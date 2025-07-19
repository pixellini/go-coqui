package coqui

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// TTS represents a text-to-speech synthesis engine.
// Configured with a specific model, language, and device settings.
// TTS holds the configuration for TTS synthesis.
type TTS struct {
	// Model specifies the TTS model to use for synthesis.
	// This can be a specific model like ModelXTTSv2 or a custom Model.
	model TTSModel
	// Vocoder specifies the vocoder model to use for audio synthesis.
	// If not set, the default vocoder for the model will be used.
	// This is useful for advanced configurations where a specific vocoder is desired.
	vocoder Vocoder
	// SpeakerSample is the path to the speaker sample file (XTTS only).
	// Should be a clear audio sample of the desired voice (1-3 minutes recommended).
	speakerSample string
	// SpeakerIdx is the speaker index identifier (VITS only).
	// Use speaker IDs like "p225", "p287", ett. from the VCTK dataset.
	speakerIdx string
	// OutputDir is the output directory for generated audio files.
	// If empty, files are saved to the current working directory.
	outputDir string
	// Device specifies the compute device (auto/cpu/cuda/mps).
	// Use "auto" for automatic detection, "cuda" for GPU acceleration if available.
	device Device
	// MaxRetries is the maximum number of synthesis attempts on failure.
	// Recommended range is 1-5; higher values increase reliability but slow down failure recovery.
	maxRetries int
}

const (
	defaultLanguage   = English
	defaultOutputDir  = "./dist"
	defaultDevice     = DeviceAuto
	defaultMaxRetries = 3
)

// New creates a new TTS instance with the specified configuration options.
func New(options ...Option) (*TTS, error) {
	// Build the config, apply the defaults
	tts := &TTS{
		model:      TTSModelXTTSv2,
		outputDir:  defaultOutputDir,
		device:     defaultDevice,
		maxRetries: defaultMaxRetries,
	}

	for _, option := range options {
		err := option.apply(tts)
		if err != nil {
			return nil, fmt.Errorf("failed to create TTS instance: %w", err)
		}
	}

	fmt.Printf("\nUsing model: %s", tts.Name())

	return tts, nil
}

// NewWithModelXttsV2 creates a new TTS instance configured for the XTTS v2 model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelXttsV2(samplePath string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(TTSModelXTTSv2),
		WithSpeakerSample(samplePath),
	}, options...)
	return New(opts...)
}

// NewWithModelXttsV1 creates a new TTS instance configured for the XTTS v1.1 model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelXttsV1(samplePath string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(TTSModelXTTSv1),
		WithSpeakerSample(samplePath),
	}, options...)
	return New(opts...)
}

// NewWithModelYourTTS creates a new TTS instance configured for the YourTTS model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelYourTTS(samplePath string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(TTSModelYourTTS),
		WithSpeakerSample(samplePath),
	}, options...)
	return New(opts...)
}

// NewWithModelBark creates a new TTS instance configured for the Bark model.
func NewWithModelBark(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(TTSModelBark),
	}, options...)
	return New(opts...)
}

// Configure applies additional configuration options to the TTS instance.
// Use this to modify settings after the TTS instance has been created.
func (t *TTS) Configure(options ...Option) {
	for _, option := range options {
		option.apply(t)
	}
}

// Synthesize converts text to speech and saves it to the specified output file.
// This is a convenience method that uses context.Background().
func (t TTS) Synthesize(text, output string) ([]byte, error) {
	return t.SynthesizeContext(context.Background(), text, output)
}

// SynthesizeContext converts text to speech with context support for cancellation.
// Supports automatic retries on failure and returns the command output on success.
// Returns an error if the output file already exists.
func (t TTS) SynthesizeContext(ctx context.Context, text, output string) ([]byte, error) {
	if text == "" {
		return nil, errors.New("text cannot be empty")
	}

	// Create the dist directory if it doesn't exist
	if err := os.MkdirAll(t.outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create dist directory: %w", err)
	}

	outputPath := t.outputDir + output

	_, err := os.Stat(outputPath)
	if err == nil {
		return nil, fmt.Errorf("audio file already created")
	}

	var lastErr error
	for attempt := 1; attempt <= t.maxRetries; attempt++ {
		cmdOutput, err := t.run(ctx, text, outputPath)
		if err == nil {
			return cmdOutput, nil
		}

		lastErr = err
		log.Print(err)
		log.Printf("TTS failed — (attempt %d/%d)\n", attempt, t.maxRetries)
	}

	return nil, lastErr
}

// Name returns the full Coqui TTS model name to use.
// Returns empty string if no model is configured.
// Format: tts_models/{language}/{dataset}/{model}
// For multilingual models, uses "multilingual" instead of specific language.
func (t TTS) Name() string {
	// Use "multilingual" for models that support all languages.
	// NOTE: This is currently a workaround so I can test the functionality.
	// But this will break if the model supports multiple languages but is not "multilingual" as defined in the model name.
	// TODO: Fix this to properly handle "multilingual" vs multilingual models.
	language := t.model.currentLanguage.String()
	if t.model.IsMultilingual() {
		language = "multilingual"
	}
	return fmt.Sprintf("%s/%s/%s/%s", t.model.category, language, t.model.dataset, t.model.model)
}

// VocoderName returns the full Coqui TTS vocoder name to use.
// Format: vocoder_models/{language}/{dataset}/{model}
func (t TTS) VocoderName() string {
	return fmt.Sprintf("%s/%s/%s/%s", t.vocoder.model, t.vocoder.defaultLanguage, t.vocoder.dataset, t.vocoder.model)
}

// CurrentModel returns the Model being used for synthesis.
func (t TTS) CurrentModel() Model {
	return t.model
}

// CurrentVocoder returns the VocoderModel being used for synthesis.
func (t TTS) CurrentVocoder() Vocoder {
	return t.vocoder
}

// CurrentModelLanguage returns the Language being used for synthesis.
func (t TTS) CurrentModelLanguage() Language {
	return t.model.currentLanguage
}

// CurrentSpeakerSample returns the path to the speaker sample file.
func (t TTS) CurrentSpeakerSample() string {
	return t.speakerSample
}

// CurrentSpeakerIndex returns the speaker index identifier.
func (t TTS) CurrentSpeakerIndex() string {
	return t.speakerIdx
}

// CurrentOutputDir returns the output directory where audio files will be saved.
func (t TTS) CurrentOutputDir() string {
	return t.outputDir
}

// CurrentDevice returns the compute device used for synthesis.
func (t TTS) CurrentDevice() Device {
	return t.device
}

// CurrentMaxRetries returns the maximum number of retries for synthesis attempts.
func (t TTS) CurrentMaxRetries() int {
	return t.maxRetries
}

// SetCurrentModel sets the TTS model to use for synthesis.
func (t *TTS) SetCurrentModelIdentifier(m ModelIdentifier) error {
	if err := m.Validate(); err != nil {
		return fmt.Errorf("invalid TTS model specified: %s", err)
	}
	t.model = m
	t.model.currentLanguage = m.defaultLanguage
	return nil
}

// SetCurrentVocoder sets the vocoder model to use for audio synthesis.
func (t *TTS) SetCurrentVocoder(v Vocoder) error {
	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid Vocoder specified: %s", err)
	}
	t.vocoder = v
	return nil
}

// SetCurrentModelLanguage sets the target language for synthesis.
func (t *TTS) SetCurrentModelLanguage(l Language) error {
	if !l.IsSupported() {
		return fmt.Errorf("invalid language specified: %s", l.String())
	}
	if !t.model.SupportsLanguage(l) {
		return fmt.Errorf("model %s does not support language %s", t.model.Name(), l.String())
	}
	t.model.currentLanguage = l
	return nil
}

// SetCurrentVocoderLanguage sets the target language for synthesis.
func (t *TTS) SetCurrentVocoderLanguage(l Language) error {
	if !l.IsSupported() {
		return fmt.Errorf("invalid language specified: %s", l.String())
	}
	if !t.vocoder.SupportsLanguage(l) {
		return fmt.Errorf("vocoder %s does not support language %s", t.vocoder.Name(), l.String())
	}
	t.vocoder.currentLanguage = l
	return nil
}

// SetCurrentSpeaker sets the current speaker for voice cloning.
func (t *TTS) SetCurrentSpeaker(s string) error {
	if s == "" {
		return fmt.Errorf("speaker cannot be empty")
	}
	fmt.Printf("Speak path: %s\n", filepath.Ext(s))
	// speaker has an extension (e.g. ".wav", ".mp3").
	if filepath.Ext(s) != "" && t.model.SupportsVoiceCloning() {
		t.speakerSample = s
	} else {
		t.speakerIdx = s
	}
	return nil
}

// SetCurrentSpeakerSample sets the path to the speaker sample file for voice cloning.
func (t *TTS) SetCurrentSpeakerSample(samplePath string) error {
	if samplePath == "" {
		return fmt.Errorf("speaker sample path cannot be empty")
	}

	t.speakerSample = samplePath
	return nil
}

// SetCurrentSpeakerIndex sets the speaker index identifier for VITS models.
func (t *TTS) SetCurrentSpeakerIndex(idx string) error {
	if idx == "" {
		return fmt.Errorf("speaker index cannot be empty")
	}

	t.speakerIdx = idx
	return nil
}

// SetCurrentOutputDir sets the output directory for generated audio files.
func (t *TTS) SetCurrentOutputDir(dir string) error {
	if dir == "" {
		return fmt.Errorf("output directory cannot be empty")
	}

	t.outputDir = dir
	return nil
}

// SetCurrentDevice sets the compute device for synthesis.
func (t *TTS) SetCurrentDevice(device Device) error {
	if !device.IsValid() {
		return fmt.Errorf("invalid device specified: %s", device.String())
	}

	t.device = device
	return nil
}

// SetCurrentMaxRetries sets the maximum number of retries for synthesis attempts.
func (t *TTS) SetCurrentMaxRetries(r int) error {
	if r < 1 {
		return fmt.Errorf("max retries must be at least 1")
	}

	t.maxRetries = r
	return nil
}

// toArgs converts the TTS configuration to command-line arguments.
// for the underlying Coqui TTS Python process.
// TODO: There are other arguments that can be added based on the model type.
// There's also a lot of room for improvement here, but for now,
// this function generates the basic arguments needed for synthesis.
func (t TTS) toArgs() []string {
	// Resolve "auto" device to actual device.
	device := t.device
	if device == DeviceAuto {
		device = detectDevice()
	}

	args := []string{
		"--model_name", t.Name(),
		"--device", device.String(),
	}

	// Explicitly set CUDA usage based on device.
	if device == DeviceCUDA {
		args = append(args, "--use_cuda", "true")
	}

	// TODO: Handle vocoder if specified.
	if t.vocoder.IsValid() {
		args = append(args, "--vocoder_name", t.VocoderName())
	}

	// TODO: Handle Voice Conversion models.

	lang := t.model.currentLanguage.String()
	// We don't know the model type at this point, and we won't know if the model supports voice cloning until we run the command.
	// So we need to handle the speaker sample and index based on what the user has set.
	if t.model.isCustom {
		if t.speakerSample != "" {
			args = append(args, "--speaker_wav", t.speakerSample)
			args = append(args, "--language_idx", lang)
		} else {
			args = append(args, "--speaker_idx", t.speakerIdx)
		}
	} else {
		// Handle voice cloning models (XTTS variants, YourTTS).
		if t.model.SupportsVoiceCloning() {
			if t.speakerSample != "" {
				args = append(args, "--speaker_wav", t.speakerSample)
			}

			args = append(args, "--language_idx", lang)
		}

		if t.speakerIdx != "" {
			args = append(args, "--speaker_idx", t.speakerIdx)
		}
	}

	fmt.Printf("\nArgs: %v\n", args)

	return args
}

// run executes the Coqui TTS command with the specified text and output path.
// This is an internal method that handles the actual subprocess execution.
func (t TTS) run(ctx context.Context, text, output string) ([]byte, error) {
	args := t.toArgs()
	args = append(args,
		"--text", text,
		"--out_path", output,
	)

	cmd := exec.CommandContext(ctx, "tts", args...)

	fmt.Printf("\nProcessing text: %q", text)

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\nTTS command failed with output: %s\n", cmdOutput)
		return cmdOutput, fmt.Errorf("TTS command failed: %w", err)
	}

	return cmdOutput, nil
}
