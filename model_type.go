package coqui

import "slices"

// Category represents the category of model.
type Category string

const (
	// modelTypeTTS represents text-to-speech models.
	modelTypeTTS Category = "tts_models"
	// modelTypeVocoder represents vocoder models for audio synthesis.
	modelTypeVocoder Category = "vocoder_models"
	// modelTypeVoiceConversion represents voice conversion models.
	modelTypeVoiceConversion Category = "voice_conversion_models"
)

// modelTypes contains all predefined model types.
var modelTypes = []Category{
	modelTypeTTS,
	modelTypeVocoder,
	modelTypeVoiceConversion,
}

func GetModelTypes() []Category {
	return slices.Clone(modelTypes)
}
