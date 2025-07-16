package coqui

import "fmt"

// Config holds the configuration for TTS synthesis.
type Config struct {
	// Model specifies the TTS model to use for synthesis.
	// This can be a specific model like ModelXTTSv2 or a custom Model.
	TTSModel TTSModel `json:"model" yaml:"model"`
	// Vocoder specifies the vocoder model to use for audio synthesis.
	// If not set, the default vocoder for the model will be used.
	// This is useful for advanced configurations where a specific vocoder is desired.
	Vocoder VocoderModel `json:"vocoder,omitempty" yaml:"vocoder,omitempty"`
	// Language specifies the target language for synthesis.
	// See language.go for all supported language codes.
	Language Language `json:"language" yaml:"language"`
	// SpeakerWavFile is the path to the speaker sample file (XTTS only).
	// Should be a clear audio sample of the desired voice (1-3 minutes recommended).
	SpeakerWavFile string `json:"speakerWavFile" yaml:"speakerWavFile"`
	// SpeakerIdx is the speaker index identifier (VITS only).
	// Use speaker IDs like "p225", "p287", etc. from the VCTK dataset.
	SpeakerIdx string `json:"speakerIdx" yaml:"speakerIdx"`
	// MaxRetries is the maximum number of synthesis attempts on failure.
	// Recommended range is 1-5; higher values increase reliability but slow down failure recovery.
	MaxRetries int `json:"maxRetries" yaml:"maxRetries"`
	// DistDir is the output directory for generated audio files.
	// If empty, files are saved to the current working directory.
	DistDir string `json:"distDir" yaml:"distDir"`
	// Device specifies the compute device (auto/cpu/cuda/mps).
	// Use "auto" for automatic detection, "cuda" for GPU acceleration if available.
	Device Device `json:"device" yaml:"device"`
}

// GetModel returns the Model being used for synthesis.
func (c *Config) GetModel() TTSModel {
	return c.TTSModel
}

// GetModelName returns the full Coqui TTS model name to use.
// Returns empty string if no model is configured.
// Format: tts_models/{language}/{dataset}/{architecture}
// For multilingual models, uses "multilingual" instead of specific language
func (c *Config) GetModelName() string {
	// Use "multilingual" for models that support all languages
	// NOTE: This is currently a workaround so I can test the functionality.
	// But this will break if the model supports multiple languages but is not "multilingual" as defined in the model name.
	// TODO: Fix this to properly handle "multilingual" vs multilingual models.
	language := c.Language.String()
	if c.TTSModel.IsMultilingual() {
		language = "multilingual"
	}
	return fmt.Sprintf("%s/%s/%s/%s", c.TTSModel.Type, language, c.TTSModel.Dataset, c.TTSModel.Architecture)
}

// GetVocoderName returns the full Coqui TTS vocoder name to use.
// Format: vocoder_models/{language}/{dataset}/{architecture}
func (c *Config) GetVocoderName() string {
	return fmt.Sprintf("%s/%s/%s/%s", c.Vocoder.Type, c.Language, c.Vocoder.Dataset, c.Vocoder.Architecture)
}

// SupportsVoiceCloning returns true if the effective model supports voice cloning.
// TODO: Add support for more models as needed.
func (c *Config) SupportsVoiceCloning() bool {
	// XTTS v1, XTTS v2, and YourTTS support voice cloning
	return c.TTSModel.Architecture == ArchXTTSv2 ||
		c.TTSModel.Architecture == ArchXTTSv1 ||
		c.TTSModel.Architecture == ArchYourTTS ||
		c.TTSModel.Architecture == ArchBark
}

// RequiresSpeakerIndex returns true if the effective model requires a speaker index.
// TODO: Configure models that require speaker index.
func (c *Config) RequiresSpeakerIndex() bool {
	return c.TTSModel.Architecture == ArchVITS
}

// Validate checks if the TTS configuration is valid and returns an error
// if any configuration values are invalid or incompatible.
// TODO: Implement proper validation logic for TTS configurations.
func (c *Config) Validate() error {
	return nil
}

// ToArgs converts the TTS configuration to command-line arguments
// for the underlying Coqui TTS Python process.
// TODO: There are other arguments that can be added based on the model type.
// There's also a lot of room for improvement here, but for now
// this function generates the basic arguments needed for synthesis.
func (c *Config) ToArgs() []string {
	// Resolve "auto" device to actual device
	device := c.Device
	if device == DeviceAuto {
		device = DetectDevice(device)
	}

	args := []string{
		"--model_name", c.GetModelName(),
		"--device", device.String(),
	}

	// Explicitly set CUDA usage based on device
	if device == DeviceCUDA {
		args = append(args, "--use_cuda", "true")
	}

	// TODO: Handle vocoder if specified

	// TODO: Handle Voice Conversion models

	// Handle voice cloning models (XTTS variants, YourTTS)
	if c.SupportsVoiceCloning() {
		if c.SpeakerWavFile != "" {
			args = append(args, "--speaker_wav", c.SpeakerWavFile)
		}

		lang := c.Language.String()
		if !c.TTSModel.SupportsLanguage(c.Language) {
			fmt.Println("\nWarning: Model does not support specified language, using default language instead.")
			lang = string(c.TTSModel.DefaultLanguage)
		}
		args = append(args, "--language_idx", lang)
	}

	// Handle models that require speaker index
	if c.RequiresSpeakerIndex() {
		args = append(args, "--speaker_idx", c.SpeakerIdx)
	}

	return args
}
