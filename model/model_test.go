package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var MockBaseModel = BaseModel("mock-base-model")
var MockDataset = Dataset("mock-dataset")
var MockDataset2 = Dataset("mock-dataset2")

func TestNewModel_Valid(t *testing.T) {
	m, err := NewModel(TypeTTS, MockBaseModel, MockDataset, English)
	require.NoError(t, err)
	assert.Equal(t, TypeTTS, m.Category)
	assert.Equal(t, English, m.CurrentLanguage)
	assert.Equal(t, MockDataset, m.Dataset)
	assert.Equal(t, MockBaseModel, m.Model)
	assert.True(t, m.IsCustom)
}

func TestNewModel_Invalid(t *testing.T) {
	_, err := NewModel("", MockBaseModel, MockDataset, English)
	assert.Error(t, err)

	_, err = NewModel(TypeTTS, MockBaseModel, MockDataset, "")
	assert.Error(t, err)

	_, err = NewModel(TypeTTS, MockBaseModel, "", English)
	assert.Error(t, err)

	_, err = NewModel(TypeTTS, "", MockDataset, English)
	assert.Error(t, err)
}

func TestIdentifier_Name(t *testing.T) {
	m := Identifier{
		Category:        TypeTTS,
		CurrentLanguage: English,
		Dataset:         MockDataset,
		Model:           MockBaseModel,
	}
	assert.Equal(t, "tts_models/en/ljspeech/tacotron2-DDC", m.Name())
}

func TestIdentifier_NameList(t *testing.T) {
	m := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset,
		Model:              MockBaseModel,
		SupportedLanguages: []Language{English, French},
	}
	expected := []string{
		"tts_models/ljspeech/en/tacotron2-DDC",
		"tts_models/ljspeech/fr/tacotron2-DDC",
	}
	assert.Equal(t, expected, m.NameList())
}

func TestIdentifier_IsValid_Validate(t *testing.T) {
	m := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset,
		Model:              MockBaseModel,
		SupportedLanguages: []Language{English},
	}
	assert.True(t, m.IsValid())
	assert.NoError(t, m.Validate())

	invalid := Identifier{}
	assert.False(t, invalid.IsValid())
	assert.Error(t, invalid.Validate())
}

func TestIdentifier_IsMultilingual(t *testing.T) {
	m := Identifier{SupportedLanguages: []Language{English, French}}
	assert.True(t, m.IsMultilingual())

	m2 := Identifier{SupportedLanguages: []Language{English}}
	assert.False(t, m2.IsMultilingual())
}

func TestIdentifier_SupportsLanguage(t *testing.T) {
	m := Identifier{SupportedLanguages: []Language{English, French}}
	assert.True(t, m.SupportsLanguage(English))
	assert.False(t, m.SupportsLanguage(German))
}

func TestIdentifier_SupportsVoiceCloning(t *testing.T) {
	m := Identifier{IsCustom: true}
	assert.True(t, m.SupportsCloning())

	m2 := Identifier{IsCustom: false, SupportsVoiceCloning: true}
	assert.True(t, m2.SupportsCloning())

	m3 := Identifier{IsCustom: false, SupportsVoiceCloning: false}
	assert.False(t, m3.SupportsCloning())
}

func TestIdentifier_Getters(t *testing.T) {
	m := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset,
		Model:              MockBaseModel,
		CurrentLanguage:    English,
		DefaultLanguage:    French,
		SupportedLanguages: []Language{English, French},
	}
	assert.Equal(t, TypeTTS, m.GetType())
	assert.Equal(t, MockBaseModel, m.GetBaseModel())
	assert.Equal(t, MockDataset, m.GetDataset())
	assert.Equal(t, English, m.GetCurrentLanguage())
	assert.Equal(t, French, m.GetDefaultLanguage())
	assert.Equal(t, []Language{English, French}, m.GetSupportedLanguages())
}

func TestModelList_FilterMethods(t *testing.T) {
	m1 := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset,
		Model:              MockBaseModel,
		DefaultLanguage:    English,
		SupportedLanguages: []Language{English, French},
	}
	m2 := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset,
		Model:              MockBaseModel,
		DefaultLanguage:    French,
		SupportedLanguages: []Language{French},
	}
	m3 := Identifier{
		Category:           TypeTTS,
		Dataset:            MockDataset2,
		Model:              MockBaseModel,
		DefaultLanguage:    German,
		SupportedLanguages: []Language{German},
	}
	list := ModelList[Identifier]{
		Models: []Identifier{m1, m2, m3},
	}

	// FilterByBaseModel
	filtered := list.FilterByBaseModel(MockBaseModel)
	assert.Len(t, filtered.Models, 2)

	// FilterByDataset
	filteredByDataset := list.FilterByDataset(MockDataset)
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
