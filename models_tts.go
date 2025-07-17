package coqui

import (
	"fmt"
	"slices"
)

// TTSModel represents a comprehensive model identifier.
type TTSModel struct {
	// Type of the model (e.g., TTS, Vocoder, Voice Conversion).
	Type ModelType
	// Dataset used by the model (e.g., LJSpeech, VCTK, Common Voice).
	Dataset Dataset
	// Architecture of the model (e.g., Tacotron2, VITS, GlowTTS).
	Architecture Architecture

	// supportedLanguages lists all languages this model supports.
	supportedLanguages []Language
	// defaultLanguage is the primary language for this model.
	// There isn't any documentation on which language is the default for each model,
	// so I assume the default language is English, or the first language in the supportedLanguages list.
	defaultLanguage Language
}

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	ModelXTTSv2 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchXTTSv2,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelXTTSv1 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchXTTSv1,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelYourTTS = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchYourTTS,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelBark = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchBark,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	// Common Voice (CV) dataset models
	ModelVITSCV = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCV,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Bulgarian, Czech, Danish, Estonian, Irish, Greek, Croatian, Lithuanian, Latvian, Maltese, Portuguese, Romanian, Slovak, Slovenian, Swedish},
		defaultLanguage:    English, // Fallback to English if available
	}

	// English dataset models
	ModelTacotron2EK1 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetEK1,
		Architecture:       ArchTacotron2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// LJSpeech dataset models
	ModelTacotron2DDCLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelTacotron2DDCPhLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DDCPh,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelGlowTTSLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelSpeedySpeechLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchSpeedySpeech,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelTacotron2DCALJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchTacotron2DCA,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelVITSLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelVITSNeonLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchVITSNeon,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelFastPitchLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchFastPitch,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelOverflowLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchOverflow,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelNeuralHMMLJSpeech = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchNeuralHMM,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// VCTK dataset models
	ModelVITSVCTK = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetVCTK,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelFastPitchVCTK = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetVCTK,
		Architecture:       ArchFastPitch,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Sam dataset models
	ModelTacotron2DDCSam = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetSam,
		Architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Blizzard2013 dataset models
	ModelCapacitronT2C50Blizzard = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBlizzard2013,
		Architecture:       ArchCapacitronT2C50,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelCapacitronT2C150v2Blizzard = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBlizzard2013,
		Architecture:       ArchCapacitronT2C150v2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Multi-dataset English models
	ModelTortoiseV2 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMultiDataset,
		Architecture:       ArchTortoise,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Jenny dataset models
	ModelJenny = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetJenny,
		Architecture:       ArchJenny,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Mai dataset models (multiple languages)
	ModelTacotron2DDCMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{Spanish, French, Dutch},
		defaultLanguage:    Spanish,
	}

	ModelGlowTTSMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Ukrainian},
		defaultLanguage:    Ukrainian,
	}

	ModelVITSMai = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMai,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Ukrainian},
		defaultLanguage:    Ukrainian,
	}

	// CSS10 dataset models (multiple languages)
	ModelVITSCSS10 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCSS10,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Spanish, French, German, Dutch, Hungarian, Finnish},
		defaultLanguage:    Spanish,
	}

	ModelVITSNeonCSS10 = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCSS10,
		Architecture:       ArchVITSNeon,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Baker dataset models (Chinese)
	ModelTacotron2DDCGSTBaker = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetBaker,
		Architecture:       ArchTacotron2DDCGST,
		supportedLanguages: []Language{Chinese},
		defaultLanguage:    Chinese,
	}

	// Thorsten dataset models (German)
	ModelTacotron2DCAThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchTacotron2DCA,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	ModelVITSThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	ModelTacotron2DDCThorsten = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetThorsten,
		Architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Kokoro dataset models (Japanese)
	ModelTacotron2DDCKokoro = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetKokoro,
		Architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{Japanese},
		defaultLanguage:    Japanese,
	}

	// Common Voice dataset models
	ModelGlowTTSCommonVoice = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCommonVoice,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Turkish, Belarusian},
		defaultLanguage:    Turkish,
	}

	// Mai Female dataset models (Italian)
	ModelGlowTTSMaiFemale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiFemale,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	ModelVITSMaiFemale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiFemale,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Italian, Polish},
		defaultLanguage:    Italian,
	}

	// Mai Male dataset models (Italian)
	ModelGlowTTSMaiMale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiMale,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	ModelVITSMaiMale = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetMaiMale,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	// OpenBible dataset models (African languages)
	ModelVITSOpenBible = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetOpenBible,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Ewe, Hausa, Lin, TwiAkuapem, TwiAsante, Yoruba},
		defaultLanguage:    Hausa,
	}

	// Custom dataset models
	ModelVITSCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITS,
		supportedLanguages: []Language{Catalan, Bengali},
		defaultLanguage:    Catalan,
	}

	ModelGlowTTSCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Persian},
		defaultLanguage:    Persian,
	}

	ModelVITSMaleCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITSMale,
		supportedLanguages: []Language{Bengali},
		defaultLanguage:    Bengali,
	}

	ModelVITSFemaleCustom = TTSModel{
		Type:               ModelTypeTTS,
		Dataset:            DatasetCustom,
		Architecture:       ArchVITSFemale,
		supportedLanguages: []Language{Bengali},
		defaultLanguage:    Bengali,
	}
)

// allTTSModels contains a list of all predefined TTS models.
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

// NewTTSModel creates a new TTS model identifier.
// TODO: Implement logic to create a TTS model based on the provided parameters.
// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
func NewTTSModel(language Language, dataset Dataset, architecture Architecture) TTSModel {
	return TTSModel{}
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (t TTSModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", t.GetType(), t.GetDataset(), t.GetDefaultLanguage(), t.GetArchitecture())
}

// IsValid checks if the model identifier is valid.
func (t TTSModel) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), t.Type)
	// Check if the model has a valid type, dataset, architecture, and supported languages.
	return containsValidType && t.Type != "" && t.Dataset != "" && t.Architecture != "" && len(t.supportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (t TTSModel) IsMultilingual() bool {
	return len(t.supportedLanguages) > 1
}

// SupportsLanguage checks if the model supports the specified language.
func (t TTSModel) SupportsLanguage(lang Language) bool {
	return slices.Contains(t.supportedLanguages, lang)
}

// GetType returns the model type.
func (t TTSModel) GetType() ModelType {
	return t.Type
}

// GetArchitecture returns the model architecture.
func (t TTSModel) GetArchitecture() Architecture {
	return t.Architecture
}

// GetDataset returns the model dataset.
func (t TTSModel) GetDataset() Dataset {
	return t.Dataset
}

// GetSupportedLanguages returns the supported languages.
func (t TTSModel) GetSupportedLanguages() []Language {
	return t.supportedLanguages
}

// GetDefaultLanguage returns the default language of the model.
// There isn't any documentation on which language is the default for each model,
// so I assume the default language is English, or the first language in the supportedLanguages list.
func (t TTSModel) GetDefaultLanguage() Language {
	return t.defaultLanguage
}

// GetTTSModelsByArchitecture returns all TTS models that use the specified architecture.
func GetTTSModelsByArchitecture(arch Architecture) []TTSModel {
	return FilterModelsByArchitecture(allTTSModels, arch)
}

// GetTTSModelsByDataset returns all TTS models of the specified type.
func GetTTSModelsByDataset(dataset Dataset) []TTSModel {
	return FilterModelsByDataset(allTTSModels, dataset)
}

// GetTTSModelsByLanguage returns all TTS models that support the specified language (string version).
func GetTTSModelsByLanguage(language Language) []TTSModel {
	return FilterModelsBySupportedLanguages(allTTSModels, []Language{language})
}

// GetTTSModelsBySupportedLanguages returns all TTS models that support any of the specified languages.
func GetTTSModelsBySupportedLanguages(languages []Language) []TTSModel {
	return FilterModelsBySupportedLanguages(allTTSModels, languages)
}

// GetTTSModelsByDefaultLanguage returns all TTS models of the specified type.
func GetTTSModelsByDefaultLanguage(language Language) []TTSModel {
	return FilterModelsByDefaultLanguage(allTTSModels, language)
}

// GetTTSModelsMultilingual returns all TTS models that support multiple languages.
func GetTTSModelsMultilingual() []TTSModel {
	return FilterModelsMultilingual(allTTSModels)
}
