package coqui

import "slices"

type VoiceConversion = ModelIdentifier

const (
	// Voice conversion architectures.
	BaseVoiceConversionFreevc24 BaseModel = "freevc24"
)

// Voice Conversion Models for converting one voice to another.
var (
	// Multilingual voice conversion.
	VoiceConversionVCTKFreeVC24 = VoiceConversion{
		category:           modelTypeVoiceConversion,
		dataset:            DatasetVCTK,
		model:              BaseVoiceConversionFreevc24,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: supportedLanguages,
	}
)

// VoiceConversions contains all predefined voice conversion model identifiers.
var VoiceConversions = ModelList[VoiceConversion]{
	models: []VoiceConversion{
		VoiceConversionVCTKFreeVC24,
	},
}

// NewVoiceConversion creates a new VoiceConversion model with the specified parameters.
func NewVoiceConversion(model BaseModel, dataset Dataset, language Language) (VoiceConversion, error) {
	return NewModel(modelTypeVoiceConversion, model, dataset, language)
}

// GetVoiceConversionModels returns a list of all predefined voice conversion models.
func GetVoiceConversionModels() []VoiceConversion {
	return slices.Clone(VoiceConversions.models)
}
