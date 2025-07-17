package coqui

import (
	"fmt"
	"slices"
)

// TTSModel represents a comprehensive model identifier.
type TTSModel struct {
	// type of the model (e.g., TTS, Vocoder, Voice Conversion).
	modelType ModelType
	// dataset used by the model (e.g., LJSpeech, VCTK, Common Voice).
	dataset Dataset
	// architecture of the model (e.g., Tacotron2, VITS, GlowTTS).
	architecture Architecture

	// supportedLanguages lists all languages this model supports.
	supportedLanguages []Language
	// defaultLanguage is the primary language for this model.
	// There isn't any documentation on which language is the default for each model,
	// so I assume the default language is English, or the first language in the supportedLanguages list.
	defaultLanguage Language

	// isCustom Indicates if this is a custom model not predefined in the library
	isCustom bool
}

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	ModelXTTSv2 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchXTTSv2,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelXTTSv1 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchXTTSv1,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelYourTTS = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchYourTTS,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	ModelBark = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchBark,
		supportedLanguages: allSupportedLanguages,
		defaultLanguage:    English,
	}

	// Common Voice (CV) dataset models
	ModelVITSCV = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCV,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Bulgarian, Czech, Danish, Estonian, Irish, Greek, Croatian, Lithuanian, Latvian, Maltese, Portuguese, Romanian, Slovak, Slovenian, Swedish},
		defaultLanguage:    English, // Fallback to English if available
	}

	// English dataset models
	ModelTacotron2EK1 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetEK1,
		architecture:       ArchTacotron2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// LJSpeech dataset models
	ModelTacotron2DDCLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelTacotron2DDCPhLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DDCPh,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelGlowTTSLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelSpeedySpeechLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchSpeedySpeech,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelTacotron2DCALJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DCA,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelVITSLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchVITS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelVITSNeonLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchVITSNeon,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelFastPitchLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchFastPitch,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelOverflowLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchOverflow,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelNeuralHMMLJSpeech = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchNeuralHMM,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// VCTK dataset models
	ModelVITSVCTK = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetVCTK,
		architecture:       ArchVITS,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelFastPitchVCTK = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetVCTK,
		architecture:       ArchFastPitch,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Sam dataset models
	ModelTacotron2DDCSam = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetSam,
		architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Blizzard2013 dataset models
	ModelCapacitronT2C50Blizzard = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchCapacitronT2C50,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	ModelCapacitronT2C150v2Blizzard = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchCapacitronT2C150v2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Multi-dataset English models
	ModelTortoiseV2 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchTortoise,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Jenny dataset models
	ModelJenny = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetJenny,
		architecture:       ArchJenny,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Mai dataset models (multiple languages)
	ModelTacotron2DDCMai = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{Spanish, French, Dutch},
		defaultLanguage:    Spanish,
	}

	ModelGlowTTSMai = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Ukrainian},
		defaultLanguage:    Ukrainian,
	}

	ModelVITSMai = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Ukrainian},
		defaultLanguage:    Ukrainian,
	}

	// CSS10 dataset models (multiple languages)
	ModelVITSCSS10 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCSS10,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Spanish, French, German, Dutch, Hungarian, Finnish},
		defaultLanguage:    Spanish,
	}

	ModelVITSNeonCSS10 = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCSS10,
		architecture:       ArchVITSNeon,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Baker dataset models (Chinese)
	ModelTacotron2DDCGSTBaker = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetBaker,
		architecture:       ArchTacotron2DDCGST,
		supportedLanguages: []Language{Chinese},
		defaultLanguage:    Chinese,
	}

	// Thorsten dataset models (German)
	ModelTacotron2DCAThorsten = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchTacotron2DCA,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	ModelVITSThorsten = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchVITS,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	ModelTacotron2DDCThorsten = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Kokoro dataset models (Japanese)
	ModelTacotron2DDCKokoro = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetKokoro,
		architecture:       ArchTacotron2DDC,
		supportedLanguages: []Language{Japanese},
		defaultLanguage:    Japanese,
	}

	// Common Voice dataset models
	ModelGlowTTSCommonVoice = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCommonVoice,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Turkish, Belarusian},
		defaultLanguage:    Turkish,
	}

	// Mai Female dataset models (Italian)
	ModelGlowTTSMaiFemale = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMaiFemale,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	ModelVITSMaiFemale = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMaiFemale,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Italian, Polish},
		defaultLanguage:    Italian,
	}

	// Mai Male dataset models (Italian)
	ModelGlowTTSMaiMale = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMaiMale,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	ModelVITSMaiMale = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetMaiMale,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Italian},
		defaultLanguage:    Italian,
	}

	// OpenBible dataset models (African languages)
	ModelVITSOpenBible = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetOpenBible,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Ewe, Hausa, Lin, TwiAkuapem, TwiAsante, Yoruba},
		defaultLanguage:    Hausa,
	}

	// Custom dataset models
	ModelVITSCustom = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITS,
		supportedLanguages: []Language{Catalan, Bengali},
		defaultLanguage:    Catalan,
	}

	ModelGlowTTSCustom = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchGlowTTS,
		supportedLanguages: []Language{Persian},
		defaultLanguage:    Persian,
	}

	ModelVITSMaleCustom = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITSMale,
		supportedLanguages: []Language{Bengali},
		defaultLanguage:    Bengali,
	}

	ModelVITSFemaleCustom = TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITSFemale,
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
	return TTSModel{
		modelType:          ModelTypeTTS,
		dataset:            dataset,
		architecture:       architecture,
		supportedLanguages: []Language{language},
		defaultLanguage:    language,
	}
}

// NewCustomTTSModel creates a new custom TTS model identifier.
// This is useful for models that are not predefined in the Coqui TTS library.
func NewCustomTTSModel(language Language, dataset Dataset, architecture Architecture) TTSModel {
	model := NewTTSModel(language, dataset, architecture)
	model.isCustom = true // Mark this model as custom
	return model
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (t TTSModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", t.GetModelType(), t.GetDataset(), t.GetDefaultLanguage(), t.GetArchitecture())
}

// IsValid checks if the model identifier is valid.
func (t TTSModel) IsValid() bool {
	// Assume it's valid if it's a custom model provided by the user.
	if t.isCustom {
		return true
	}

	containsValidType := slices.Contains(GetAllModelTypes(), t.modelType)
	// Check if the model has a valid type, dataset, architecture, and supported languages.
	return containsValidType && t.modelType != "" && t.dataset != "" && t.architecture != "" && len(t.supportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (t TTSModel) IsMultilingual() bool {
	return len(t.supportedLanguages) > 1
}

// SupportsLanguage checks if the model supports the specified language.
func (t TTSModel) SupportsLanguage(lang Language) bool {
	return slices.Contains(t.supportedLanguages, lang)
}

// GetModelType returns the model type.
func (t TTSModel) GetModelType() ModelType {
	return t.modelType
}

// GetArchitecture returns the model architecture.
func (t TTSModel) GetArchitecture() Architecture {
	return t.architecture
}

// GetDataset returns the model dataset.
func (t TTSModel) GetDataset() Dataset {
	return t.dataset
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
