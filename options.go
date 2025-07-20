package coqui

import (
	"github.com/pixellini/go-coqui/model"
	"github.com/pixellini/go-coqui/models/vocoder"
	"github.com/pixellini/go-coqui/models/voiceconversion"
)

// Option defines an interface for TTS configuration options.
type Option interface {
	apply(*TTS) error
}

type optionFunc func(*TTS) error

// apply will set the configuration option on the TTS instance.
func (c optionFunc) apply(tts *TTS) error {
	return c(tts)
}

// WithModel sets the TTS model to use for synthesis.
// This allows access to all available Coqui TTS models.
// If the model is not valid, it will panic with an error message.
// Use WithCustomModel for custom models that may not be predefined.
func WithModelId(modelId model.Identifier) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentIdentifier(modelId)
	})
}

// WithCustomModel sets a custom TTS model by providing the model path.
// This is useful for models that are not predefined in the Coqui TTS library.
func WithModelPath(path string) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentModelPath(path)
	})
}

// WithModelLanguage sets the target language for TTS synthesis.
// Note: Language support varies by model.
func WithModelLanguage(language model.Language) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentModelLanguage(language)
	})
}

// WithVocoder sets a vocoder model to use alongside the TTS model.
func WithVocoder(v vocoder.Model) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentVocoder(v)
	})
}

// WithVocoderLanguage sets the target language for TTS synthesis.
// Note: Language support varies by model.
func WithVocoderLanguage(language model.Language) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentVocoderLanguage(language)
	})
}

// WithVoiceConversion sets a voice conversion model to use alongside the TTS model.
// TODO: Implement proper handling for voice conversion models.
func WithVoiceConversion(vcModel voiceconversion.Model) Option {
	return nil
}

// WithSpeaker sets the speaker for TTS synthesis.
// Automatically selects the appropriate configuration based on the model type.
func WithSpeaker(speaker string) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentSpeaker(speaker)
	})
}

// WithSpeakerSample sets the speaker sample file path for XTTS.
func WithSpeakerSample(path string) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentSpeakerSample(path)
	})
}

// WithSpeakerIndex sets the speaker index identifier for VITS.
func WithSpeakerIndex(idx string) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentSpeakerIndex(idx)
	})
}

// WithOutputDir sets the output directory for generated audio files.
func WithOutputDir(dir string) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentOutputDir(dir)
	})
}

// WithDevice sets the compute device for TTS synthesis.
// If Auto is specified, the best available device will be detected automatically.
func WithDevice(device model.Device) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentDevice(device)
	})
}

// WithMaxRetries sets the maximum number of synthesis attempts on failure.
func WithMaxRetries(mr int) Option {
	return optionFunc(func(t *TTS) error {
		return t.SetCurrentMaxRetries(mr)
	})
}
