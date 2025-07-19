package coqui

import "slices"

type TTSModel = ModelIdentifier

const (
	// Multilingual architectures.
	BaseModelXTTSv2  BaseModel = "xtts_v2"
	BaseModelXTTSv1  BaseModel = "xtts_v1.1"
	BaseModelYourTTS BaseModel = "your_tts"
	BaseModelBark    BaseModel = "bark"

	// Single language architectures.
	BaseModelVITS         BaseModel = "vits"
	BaseModelVITSNeon     BaseModel = "vits--neon"
	BaseModelVITSNeonDash BaseModel = "vits-neon"
	BaseModelVITSMale     BaseModel = "vits-male"
	BaseModelVITSFemale   BaseModel = "vits-female"

	// Tacotron variants.
	BaseModelTacotron2       BaseModel = "tacotron2"
	BaseModelTacotron2DDC    BaseModel = "tacotron2-DDC"
	BaseModelTacotron2DDCPh  BaseModel = "tacotron2-DDC_ph"
	BaseModelTacotron2DCA    BaseModel = "tacotron2-DCA"
	BaseModelTacotron2DDCGST BaseModel = "tacotron2-DDC-GST"

	// Other architectures.
	BaseModelGlowTTS            BaseModel = "glow-tts"
	BaseModelFastPitch          BaseModel = "fast_pitch"
	BaseModelSpeedySpeech       BaseModel = "speedy-speech"
	BaseModelOverflow           BaseModel = "overflow"
	BaseModelNeuralHMM          BaseModel = "neural_hmm"
	BaseModelTortoise           BaseModel = "tortoise-v2"
	BaseModelCapacitronT2C50    BaseModel = "capacitron-t2-c50"
	BaseModelCapacitronT2C150v2 BaseModel = "capacitron-t2-c150_v2"
	BaseModelJenny              BaseModel = "jenny"
)

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	TTSModelXTTSv2 = TTSModel{
		category:             modelTypeTTS,
		dataset:              DatasetMultiDataset,
		model:                BaseModelXTTSv2,
		defaultLanguage:      English,
		currentLanguage:      English,
		supportedLanguages:   supportedLanguages,
		supportsVoiceCloning: true,
	}

	TTSModelXTTSv1 = TTSModel{
		category:             modelTypeTTS,
		dataset:              DatasetMultiDataset,
		model:                BaseModelXTTSv1,
		defaultLanguage:      English,
		currentLanguage:      English,
		supportedLanguages:   supportedLanguages,
		supportsVoiceCloning: true,
	}

	TTSModelYourTTS = TTSModel{
		category:             modelTypeTTS,
		dataset:              DatasetMultiDataset,
		model:                BaseModelYourTTS,
		defaultLanguage:      English,
		currentLanguage:      English,
		supportedLanguages:   supportedLanguages,
		supportsVoiceCloning: true,
	}

	TTSModelBark = TTSModel{
		category:             modelTypeTTS,
		dataset:              DatasetMultiDataset,
		model:                BaseModelBark,
		defaultLanguage:      English,
		currentLanguage:      English,
		supportedLanguages:   supportedLanguages,
		supportsVoiceCloning: true,
	}

	// Common Voice (CV) dataset models
	TTSModelVITSCV = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCV,
		model:              BaseModelVITS,
		defaultLanguage:    English, // Fallback to English if available
		currentLanguage:    English, // Fallback to English if available
		supportedLanguages: []Language{Bulgarian, Czech, Danish, Estonian, Irish, Greek, Croatian, Lithuanian, Latvian, Maltese, Portuguese, Romanian, Slovak, Slovenian, Swedish},
	}

	// English dataset models
	TTSModelTacotron2EK1 = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetEK1,
		model:              BaseModelTacotron2,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// LJSpeech dataset models
	TTSModelTacotron2DDCLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelTacotron2DDCPhLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DDCPh,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelGlowTTSLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelGlowTTS,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelSpeedySpeechLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelSpeedySpeech,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelTacotron2DCALJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelTacotron2DCA,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelVITSLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelVITS,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelVITSNeonLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelVITSNeon,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelFastPitchLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelFastPitch,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelOverflowLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelOverflow,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelNeuralHMMLJSpeech = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetLJSpeech,
		model:              BaseModelNeuralHMM,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// VCTK dataset models
	TTSModelVITSVCTK = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetVCTK,
		model:              BaseModelVITS,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelFastPitchVCTK = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetVCTK,
		model:              BaseModelFastPitch,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Sam dataset models
	TTSModelTacotron2DDCSam = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetSam,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Blizzard2013 dataset models
	TTSModelCapacitronT2C50Blizzard = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetBlizzard2013,
		model:              BaseModelCapacitronT2C50,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	TTSModelCapacitronT2C150v2Blizzard = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetBlizzard2013,
		model:              BaseModelCapacitronT2C150v2,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Multi-dataset English models
	TTSModelTortoiseV2 = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMultiDataset,
		model:              BaseModelTortoise,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Jenny dataset models
	TTSModelJenny = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetJenny,
		model:              BaseModelJenny,
		defaultLanguage:    English,
		currentLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Mai dataset models (multiple languages)
	TTSModelTacotron2DDCMai = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMai,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    Spanish,
		currentLanguage:    Spanish,
		supportedLanguages: []Language{Spanish, French, Dutch},
	}

	TTSModelGlowTTSMai = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMai,
		model:              BaseModelGlowTTS,
		defaultLanguage:    Ukrainian,
		currentLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	TTSModelVITSMai = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMai,
		model:              BaseModelVITS,
		defaultLanguage:    Ukrainian,
		currentLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	// CSS10 dataset models (multiple languages)
	TTSModelVITSCSS10 = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCSS10,
		model:              BaseModelVITS,
		defaultLanguage:    Spanish,
		currentLanguage:    Spanish,
		supportedLanguages: []Language{Spanish, French, German, Dutch, Hungarian, Finnish},
	}

	TTSModelVITSNeonCSS10 = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCSS10,
		model:              BaseModelVITSNeon,
		defaultLanguage:    German,
		currentLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Baker dataset models (Chinese)
	TTSModelTacotron2DDCGSTBaker = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetBaker,
		model:              BaseModelTacotron2DDCGST,
		defaultLanguage:    Chinese,
		currentLanguage:    Chinese,
		supportedLanguages: []Language{Chinese},
	}

	// Thorsten dataset models (German)
	TTSModelTacotron2DCAThorsten = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetThorsten,
		model:              BaseModelTacotron2DCA,
		defaultLanguage:    German,
		currentLanguage:    German,
		supportedLanguages: []Language{German},
	}

	TTSModelVITSThorsten = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetThorsten,
		model:              BaseModelVITS,
		defaultLanguage:    German,
		currentLanguage:    German,
		supportedLanguages: []Language{German},
	}

	TTSModelTacotron2DDCThorsten = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetThorsten,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    German,
		currentLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Kokoro dataset models (Japanese)
	TTSModelTacotron2DDCKokoro = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetKokoro,
		model:              BaseModelTacotron2DDC,
		defaultLanguage:    Japanese,
		currentLanguage:    Japanese,
		supportedLanguages: []Language{Japanese},
	}

	// Common Voice dataset models
	TTSModelGlowTTSCommonVoice = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCommonVoice,
		model:              BaseModelGlowTTS,
		defaultLanguage:    Turkish,
		currentLanguage:    Turkish,
		supportedLanguages: []Language{Turkish, Belarusian},
	}

	// Mai Female dataset models (Italian)
	TTSModelGlowTTSMaiFemale = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMaiFemale,
		model:              BaseModelGlowTTS,
		defaultLanguage:    Italian,
		currentLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	TTSModelVITSMaiFemale = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMaiFemale,
		model:              BaseModelVITS,
		defaultLanguage:    Italian,
		currentLanguage:    Italian,
		supportedLanguages: []Language{Italian, Polish},
	}

	// Mai Male dataset models (Italian)
	TTSModelGlowTTSMaiMale = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMaiMale,
		model:              BaseModelGlowTTS,
		defaultLanguage:    Italian,
		currentLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	TTSModelVITSMaiMale = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetMaiMale,
		model:              BaseModelVITS,
		defaultLanguage:    Italian,
		currentLanguage:    Italian,
		supportedLanguages: []Language{Italian},
	}

	// OpenBible dataset models (African languages)
	TTSModelVITSOpenBible = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetOpenBible,
		model:              BaseModelVITS,
		defaultLanguage:    Hausa,
		currentLanguage:    Hausa,
		supportedLanguages: []Language{Ewe, Hausa, Lin, TwiAkuapem, TwiAsante, Yoruba},
	}

	// Custom dataset models
	TTSModelVITSCustom = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCustom,
		model:              BaseModelVITS,
		defaultLanguage:    Catalan,
		currentLanguage:    Catalan,
		supportedLanguages: []Language{Catalan, Bengali},
	}

	TTSModelGlowTTSCustom = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCustom,
		model:              BaseModelGlowTTS,
		defaultLanguage:    Persian,
		currentLanguage:    Persian,
		supportedLanguages: []Language{Persian},
	}

	TTSModelVITSMaleCustom = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCustom,
		model:              BaseModelVITSMale,
		defaultLanguage:    Bengali,
		currentLanguage:    Bengali,
		supportedLanguages: []Language{Bengali},
	}

	TTSModelVITSFemaleCustom = TTSModel{
		category:           modelTypeTTS,
		dataset:            DatasetCustom,
		model:              BaseModelVITSFemale,
		defaultLanguage:    Bengali,
		currentLanguage:    Bengali,
		supportedLanguages: []Language{Bengali},
	}
)

// ttsModels contains a list of all predefined TTS models.
var TTSModels = ModelList[TTSModel]{
	models: []TTSModel{
		// Multilingual models
		TTSModelXTTSv2,
		TTSModelXTTSv1,
		TTSModelYourTTS,
		TTSModelBark,

		// Common Voice models
		TTSModelVITSCV,

		// English models
		TTSModelTacotron2EK1,
		TTSModelTacotron2DDCLJSpeech,
		TTSModelTacotron2DDCPhLJSpeech,
		TTSModelGlowTTSLJSpeech,
		TTSModelSpeedySpeechLJSpeech,
		TTSModelTacotron2DCALJSpeech,
		TTSModelVITSLJSpeech,
		TTSModelVITSNeonLJSpeech,
		TTSModelFastPitchLJSpeech,
		TTSModelOverflowLJSpeech,
		TTSModelNeuralHMMLJSpeech,
		TTSModelVITSVCTK,
		TTSModelFastPitchVCTK,
		TTSModelTacotron2DDCSam,
		TTSModelCapacitronT2C50Blizzard,
		TTSModelCapacitronT2C150v2Blizzard,
		TTSModelTortoiseV2,
		TTSModelJenny,

		// Multi-language models
		TTSModelTacotron2DDCMai,
		TTSModelGlowTTSMai,
		TTSModelVITSMai,
		TTSModelVITSCSS10,
		TTSModelVITSNeonCSS10,

		// Language-specific models
		TTSModelTacotron2DDCGSTBaker,
		TTSModelTacotron2DCAThorsten,
		TTSModelVITSThorsten,
		TTSModelTacotron2DDCThorsten,
		TTSModelTacotron2DDCKokoro,
		TTSModelGlowTTSCommonVoice,
		TTSModelGlowTTSMaiFemale,
		TTSModelVITSMaiFemale,
		TTSModelGlowTTSMaiMale,
		TTSModelVITSMaiMale,
		TTSModelVITSOpenBible,
		TTSModelVITSCustom,
		TTSModelGlowTTSCustom,
		TTSModelVITSMaleCustom,
		TTSModelVITSFemaleCustom,
	},
}

// NewTTSModel creates a new custom TTS model identifier.
func NewTTSModel(model BaseModel, dataset Dataset, language Language) (ModelIdentifier, error) {
	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return NewModel(modelTypeTTS, model, dataset, language)
}

// GetTTSModels returns a list of all predefined TTS models.
func GetTTSModels() []TTSModel {
	return slices.Clone(TTSModels.models)
}
