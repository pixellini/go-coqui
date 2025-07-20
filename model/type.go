package model

import "slices"

// Type represents the category of model.
type Type string

const (
	// TypeTTS represents text-to-speech models.
	TypeTTS Type = "tts_models"
	// TypeVocoder represents vocoder models for audio synthesis.
	TypeVocoder Type = "vocoder_models"
	// TypeVoiceConversion represents voice conversion models.
	TypeVoiceConversion Type = "voice_conversion_models"
)

// types contains all predefined model types.
var types = []Type{
	TypeTTS,
	TypeVocoder,
	TypeVoiceConversion,
}

func GetTypes() []Type {
	return slices.Clone(types)
}
