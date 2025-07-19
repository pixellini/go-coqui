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
		supportedLanguages: supportedLanguages,
	}
)

// VoiceConversionModelList contains all predefined voice conversion model identifiers.
var VoiceConversionModelList = ModelList[VoiceConversion]{
	models: []VoiceConversion{
		VoiceConversionVCTKFreeVC24,
	},
}

// NewVoiceConversion creates a new VoiceConversion model with the specified parameters.
func NewVoiceConversion(language Language, dataset Dataset, model BaseModel) (VoiceConversion, error) {
	return NewModel(modelTypeVoiceConversion, language, dataset, model)
}

// GetVoiceConversionModels returns a list of all predefined voice conversion models.
func GetVoiceConversionModels() []VoiceConversion {
	return slices.Clone(VoiceConversionModelList.models)
}
