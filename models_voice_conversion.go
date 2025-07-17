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
	supportedLanguages []Language
	defaultLanguage    Language
}

// Voice Conversion Models for converting one voice to another.
var (
	// Multilingual voice conversion.
	VoiceConversionVCTKFreeVC24 = VoiceConversionModel{
		modelType:          ModelTypeVoiceConversion,
		supportedLanguages: GetAllSupportedLanguages(),
		dataset:            DatasetVCTK,
		architecture:       ArchFreevc24,
		defaultLanguage:    English,
	}
)

// allVoiceConversions contains all predefined voice conversion model identifiers.
var allVoiceConversions = []VoiceConversionModel{
	VoiceConversionVCTKFreeVC24,
}

func NewVoiceConversionModel(language Language, dataset Dataset, architecture Architecture) *VoiceConversionModel {
	return &VoiceConversionModel{
		modelType:          ModelTypeVoiceConversion,
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

// GetVoiceConversionModelByArchitecture returns all voice conversion models that use the specified architecture.
func GetVoiceConversionModelByArchitecture(arch Architecture) []VoiceConversionModel {
	return FilterModelsByArchitecture(allVoiceConversions, arch)
}

// GetVoiceConversionModelByDataset returns all voice conversion models of the specified dataset.
func GetVoiceConversionModelByDataset(dataset Dataset) []VoiceConversionModel {
	return FilterModelsByDataset(allVoiceConversions, dataset)
}

// GetVoiceConversionModelByLanguage returns all voice conversion models that support the specified language.
func GetVoiceConversionModelByLanguage(language Language) []VoiceConversionModel {
	return FilterModelsBySupportedLanguages(allVoiceConversions, []Language{language})
}

// GetVoiceConversionModelBySupportedLanguages returns all voice conversion models that support any of the specified languages.
func GetVoiceConversionModelBySupportedLanguages(languages []Language) []VoiceConversionModel {
	return FilterModelsBySupportedLanguages(allVoiceConversions, languages)
}

// GetVoiceConversionModelByDefaultLanguage returns all voice conversion models with the specified default language.
func GetVoiceConversionModelByDefaultLanguage(language Language) []VoiceConversionModel {
	return FilterModelsByDefaultLanguage(allVoiceConversions, language)
}

// GetVoiceConversionModelMultilingual returns all voice conversion models that support multiple languages.
func GetVoiceConversionModelMultilingual() []VoiceConversionModel {
	return FilterModelsMultilingual(allVoiceConversions)
}
