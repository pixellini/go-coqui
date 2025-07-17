package coqui

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// TTS represents a text-to-speech synthesis engine.
// Configured with a specific model, language, and device settings.
type TTS struct {
	// config holds the TTS configuration options.
	// This includes model type, language, device, and other parameters.
	config *Config
}

// ErrInvalidConfig is returned when TTS configuration validation fails.
var ErrInvalidConfig = errors.New("invalid configuration")

// New creates a new TTS instance with the specified configuration options.
func New(options ...Option) (*TTS, error) {
	// Build the config, apply the defaults
	tts := &TTS{
		config: &Config{
			Language:   English,
			TTSModel:   ModelXTTSv2, // Default to XTTS v2.
			MaxRetries: 3,           // Default retry count.
			Device:     DeviceAuto,  // Auto-detect the best device.
		},
	}

	for _, option := range options {
		option.apply(tts)
	}

	// Validate configuration.
	if err := tts.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid TTS configuration: %w", err)
	}

	fmt.Printf("\nUsing model: %s", tts.config.GetModelName())

	return tts, nil
}

/* The following are convenience constructors for common TTS models. */

// NewWithModelXttsV2 creates a new TTS instance configured for the XTTS v2 model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelXttsV2(speakerWav string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithTTSModel(ModelXTTSv2),
		WithSpeakerWav(speakerWav),
	}, options...)
	return New(opts...)
}

// NewWithModelXttsV1 creates a new TTS instance configured for the XTTS v1.1 model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelXttsV1(speakerWav string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithTTSModel(ModelXTTSv1),
		WithSpeakerWav(speakerWav),
	}, options...)
	return New(opts...)
}

// NewWithModelYourTTS creates a new TTS instance configured for the YourTTS model.
// Requires a speaker sample file path for voice cloning.
func NewWithModelYourTTS(speakerWav string, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithTTSModel(ModelYourTTS),
		WithSpeakerWav(speakerWav),
	}, options...)
	return New(opts...)
}

// NewWithModelBark creates a new TTS instance configured for the Bark model.
func NewWithModelBark(options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithTTSModel(ModelBark),
	}, options...)
	return New(opts...)
}

// NewWithSpecificModel creates a new TTS instance with a specific Model.
// This provides access to all available Coqui TTS models.
func NewFromModel(modelId TTSModel, options ...Option) (*TTS, error) {
	opts := append([]Option{
		WithTTSModel(modelId),
	}, options...)
	return New(opts...)
}

// NewFromCustomModel creates a new TTS instance from a custom TTSModel.
// This is useful for models that are not predefined in the Coqui TTS library.
func NewFromCustomModel(model TTSModel, options ...Option) (*TTS, error) {
	customModel := NewCustomTTSModel(model.defaultLanguage, model.dataset, model.architecture)
	if !customModel.IsValid() {
		return nil, fmt.Errorf("invalid custom TTS model specified: %s", model.String())
	}

	opts := append([]Option{
		WithCustomModel(customModel),
	}, options...)

	return New(opts...)
}

// NewFromConfig creates a new TTS instance from a Config struct.
// Allows loading configuration from JSON/YAML files with optional overrides.
// Additional Option parameters can override config file settings.
func NewFromConfig(config *Config, options ...Option) (*TTS, error) {
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	opts := []Option{
		WithTTSModel(config.TTSModel),
		WithLanguage(config.Language),
		WithDevice(config.Device),
		WithMaxRetries(config.MaxRetries),
	}

	if config.SpeakerWavFile != "" {
		opts = append(opts, WithSpeakerWav(config.SpeakerWavFile))
	}
	if config.SpeakerIdx != "" {
		opts = append(opts, WithSpeakerIndex(config.SpeakerIdx))
	}
	if config.DistDir != "" {
		opts = append(opts, WithDistDir(config.DistDir))
	}

	// Allow additional options to override config file settings.
	opts = append(opts, options...)

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

	_, err := os.Stat(output)
	if err == nil {
		return nil, fmt.Errorf("audio file already created")
	}

	var lastErr error
	for attempt := 1; attempt <= t.config.MaxRetries; attempt++ {
		cmdOutput, err := t.exec(ctx, text, output)
		if err == nil {
			return cmdOutput, nil
		}

		lastErr = err
		log.Print(err)
		log.Printf("TTS failed — (attempt %d/%d)\n", attempt, t.config.MaxRetries)
	}

	return nil, lastErr
}

// Config returns a copy of the current TTS configuration.
// The returned Config can be safely modified without affecting the TTS instance.
func (t TTS) GetConfig() Config {
	return *t.config
}

// exec executes the Coqui TTS command with the specified text and output path.
// This is an internal method that handles the actual subprocess execution.
func (t TTS) exec(ctx context.Context, text, output string) ([]byte, error) {
	args := t.config.ToArgs()
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
