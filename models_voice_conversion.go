package coqui

import (
	"fmt"
	"slices"
)

// VoiceConversion represents a voice conversion model identifier
type VoiceConversion struct {
	Type               ModelType
	Dataset            Dataset
	Architecture       Architecture
	SupportedLanguages []Language
	DefaultLanguage    Language
}

// Voice Conversion Models for converting one voice to another
var (
	// Multilingual voice conversion
	VoiceConversionVCTKFreeVC24 = VoiceConversion{
		Type:               ModelTypeVoiceConversion,
		SupportedLanguages: GetAllSupportedLanguages(),
		Dataset:            DatasetVCTK,
		Architecture:       ArchFreevc24,
		DefaultLanguage:    English,
	}
)

// allVoiceConversions contains all predefined voice conversion model identifiers as pointers
var allVoiceConversions = []*VoiceConversion{
	&VoiceConversionVCTKFreeVC24,
}

// GetType returns the model type
func (vc *VoiceConversion) GetType() ModelType {
	return vc.Type
}

// GetArchitecture returns the model architecture
func (vc *VoiceConversion) GetArchitecture() Architecture {
	return vc.Architecture
}

// GetDataset returns the model dataset
func (vc *VoiceConversion) GetDataset() Dataset {
	return vc.Dataset
}

// GetSupportedLanguages returns the supported languages
func (vc *VoiceConversion) GetSupportedLanguages() []Language {
	return vc.SupportedLanguages
}

// GetDefaultLanguage returns the default language
func (vc *VoiceConversion) GetDefaultLanguage() Language {
	return vc.DefaultLanguage
}

// IsValid checks if the model identifier is valid
func (vc *VoiceConversion) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), vc.Type)
	// Check if the model has a valid type, dataset, architecture, and supported languages
	return containsValidType && vc.Type != "" && vc.Dataset != "" && vc.Architecture != "" && len(vc.SupportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (vc *VoiceConversion) IsMultilingual() bool {
	return len(vc.SupportedLanguages) > 1
}

// String returns a string representation of the model identifier
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (vc *VoiceConversion) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", vc.GetType(), vc.GetDataset(), vc.GetDefaultLanguage(), vc.GetArchitecture())
}

// NewVoiceConversionModel creates a new vocoder model identifier
func NewVoiceConversionModel(language, dataset string, architecture Architecture) VoiceConversion {
	return VoiceConversion{}
}

// GetVoiceConversionModelByArchitecture returns all voice conversion models that use the specified architecture
func GetVoiceConversionModelByArchitecture(arch Architecture) []*VoiceConversion {
	return FilterModelsByArchitecture(allVoiceConversions, arch)
}

// GetVoiceConversionModelByDataset returns all voice conversion models of the specified dataset
func GetVoiceConversionModelByDataset(dataset Dataset) []*VoiceConversion {
	return FilterModelsByDataset(allVoiceConversions, dataset)
}

// GetVoiceConversionModelByLanguage returns all voice conversion models that support the specified language
func GetVoiceConversionModelByLanguage(language Language) []*VoiceConversion {
	return FilterModelsBySupportedLanguages(allVoiceConversions, []Language{language})
}

// GetVoiceConversionModelBySupportedLanguages returns all voice conversion models that support any of the specified languages
func GetVoiceConversionModelBySupportedLanguages(languages []Language) []*VoiceConversion {
	return FilterModelsBySupportedLanguages(allVoiceConversions, languages)
}

// GetVoiceConversionModelByDefaultLanguage returns all voice conversion models with the specified default language
func GetVoiceConversionModelByDefaultLanguage(language Language) []*VoiceConversion {
	return FilterModelsByDefaultLanguage(allVoiceConversions, language)
}

// GetVoiceConversionModelMultilingual returns all voice conversion models that support multiple languages
func GetVoiceConversionModelMultilingual() []*VoiceConversion {
	return FilterModelsMultilingual(allVoiceConversions)
}
