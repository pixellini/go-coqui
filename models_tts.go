package coqui

import (
	"fmt"
	"slices"
)

// TTSModel represents a comprehensive model identifier
type TTSModel struct {
	Type               ModelType
	Dataset            Dataset
	Architecture       Architecture
	SupportedLanguages []Language
	DefaultLanguage    Language
}

// TTS Model definitions organized by dataset and architecture
// Each model supports multiple languages as specified in SupportedLanguages
var (
	// Multilingual models (support all languages)
	ModelXTTSv2 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchXTTSv2,
		SupportedLanguages: GetAllSupportedLanguages(),
		DefaultLanguage:    English,
	}

	ModelXTTSv1 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchXTTSv1,
		SupportedLanguages: GetAllSupportedLanguages(),
		DefaultLanguage:    English,
	}

	ModelYourTTS = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchYourTTS,
		SupportedLanguages: GetAllSupportedLanguages(),
		DefaultLanguage:    English,
	}

	ModelBark = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchBark,
		SupportedLanguages: GetAllSupportedLanguages(),
		DefaultLanguage:    English,
	}

	// Common Voice (CV) dataset models
	ModelVITSCV = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCV,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Bulgarian, Czech, Danish, Estonian, Irish, Greek, Croatian, Lithuanian, Latvian, Maltese, Portuguese, Romanian, Slovak, Slovenian, Swedish},
		DefaultLanguage:    English, // Fallback to English if available
	}

	// English dataset models
	ModelTacotron2EK1 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetEK1,
		Architecture:       ArchTacotron2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// LJSpeech dataset models
	ModelTacotron2DDCLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DDC,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelTacotron2DDCPhLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DDCPh,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelGlowTTSLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelSpeedySpeechLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchSpeedySpeech,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelTacotron2DCALJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DCA,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelVITSLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelVITSNeonLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchVITSNeon,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelFastPitchLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchFastPitch,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelOverflowLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchOverflow,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelNeuralHMMLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchNeuralHMM,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// VCTK dataset models
	ModelVITSVCTK = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetVCTK,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelFastPitchVCTK = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetVCTK,
		Architecture:       ArchFastPitch,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Sam dataset models
	ModelTacotron2DDCSam = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetSam,
		Architecture:       ArchTacotron2DDC,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Blizzard2013 dataset models
	ModelCapacitronT2C50Blizzard = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBlizzard2013,
		Architecture:       ArchCapacitronT2C50,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	ModelCapacitronT2C150v2Blizzard = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBlizzard2013,
		Architecture:       ArchCapacitronT2C150v2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Multi-dataset English models
	ModelTortoiseV2 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchTortoise,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Jenny dataset models
	ModelJenny = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetJenny,
		Architecture:       ArchJenny,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Mai dataset models (multiple languages)
	ModelTacotron2DDCMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchTacotron2DDC,
		SupportedLanguages: []Language{Spanish, French, Dutch},
		DefaultLanguage:    Spanish,
	}

	ModelGlowTTSMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{Ukrainian},
		DefaultLanguage:    Ukrainian,
	}

	ModelVITSMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Ukrainian},
		DefaultLanguage:    Ukrainian,
	}

	// CSS10 dataset models (multiple languages)
	ModelVITSCSS10 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCSS10,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Spanish, French, German, Dutch, Hungarian, Finnish},
		DefaultLanguage:    Spanish,
	}

	ModelVITSNeonCSS10 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCSS10,
		Architecture:       ArchVITSNeon,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	// Baker dataset models (Chinese)
	ModelTacotron2DDCGSTBaker = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBaker,
		Architecture:       ArchTacotron2DDCGST,
		SupportedLanguages: []Language{Chinese},
		DefaultLanguage:    Chinese,
	}

	// Thorsten dataset models (German)
	ModelTacotron2DCAThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchTacotron2DCA,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	ModelVITSThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	ModelTacotron2DDCThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchTacotron2DDC,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	// Kokoro dataset models (Japanese)
	ModelTacotron2DDCKokoro = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetKokoro,
		Architecture:       ArchTacotron2DDC,
		SupportedLanguages: []Language{Japanese},
		DefaultLanguage:    Japanese,
	}

	// Common Voice dataset models
	ModelGlowTTSCommonVoice = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCommonVoice,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{Turkish, Belarusian},
		DefaultLanguage:    Turkish,
	}

	// Mai Female dataset models (Italian)
	ModelGlowTTSMaiFemale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiFemale,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{Italian},
		DefaultLanguage:    Italian,
	}

	ModelVITSMaiFemale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiFemale,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Italian, Polish},
		DefaultLanguage:    Italian,
	}

	// Mai Male dataset models (Italian)
	ModelGlowTTSMaiMale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiMale,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{Italian},
		DefaultLanguage:    Italian,
	}

	ModelVITSMaiMale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiMale,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Italian},
		DefaultLanguage:    Italian,
	}

	// OpenBible dataset models (African languages)
	ModelVITSOpenBible = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetOpenBible,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Ewe, Hausa, Lin, TwiAkuapem, TwiAsante, Yoruba},
		DefaultLanguage:    Hausa,
	}

	// Custom dataset models
	ModelVITSCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITS,
		SupportedLanguages: []Language{Catalan, Bengali},
		DefaultLanguage:    Catalan,
	}

	ModelGlowTTSCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchGlowTTS,
		SupportedLanguages: []Language{Persian},
		DefaultLanguage:    Persian,
	}

	ModelVITSMaleCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITSMale,
		SupportedLanguages: []Language{Bengali},
		DefaultLanguage:    Bengali,
	}

	ModelVITSFemaleCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITSFemale,
		SupportedLanguages: []Language{Bengali},
		DefaultLanguage:    Bengali,
	}
)

// allTTSModels contains all predefined TTS model identifiers as pointers
var allTTSModels = []TTSModel{
	// Multilingual models
	ModelXTTSv2,
	ModelXTTSv1,
	ModelYourTTS,
	ModelBark,

	// Common Voice models
	ModelVITSCV,

	// English models
	ModelTacotron2EK1,
	ModelTacotron2DDCLJSpeech,
	ModelTacotron2DDCPhLJSpeech,
	ModelGlowTTSLJSpeech,
	ModelSpeedySpeechLJSpeech,
	ModelTacotron2DCALJSpeech,
	ModelVITSLJSpeech,
	ModelVITSNeonLJSpeech,
	ModelFastPitchLJSpeech,
	ModelOverflowLJSpeech,
	ModelNeuralHMMLJSpeech,
	ModelVITSVCTK,
	ModelFastPitchVCTK,
	ModelTacotron2DDCSam,
	ModelCapacitronT2C50Blizzard,
	ModelCapacitronT2C150v2Blizzard,
	ModelTortoiseV2,
	ModelJenny,

	// Multi-language models
	ModelTacotron2DDCMai,
	ModelGlowTTSMai,
	ModelVITSMai,
	ModelVITSCSS10,
	ModelVITSNeonCSS10,

	// Language-specific models
	ModelTacotron2DDCGSTBaker,
	ModelTacotron2DCAThorsten,
	ModelVITSThorsten,
	ModelTacotron2DDCThorsten,
	ModelTacotron2DDCKokoro,
	ModelGlowTTSCommonVoice,
	ModelGlowTTSMaiFemale,
	ModelVITSMaiFemale,
	ModelGlowTTSMaiMale,
	ModelVITSMaiMale,
	ModelVITSOpenBible,
	ModelVITSCustom,
	ModelGlowTTSCustom,
	ModelVITSMaleCustom,
	ModelVITSFemaleCustom,
}

// GetType returns the model type
func (t TTSModel) GetType() ModelType {
	return t.Type
}

// GetArchitecture returns the model architecture
func (t TTSModel) GetArchitecture() Architecture {
	return t.Architecture
}

// GetDataset returns the model dataset
func (t TTSModel) GetDataset() Dataset {
	return t.Dataset
}

// GetSupportedLanguages returns the supported languages
func (t TTSModel) GetSupportedLanguages() []Language {
	return t.SupportedLanguages
}

// GetDefaultLanguage returns the default language
func (t TTSModel) GetDefaultLanguage() Language {
	return t.DefaultLanguage
}

// IsValid checks if the model identifier is valid
func (t TTSModel) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), t.Type)
	// Check if the model has a valid type, dataset, architecture, and supported languages
	return containsValidType && t.Type != "" && t.Dataset != "" && t.Architecture != "" && len(t.SupportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (t TTSModel) IsMultilingual() bool {
	return len(t.SupportedLanguages) > 1
}

func (t TTSModel) SupportsLanguage(lang Language) bool {
	for _, supportedLang := range t.SupportedLanguages {
		if supportedLang == lang {
			return true
		}
	}
	return false
}

// String returns a string representation of the model identifier
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (t TTSModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", t.GetType(), t.GetDataset(), t.GetDefaultLanguage(), t.GetArchitecture())
}

// NewTTSModel creates a new TTS model identifier
func NewTTSModel(language, dataset string, architecture Architecture) TTSModel {
	return TTSModel{}
}

// GetTTSModelsByArchitecture returns all TTS models that use the specified architecture
func GetTTSModelsByArchitecture(arch Architecture) []TTSModel {
	return FilterModelsByArchitecture(allTTSModels, arch)
}

// GetTTSModelsByDataset returns all TTS models of the specified type
func GetTTSModelsByDataset(dataset Dataset) []TTSModel {
	return FilterModelsByDataset(allTTSModels, dataset)
}

// GetTTSModelsByLanguage returns all TTS models that support the specified language (string version)
func GetTTSModelsByLanguage(language Language) []TTSModel {
	return FilterModelsBySupportedLanguages(allTTSModels, []Language{language})
}

// GetTTSModelsBySupportedLanguages returns all TTS models that support any of the specified languages
func GetTTSModelsBySupportedLanguages(languages []Language) []TTSModel {
	return FilterModelsBySupportedLanguages(allTTSModels, languages)
}

// GetTTSModelsByDefaultLanguage returns all TTS models of the specified type
func GetTTSModelsByDefaultLanguage(language Language) []TTSModel {
	return FilterModelsByDefaultLanguage(allTTSModels, language)
}

// GetTTSModelsMultilingual returns all TTS models that support multiple languages
func GetTTSModelsMultilingual() []TTSModel {
	return FilterModelsMultilingual(allTTSModels)
}
