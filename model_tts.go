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
	// defaultLanguage is the primary language for this model.
	// There isn't any documentation on which language is the default for each model,
	// so I assume the default language is English, or the first language in the supportedLanguages list.
	defaultLanguage Language
	// supportedLanguages lists all languages this model supports.
	supportedLanguages []Language
	// supportsVoiceCloning indicates if the model supports voice cloning by providing a speaker sample.
	supportsVoiceCloning bool
	// isCustom Indicates if this is a custom model not predefined in the library
	isCustom bool
}

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	ModelXTTSv2 = TTSModel{
		modelType:            modelTypeTTS,
		dataset:              DatasetMultiDataset,
		architecture:         ArchXTTSv2,
		defaultLanguage:      English,
		supportedLanguages:   allSupportedLanguages,
		supportsVoiceCloning: true,
	}

	ModelXTTSv1 = TTSModel{
		modelType:            modelTypeTTS,
		dataset:              DatasetMultiDataset,
		architecture:         ArchXTTSv1,
		defaultLanguage:      English,
		supportedLanguages:   allSupportedLanguages,
		supportsVoiceCloning: true,
	}

	ModelYourTTS = TTSModel{
		modelType:            modelTypeTTS,
		dataset:              DatasetMultiDataset,
		architecture:         ArchYourTTS,
		defaultLanguage:      English,
		supportedLanguages:   allSupportedLanguages,
		supportsVoiceCloning: true,
	}

	ModelBark = TTSModel{
		modelType:            modelTypeTTS,
		dataset:              DatasetMultiDataset,
		architecture:         ArchBark,
		defaultLanguage:      English,
		supportedLanguages:   allSupportedLanguages,
		supportsVoiceCloning: true,
	}

	// Common Voice (CV) dataset models
	ModelVITSCV = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCV,
		architecture:       ArchVITS,
		defaultLanguage:    English, // Fallback to English if available
		supportedLanguages: []Language{Bulgarian, Czech, Danish, Estonian, Irish, Greek, Croatian, Lithuanian, Latvian, Maltese, Portuguese, Romanian, Slovak, Slovenian, Swedish},
	}

	// English dataset models
	ModelTacotron2EK1 = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetEK1,
		architecture:       ArchTacotron2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// LJSpeech dataset models
	ModelTacotron2DDCLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DDC,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelTacotron2DDCPhLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DDCPh,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelGlowTTSLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchGlowTTS,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelSpeedySpeechLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchSpeedySpeech,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelTacotron2DCALJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchTacotron2DCA,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelVITSLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchVITS,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelVITSNeonLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchVITSNeon,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelFastPitchLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchFastPitch,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelOverflowLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchOverflow,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelNeuralHMMLJSpeech = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetLJSpeech,
		architecture:       ArchNeuralHMM,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// VCTK dataset models
	ModelVITSVCTK = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetVCTK,
		architecture:       ArchVITS,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelFastPitchVCTK = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetVCTK,
		architecture:       ArchFastPitch,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Sam dataset models
	ModelTacotron2DDCSam = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetSam,
		architecture:       ArchTacotron2DDC,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Blizzard2013 dataset models
	ModelCapacitronT2C50Blizzard = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchCapacitronT2C50,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	ModelCapacitronT2C150v2Blizzard = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchCapacitronT2C150v2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Multi-dataset English models
	ModelTortoiseV2 = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMultiDataset,
		architecture:       ArchTortoise,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Jenny dataset models
	ModelJenny = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetJenny,
		architecture:       ArchJenny,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Mai dataset models (multiple languages)
	ModelTacotron2DDCMai = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchTacotron2DDC,
		defaultLanguage:    Spanish,
		supportedLanguages: []Language{Spanish, French, Dutch},
	}

	ModelGlowTTSMai = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchGlowTTS,
		defaultLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	ModelVITSMai = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMai,
		architecture:       ArchVITS,
		defaultLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	// CSS10 dataset models (multiple languages)
	ModelVITSCSS10 = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCSS10,
		architecture:       ArchVITS,
		defaultLanguage:    Spanish,
		supportedLanguages: []Language{Spanish, French, German, Dutch, Hungarian, Finnish},
	}

	ModelVITSNeonCSS10 = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCSS10,
		architecture:       ArchVITSNeon,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Baker dataset models (Chinese)
	ModelTacotron2DDCGSTBaker = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetBaker,
		architecture:       ArchTacotron2DDCGST,
		defaultLanguage:    Chinese,
		supportedLanguages: []Language{Chinese},
	}

	// Thorsten dataset models (German)
	ModelTacotron2DCAThorsten = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchTacotron2DCA,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	ModelVITSThorsten = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchVITS,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	ModelTacotron2DDCThorsten = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetThorsten,
		architecture:       ArchTacotron2DDC,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Kokoro dataset models (Japanese)
	ModelTacotron2DDCKokoro = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetKokoro,
		architecture:       ArchTacotron2DDC,
		defaultLanguage:    Japanese,
		supportedLanguages: []Language{Japanese},
	}

	// Common Voice dataset models
	ModelGlowTTSCommonVoice = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCommonVoice,
		architecture:       ArchGlowTTS,
		defaultLanguage:    Turkish,
		supportedLanguages: []Language{Turkish, Belarusian},
	}

	// Mai Female dataset models (Italian)
	ModelGlowTTSMaiFemale = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMaiFemale,
		architecture:       ArchGlowTTS,
		defaultLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	ModelVITSMaiFemale = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMaiFemale,
		architecture:       ArchVITS,
		defaultLanguage:    Italian,
		supportedLanguages: []Language{Italian, Polish},
	}

	// Mai Male dataset models (Italian)
	ModelGlowTTSMaiMale = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMaiMale,
		architecture:       ArchGlowTTS,
		defaultLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	ModelVITSMaiMale = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetMaiMale,
		architecture:       ArchVITS,
		defaultLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	// OpenBible dataset models (African languages)
	ModelVITSOpenBible = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetOpenBible,
		architecture:       ArchVITS,
		defaultLanguage:    Hausa,
		supportedLanguages: []Language{Ewe, Hausa, Lin, TwiAkuapem, TwiAsante, Yoruba},
	}

	// Custom dataset models
	ModelVITSCustom = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITS,
		defaultLanguage:    Catalan,
		supportedLanguages: []Language{Catalan, Bengali},
	}

	ModelGlowTTSCustom = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchGlowTTS,
		defaultLanguage:    Persian,
		supportedLanguages: []Language{Persian},
	}

	ModelVITSMaleCustom = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITSMale,
		defaultLanguage:    Bengali,
		supportedLanguages: []Language{Bengali},
	}

	ModelVITSFemaleCustom = TTSModel{
		modelType:          modelTypeTTS,
		dataset:            DatasetCustom,
		architecture:       ArchVITSFemale,
		defaultLanguage:    Bengali,
		supportedLanguages: []Language{Bengali},
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

// NewModel creates a new custom TTS model identifier.
// This is useful for models that are not predefined in the Coqui TTS library.
func NewModel(language Language, dataset Dataset, architecture Architecture) TTSModel {
	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return TTSModel{
		modelType:          modelTypeTTS,
		dataset:            dataset,
		architecture:       architecture,
		defaultLanguage:    language,
		supportedLanguages: []Language{language},
		isCustom:           true,
	}
}

// String returns a string representation of the model identifier.
// It will only print the default language.
func (t TTSModel) String() string {
	// Will need to be updated if we want to include all supported languages.
	return fmt.Sprintf("%s/%s/%s/%s", t.GetModelType(), t.GetDataset(), t.GetDefaultLanguage(), t.GetArchitecture())
}

// NameList returns a list of string representations of the model identifier for each supported language.
func (t TTSModel) NameList() []string {
	var names []string
	for _, lang := range t.supportedLanguages {
		name := fmt.Sprintf("%s/%s/%s/%s", t.GetModelType(), t.GetDataset(), lang, t.GetArchitecture())
		names = append(names, name)
	}
	return names
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

// SupportsVoiceCloning checks if the model supports voice cloning by providing a speaker sample.
func (t TTSModel) SupportsVoiceCloning() bool {
	return t.supportsVoiceCloning
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

// GetAllModels returns a copy of the list of all predefined TTS models.
func GetAllModels() []TTSModel {
	return slices.Clone(allTTSModels)
}

// GetModelsByArchitecture returns all TTS models that use the specified architecture.
func GetModelsByArchitecture(arch Architecture) []TTSModel {
	return filterModelsByArchitecture(allTTSModels, arch)
}

// GetModelsByDataset returns all TTS models of the specified type.
func GetModelsByDataset(dataset Dataset) []TTSModel {
	return filterModelsByDataset(allTTSModels, dataset)
}

// GetModelsByLanguage returns all TTS models that support the specified language (string version).
func GetModelsByLanguage(language Language) []TTSModel {
	return filterModelsBySupportedLanguages(allTTSModels, []Language{language})
}

// GetModelsBySupportedLanguages returns all TTS models that support any of the specified languages.
func GetModelsBySupportedLanguages(languages []Language) []TTSModel {
	return filterModelsBySupportedLanguages(allTTSModels, languages)
}

// GetModelsByDefaultLanguage returns all TTS models of the specified type.
func GetModelsByDefaultLanguage(language Language) []TTSModel {
	return filterModelsByDefaultLanguage(allTTSModels, language)
}

// GetModelsMultilingual returns all TTS models that support multiple languages.
func GetModelsMultilingual() []TTSModel {
	return filterModelsMultilingual(allTTSModels)
}
