package voiceconversion

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type VoiceConversion = model.Identifier

const (
	// Voice conversion architectures.
	BaseVoiceConversionFreevc24 model.BaseModel = "freevc24"
)

// Voice Conversion Models for converting one voice to another.
var (
	// Multilingual voice conversion.
	VoiceConversionVCTKFreeVC24 = VoiceConversion{
		Category:           model.TypeVoiceConversion,
		Dataset:            model.DatasetVCTK,
		Model:              BaseVoiceConversionFreevc24,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(),
	}
)

// VoiceConversions contains all predefined voice conversion model identifiers.
var VoiceConversions = model.ModelList[VoiceConversion]{
	Models: []VoiceConversion{
		VoiceConversionVCTKFreeVC24,
	},
}

// NewVoiceConversion creates a new VoiceConversion model with the specified parameters.
func NewVoiceConversion(base model.BaseModel, dataset model.Dataset, language model.Language) (VoiceConversion, error) {
	return model.NewModel(model.TypeVoiceConversion, base, dataset, language)
}

// GetVoiceConversionModels returns a list of all predefined voice conversion models.
func GetVoiceConversionModels() []VoiceConversion {
	return slices.Clone(VoiceConversions.Models)
}
