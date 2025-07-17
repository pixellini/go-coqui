package coqui

import (
	"fmt"
	"slices"
)

// VoiceConversionModel represents a voice conversion model identifier.
type VoiceConversionModel struct {
	modelType          ModelType
	dataset            Dataset
	architecture       Architecture
	defaultLanguage    Language
	supportedLanguages []Language
}

// Voice Conversion Models for converting one voice to another.
var (
	// Multilingual voice conversion.
	VoiceConversionVCTKFreeVC24 = VoiceConversionModel{
		modelType:          modelTypeVoiceConversion,
		dataset:            DatasetVCTK,
		architecture:       ArchFreevc24,
		defaultLanguage:    English,
		supportedLanguages: allSupportedLanguages,
	}
)

// allVoiceConversions contains all predefined voice conversion model identifiers.
var allVoiceConversions = []VoiceConversionModel{
	VoiceConversionVCTKFreeVC24,
}

func NewVoiceConversion(language Language, dataset Dataset, architecture Architecture) VoiceConversionModel {
	return VoiceConversionModel{
		modelType:          modelTypeVoiceConversion,
		dataset:            dataset,
		architecture:       architecture,
		supportedLanguages: []Language{Language(language)},
		defaultLanguage:    Language(language),
	}
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (vc VoiceConversionModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", vc.GetModelType(), vc.GetDataset(), vc.GetDefaultLanguage(), vc.GetArchitecture())
}

// NameList returns a list of string representations of the model identifier for each supported language.
func (vc VoiceConversionModel) NameList() []string {
	var names []string
	for _, lang := range vc.supportedLanguages {
		name := fmt.Sprintf("%s/%s/%s/%s", vc.GetModelType(), vc.GetDataset(), lang, vc.GetArchitecture())
		names = append(names, name)
	}
	return names
}

// IsValid checks if the model identifier is valid.
func (vc VoiceConversionModel) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), vc.modelType)
	// Check if the model has a valid type, dataset, architecture, and supported languages.
	return containsValidType && vc.modelType != "" && vc.dataset != "" && vc.architecture != "" && len(vc.supportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (vc VoiceConversionModel) IsMultilingual() bool {
	return len(vc.supportedLanguages) > 1
}

// GetModelType returns the model type.
func (vc VoiceConversionModel) GetModelType() ModelType {
	return vc.modelType
}

// GetArchitecture returns the model architecture.
func (vc VoiceConversionModel) GetArchitecture() Architecture {
	return vc.architecture
}

// GetDataset returns the model dataset.
func (vc VoiceConversionModel) GetDataset() Dataset {
	return vc.dataset
}

// GetSupportedLanguages returns the supported languages.
func (vc VoiceConversionModel) GetSupportedLanguages() []Language {
	return vc.supportedLanguages
}

// GetDefaultLanguage returns the default language.
func (vc VoiceConversionModel) GetDefaultLanguage() Language {
	return vc.defaultLanguage
}

// GetAllVoiceConversions returns a copy of the list of all predefined voice conversion models.
func GetAllVoiceConversions() []VoiceConversionModel {
	return slices.Clone(allVoiceConversions)
}

// GetVoiceConversionByArchitecture returns all voice conversion models that use the specified architecture.
func GetVoiceConversionByArchitecture(arch Architecture) []VoiceConversionModel {
	return filterModelsByArchitecture(allVoiceConversions, arch)
}

// GetVoiceConversionByDataset returns all voice conversion models of the specified dataset.
func GetVoiceConversionByDataset(dataset Dataset) []VoiceConversionModel {
	return filterModelsByDataset(allVoiceConversions, dataset)
}

// GetVoiceConversionByLanguage returns all voice conversion models that support the specified language.
func GetVoiceConversionByLanguage(language Language) []VoiceConversionModel {
	return filterModelsBySupportedLanguages(allVoiceConversions, []Language{language})
}

// GetVoiceConversionBySupportedLanguages returns all voice conversion models that support any of the specified languages.
func GetVoiceConversionBySupportedLanguages(languages []Language) []VoiceConversionModel {
	return filterModelsBySupportedLanguages(allVoiceConversions, languages)
}

// GetVoiceConversionByDefaultLanguage returns all voice conversion models with the specified default language.
func GetVoiceConversionByDefaultLanguage(language Language) []VoiceConversionModel {
	return filterModelsByDefaultLanguage(allVoiceConversions, language)
}

// GetVoiceConversionMultilingual returns all voice conversion models that support multiple languages.
func GetVoiceConversionMultilingual() []VoiceConversionModel {
	return filterModelsMultilingual(allVoiceConversions)
}
