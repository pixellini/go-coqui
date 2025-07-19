package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewModel_Valid(t *testing.T) {
	m, err := NewModel(modelTypeTTS, English, DatasetLJSpeech, BaseModelTacotron2DDC)
	require.NoError(t, err)
	assert.Equal(t, modelTypeTTS, m.category)
	assert.Equal(t, English, m.currentLanguage)
	assert.Equal(t, DatasetLJSpeech, m.dataset)
	assert.Equal(t, BaseModelTacotron2DDC, m.model)
	assert.True(t, m.isCustom)
}

func TestNewModel_Invalid(t *testing.T) {
	_, err := NewModel("", English, DatasetLJSpeech, BaseModelTacotron2DDC)
	assert.Error(t, err)

	_, err = NewModel(modelTypeTTS, "", DatasetLJSpeech, BaseModelTacotron2DDC)
	assert.Error(t, err)

	_, err = NewModel(modelTypeTTS, English, "", BaseModelTacotron2DDC)
	assert.Error(t, err)

	_, err = NewModel(modelTypeTTS, English, DatasetLJSpeech, "")
	assert.Error(t, err)
}

func TestModelIdentifier_Name(t *testing.T) {
	m := ModelIdentifier{
		category:        modelTypeTTS,
		currentLanguage: English,
		dataset:         DatasetLJSpeech,
		model:           BaseModelTacotron2DDC,
	}
	assert.Equal(t, "tts_models/en/ljspeech/tacotron2-DDC", m.Name())
}

func TestModelIdentifier_NameList(t *testing.T) {
	m := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		supportedLanguages: []Language{English, French},
	}
	expected := []string{
		"tts_models/ljspeech/en/tacotron2-DDC",
		"tts_models/ljspeech/fr/tacotron2-DDC",
	}
	assert.Equal(t, expected, m.NameList())
}

func TestModelIdentifier_IsValid_Validate(t *testing.T) {
	m := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		supportedLanguages: []Language{English},
	}
	assert.True(t, m.IsValid())
	assert.NoError(t, m.Validate())

	invalid := ModelIdentifier{}
	assert.False(t, invalid.IsValid())
	assert.Error(t, invalid.Validate())
}

func TestModelIdentifier_IsMultilingual(t *testing.T) {
	m := ModelIdentifier{supportedLanguages: []Language{English, French}}
	assert.True(t, m.IsMultilingual())

	m2 := ModelIdentifier{supportedLanguages: []Language{English}}
	assert.False(t, m2.IsMultilingual())
}

func TestModelIdentifier_SupportsLanguage(t *testing.T) {
	m := ModelIdentifier{supportedLanguages: []Language{English, French}}
	assert.True(t, m.SupportsLanguage(English))
	assert.False(t, m.SupportsLanguage(German))
}

func TestModelIdentifier_SupportsVoiceCloning(t *testing.T) {
	m := ModelIdentifier{isCustom: true}
	assert.True(t, m.SupportsVoiceCloning())

	m2 := ModelIdentifier{isCustom: false, supportsVoiceCloning: true}
	assert.True(t, m2.SupportsVoiceCloning())

	m3 := ModelIdentifier{isCustom: false, supportsVoiceCloning: false}
	assert.False(t, m3.SupportsVoiceCloning())
}

func TestModelIdentifier_Getters(t *testing.T) {
	m := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		currentLanguage:    English,
		defaultLanguage:    French,
		supportedLanguages: []Language{English, French},
	}
	assert.Equal(t, modelTypeTTS, m.GetCategory())
	assert.Equal(t, BaseModelTacotron2DDC, m.GetBaseModel())
	assert.Equal(t, DatasetLJSpeech, m.GetDataset())
	assert.Equal(t, English, m.GetCurrentLanguage())
	assert.Equal(t, French, m.GetDefaultLanguage())
	assert.Equal(t, []Language{English, French}, m.GetSupportedLanguages())
}

func TestModelList_FilterMethods(t *testing.T) {
	m1 := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    English,
		supportedLanguages: []Language{English, French},
	}
	m2 := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    French,
		supportedLanguages: []Language{French},
	}
	m3 := ModelIdentifier{
		category:           modelTypeTTS,
		dataset:            DatasetVCTK,
		model:              BaseModelVITS,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}
	list := ModelList[ModelIdentifier]{
		models: []ModelIdentifier{m1, m2, m3},
	}

	// FilterByBaseModel
	filtered := list.FilterByBaseModel(BaseModelTacotron2DDC)
	assert.Len(t, filtered.models, 2)

	// FilterByDataset
	filteredByDataset := list.FilterByDataset(DatasetLJSpeech)
	assert.Len(t, filteredByDataset, 2)

	// FilterBySupportedLanguages
	filteredByLang := list.FilterBySupportedLanguages([]Language{French})
	assert.Len(t, filteredByLang, 2)

	// FilterByMultilingual
	filteredMulti := list.FilterByMultilingual()
	assert.Len(t, filteredMulti, 1)

	// FilterByDefaultLanguage
	filteredByDefaultLang := list.FilterByDefaultLanguage(English)
	assert.Len(t, filteredByDefaultLang, 1)
}
