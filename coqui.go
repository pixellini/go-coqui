package coqui

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pixellini/go-coqui/model"
)

// TTS represents a text-to-speech synthesis engine.
// Configured with a specific model, language, and device settings.
// TTS holds the configuration for TTS synthesis.
type TTS struct {
	// model specifies the TTS model to use for synthesis.
	// This can be a specific model like ModelXTTSv2 or a custom Model.
	model model.TTSModel
	// modelPath is the path to a custom TTS model.
	// If set, this overrides the default model and uses the specified path.
	modelPath string
	// vocoder specifies the vocoder model to use for audio synthesis.
	// If not set, the default vocoder for the model will be used.
	// This is useful for advanced configurations where a specific vocoder is desired.
	vocoder model.Vocoder
	// speakerSample is the path to the speaker sample file (XTTS only).
	// Should be a clear audio sample of the desired voice (1-3 minutes recommended).
	speakerSample string
	// speakerIdx is the speaker index identifier (VITS only).
	// Use speaker IDs like "p225", "p287", ett. from the VCTK dataset.
	speakerIdx string
	// outputDir is the output directory for generated audio files.
	// If empty, files are saved to the current working directory.
	outputDir string
	// device specifies the compute device (auto/cpu/cuda/mps).
	// Use "auto" for automatic detection, "cuda" for GPU acceleration if available.
	device model.Device
	// maxRetries is the maximum number of synthesis attempts on failure.
	// Recommended range is 1-5; higher values increase reliability but slow down failure recovery.
	maxRetries int
}

const (
	defaultLanguage   = model.English
	defaultOutputDir  = "./dist/"
	defaultDevice     = model.DeviceAuto
	defaultMaxRetries = 3
)

// New creates a new TTS instance with the specified configuration options.
func New(options ...Option) (*TTS, error) {
	// Build the config, apply the defaults
	tts := &TTS{
		model:      model.TTSModelXTTSv2,
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

	if tts.model.defaultLanguage != "" && tts.model.currentLanguage == "" {
		tts.model.currentLanguage = defaultLanguage
	}
	if tts.vocoder.defaultLanguage != "" && tts.vocoder.currentLanguage == "" {
		tts.vocoder.currentLanguage = defaultLanguage
	}

	return tts, nil
}

// NewWithModelXttsV2 creates a new TTS instance configured for the XTTS v2 model.
func NewWithModelXttsV2(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(model.TTSModelXTTSv2),
	}, options...)
	return New(opts...)
}

// NewWithModelXttsV1 creates a new TTS instance configured for the XTTS v1.1 model.
func NewWithModelXttsV1(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(model.TTSModelXTTSv1),
	}, options...)
	return New(opts...)
}

// NewWithModelYourTTS creates a new TTS instance configured for the YourTTS model.
func NewWithModelYourTTS(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(model.TTSModelYourTTS),
	}, options...)
	return New(opts...)
}

// NewWithModelBark creates a new TTS instance configured for the Bark model.
func NewWithModelBark(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithModelId(model.TTSModelBark),
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
func (t TTS) Synthesize(text, outputPath string) ([]byte, error) {
	return t.SynthesizeContext(context.Background(), text, outputPath)
}

// SynthesizeContext converts text to speech with context support for cancellation.
// Supports automatic retries on failure and returns the command output on success.
// Returns an error if the output file already exists.
func (t TTS) SynthesizeContext(ctx context.Context, text, outputPath string) ([]byte, error) {
	if text == "" {
		return nil, errors.New("text cannot be empty")
	}

	return t.synthesize(ctx, text, outputPath)
}

// SynthesizeFromFile converts text from a file to speech and saves it to the specified output file.
func (t TTS) SynthesizeFromFile(filePath, outputPath string) ([]byte, error) {
	return t.SynthesizeFromFileContext(context.Background(), filePath, outputPath)
}

// SynthesizeFromFileContext converts text from a file to speech with context support.
func (t TTS) SynthesizeFromFileContext(ctx context.Context, filePath, outputPath string) ([]byte, error) {
	if filePath == "" {
		return nil, errors.New("file path cannot be empty")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	return t.synthesize(ctx, string(content), outputPath)
}

// synthesize runs the TTS command to convert text to speech.
func (t TTS) synthesize(ctx context.Context, text, outputPath string) ([]byte, error) {
	// Create the dist directory if it doesn't exist
	if err := os.MkdirAll(t.outputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create dist directory: %w", err)
	}

	outputPath = t.outputDir + outputPath

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

// run executes the Coqui TTS command with the specified text and output path.
// This is an internal method that handles the actual subprocess execution.
func (t TTS) run(ctx context.Context, text, outputPath string) ([]byte, error) {
	args := toArgs(t)
	args = append(args,
		argText, text,
		argOutPath, outputPath,
	)

	fmt.Printf("\nProcessing text: %q\n", text)

	cmd := exec.CommandContext(ctx, "tts", args...)

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\nTTS command failed with output: %s\n", cmdOutput)
		return cmdOutput, fmt.Errorf("TTS command failed: %w", err)
	}

	return cmdOutput, nil
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
	return fmt.Sprintf("%s/%s/%s/%s", t.vocoder.category, t.vocoder.defaultLanguage, t.vocoder.dataset, t.vocoder.model)
}

// CurrentModel returns the Model being used for synthesis.
func (t TTS) CurrentModel() model.Model {
	return t.model
}

// CurrentVocoder returns the VocoderModel being used for synthesis.
func (t TTS) CurrentVocoder() model.Vocoder {
	return t.vocoder
}

// CurrentModelLanguage returns the model.Language being used for synthesis.
func (t TTS) CurrentModelLanguage() model.Language {
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
func (t TTS) CurrentDevice() model.Device {
	return t.device
}

// CurrentMaxRetries returns the maximum number of retries for synthesis attempts.
func (t TTS) CurrentMaxRetries() int {
	return t.maxRetries
}

// SetCurrentModel sets the TTS model to use for synthesis.
func (t *TTS) SetCurrentModelIdentifier(m model.ModelIdentifier) error {
	if err := m.Validate(); err != nil {
		return fmt.Errorf("invalid TTS model specified: %s", err)
	}
	t.model = m
	t.model.currentLanguage = m.defaultLanguage
	t.modelPath = ""
	return nil
}

// SetCurrentModelPath sets the path to a custom TTS model.
func (t *TTS) SetCurrentModelPath(p string) error {
	if p == "" {
		return fmt.Errorf("model path cannot be empty")
	}

	// Check if the model path exists
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return fmt.Errorf("model path does not exist: %s", p)
	}

	t.modelPath = p
	t.model.isCustom = true // Mark as custom model
	return nil
}

// SetCurrentVocoder sets the vocoder model to use for audio synthesis.
func (t *TTS) SetCurrentVocoder(v model.Vocoder) error {
	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid Vocoder specified: %s", err)
	}
	t.vocoder = v
	return nil
}

// SetCurrentModelLanguage sets the target language for synthesis.
func (t *TTS) SetCurrentModelLanguage(l model.Language) error {
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
func (t *TTS) SetCurrentVocoderLanguage(l model.Language) error {
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
func (t *TTS) SetCurrentDevice(device model.Device) error {
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
