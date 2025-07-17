package coqui

// Option defines an interface for TTS configuration options.
type Option interface {
	apply(*TTS)
}

type optionFunc func(*TTS)

// apply will set the configuration option on the TTS instance.
func (c optionFunc) apply(tts *TTS) {
	c(tts)
}

// WithTTSModel sets the TTS model to use for synthesis.
// This allows access to all available Coqui TTS models.
// If the model is not valid, it will panic with an error message.
// Use WithCustomModel for custom models that may not be predefined.
func WithTTSModel(model TTSModel) Option {
	return optionFunc(func(t *TTS) {
		// If the model is not valid, return an error.
		if !model.IsValid() {
			panic("invalid TTS model specified: " + model.String())
		}
		t.config.TTSModel = model
	})
}

// WithCustomModel sets a custom TTS model to use for synthesis.
// This is useful for models that are not predefined in the Coqui TTS library.
func WithCustomModel(ttsModel TTSModel) Option {
	return optionFunc(func(t *TTS) {
		// We don't need to check if the model is valid here because it's custom.
		t.config.TTSModel = NewTTSModel(ttsModel.defaultLanguage, ttsModel.dataset, ttsModel.architecture)
	})
}

// WithVocoderModel sets a vocoder model to use alongside the TTS model.
// TODO: Implement proper handling for vocoder models.
func WithVocoderModel(vModel VocoderModel) Option {
	return nil
}

// WithVoiceConversionModel sets a voice conversion model to use alongside the TTS model.
// TODO: Implement proper handling for voice conversion models.
func WithVoiceConversionModel(vcModel VoiceConversionModel) Option {
	return nil
}

// WithLanguage sets the target language for TTS synthesis.
// Note: Language support varies by model.
func WithLanguage(language Language) Option {
	return optionFunc(func(t *TTS) {
		t.config.Language = language
	})
}

// WithSpeaker sets the speaker for TTS synthesis.
// Automatically selects the appropriate configuration based on the model type.
func WithSpeaker(speaker string) Option {
	return optionFunc(func(t *TTS) {
		if t.config.SupportsVoiceCloning() {
			WithSpeakerWav(speaker).apply(t)
		}
		if t.config.RequiresSpeakerIndex() {
			WithSpeakerIndex(speaker).apply(t)
		}
	})
}

// WithSpeakerWav sets the speaker sample file path for XTTS.
func WithSpeakerWav(path string) Option {
	return optionFunc(func(t *TTS) {
		t.config.SpeakerWavFile = path
	})
}

// WithSpeakerIndex sets the speaker index identifier for VITS.
func WithSpeakerIndex(idx string) Option {
	return optionFunc(func(t *TTS) {
		t.config.SpeakerIdx = idx
	})
}

// WithMaxRetries sets the maximum number of synthesis attempts on failure.
func WithMaxRetries(mr int) Option {
	return optionFunc(func(t *TTS) {
		t.config.MaxRetries = mr
	})
}

// WithDistDir sets the output directory for generated audio files.
func WithDistDir(outputPath string) Option {
	return optionFunc(func(t *TTS) {
		t.config.DistDir = outputPath
	})
}

// WithDevice sets the compute device for TTS synthesis.
// If Auto is specified, the best available device will be detected automatically.
func WithDevice(device Device) Option {
	return optionFunc(func(t *TTS) {
		// If user explicitly chooses Auto, detect the best device.
		if device == DeviceAuto {
			device = DetectDevice(DeviceAuto)
		}
		t.config.Device = device
	})
}
