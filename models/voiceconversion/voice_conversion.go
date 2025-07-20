package voiceconversion

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type Model = model.Identifier

const (
	// Voice conversion architectures.
	Freevc24    model.BaseModel = "freevc24"
	Knnvc       model.BaseModel = "knnvc"
	OpenvoiceV1 model.BaseModel = "openvoice_v1"
	OpenvoiceV2 model.BaseModel = "openvoice_v2"
)

// Voice Conversion Models for converting one voice to another.
var (
	// voice_conversion_models/multilingual/vctk/freevc24
	PresetVCTKFreeVC24 = Model{
		Category:           model.TypeVoiceConversion,
		Dataset:            model.DatasetVCTK,
		Model:              Freevc24,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(),
	}

	// voice_conversion_models/multilingual/multi-dataset/knnvc
	PresetMultidataKnnvc = Model{
		Category:           model.TypeVoiceConversion,
		Dataset:            model.DatasetVCTK,
		Model:              Knnvc,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(),
	}

	// voice_conversion_models/multilingual/multi-dataset/openvoice_v1
	PresetMultidataOpenVoiceV1 = Model{
		Category:           model.TypeVoiceConversion,
		Dataset:            model.DatasetVCTK,
		Model:              OpenvoiceV1,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(),
	}

	// voice_conversion_models/multilingual/multi-dataset/openvoice_v2
	PresetMultidataOpenVoiceV2 = Model{
		Category:           model.TypeVoiceConversion,
		Dataset:            model.DatasetVCTK,
		Model:              OpenvoiceV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(),
	}
)

// presets contains all predefined voice conversion model identifiers.
var presets = model.ModelList[Model]{
	Models: []Model{
		PresetVCTKFreeVC24,
	},
}

// New creates a new VoiceConversion model with the specified parameters.
func New(base model.BaseModel, dataset model.Dataset, language model.Language) (Model, error) {
	return model.NewModel(model.TypeVoiceConversion, base, dataset, language)
}

// GetPresets returns a list of all predefined voice conversion models.
func GetPresets() []Model {
	return slices.Clone(presets.Models)
}
