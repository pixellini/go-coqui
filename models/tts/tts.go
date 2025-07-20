package tts

import (
	"github.com/pixellini/go-coqui/model"
)

type TTSModel = model.Identifier

const (
	// Multilingual architectures.
	BaseModelXTTSv2  model.BaseModel = "xtts_v2"
	BaseModelXTTSv1  model.BaseModel = "xtts_v1.1"
	BaseModelYourTTS model.BaseModel = "your_tts"
	BaseModelBark    model.BaseModel = "bark"

	// Single language architectures.
	BaseModelVITS         model.BaseModel = "vits"
	BaseModelVITSNeon     model.BaseModel = "vits--neon"
	BaseModelVITSNeonDash model.BaseModel = "vits-neon"
	BaseModelVITSMale     model.BaseModel = "vits-male"
	BaseModelVITSFemale   model.BaseModel = "vits-female"

	// Tacotron variants.
	BaseModelTacotron2       model.BaseModel = "tacotron2"
	BaseModelTacotron2DDC    model.BaseModel = "tacotron2-DDC"
	BaseModelTacotron2DDCPh  model.BaseModel = "tacotron2-DDC_ph"
	BaseModelTacotron2DCA    model.BaseModel = "tacotron2-DCA"
	BaseModelTacotron2DDCGST model.BaseModel = "tacotron2-DDC-GST"

	// Other architectures.
	BaseModelGlowTTS            model.BaseModel = "glow-tts"
	BaseModelFastPitch          model.BaseModel = "fast_pitch"
	BaseModelSpeedySpeech       model.BaseModel = "speedy-speech"
	BaseModelOverflow           model.BaseModel = "overflow"
	BaseModelNeuralHMM          model.BaseModel = "neural_hmm"
	BaseModelTortoise           model.BaseModel = "tortoise-v2"
	BaseModelCapacitronT2C50    model.BaseModel = "capacitron-t2-c50"
	BaseModelCapacitronT2C150v2 model.BaseModel = "capacitron-t2-c150_v2"
	BaseModelJenny              model.BaseModel = "jenny"
)

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	TTSModelXTTSv2 = TTSModel{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                BaseModelXTTSv2,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	TTSModelXTTSv1 = TTSModel{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                BaseModelXTTSv1,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	TTSModelYourTTS = TTSModel{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                BaseModelYourTTS,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	TTSModelBark = TTSModel{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                BaseModelBark,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	// Common Voice (CV) dataset models
	TTSModelVITSCV = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCV,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.English, // Fallback to model.English if available
		CurrentLanguage:    model.English, // Fallback to model.English if available
		SupportedLanguages: []model.Language{model.Bulgarian, model.Czech, model.Danish, model.Estonian, model.Irish, model.Greek, model.Croatian, model.Lithuanian, model.Latvian, model.Maltese, model.Portuguese, model.Romanian, model.Slovak, model.Slovenian, model.Swedish},
	}

	// model.English dataset models
	TTSModelTacotron2EK1 = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetEK1,
		Model:              BaseModelTacotron2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// LJSpeech dataset models
	TTSModelTacotron2DDCLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelTacotron2DDC,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelTacotron2DDCPhLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelTacotron2DDCPh,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelGlowTTSLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelSpeedySpeechLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelSpeedySpeech,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelTacotron2DCALJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelTacotron2DCA,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelVITSLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelVITSNeonLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelVITSNeon,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelFastPitchLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelFastPitch,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelOverflowLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelOverflow,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelNeuralHMMLJSpeech = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseModelNeuralHMM,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// VCTK dataset models
	TTSModelVITSVCTK = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetVCTK,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelFastPitchVCTK = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetVCTK,
		Model:              BaseModelFastPitch,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Sam dataset models
	TTSModelTacotron2DDCSam = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetSam,
		Model:              BaseModelTacotron2DDC,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Blizzard2013 dataset models
	TTSModelCapacitronT2C50Blizzard = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBlizzard2013,
		Model:              BaseModelCapacitronT2C50,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	TTSModelCapacitronT2C150v2Blizzard = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBlizzard2013,
		Model:              BaseModelCapacitronT2C150v2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Multi-dataset model.English models
	TTSModelTortoiseV2 = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMultiDataset,
		Model:              BaseModelTortoise,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Jenny dataset models
	TTSModelJenny = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetJenny,
		Model:              BaseModelJenny,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Mai dataset models (multiple languages)
	TTSModelTacotron2DDCMai = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              BaseModelTacotron2DDC,
		DefaultLanguage:    model.Spanish,
		CurrentLanguage:    model.Spanish,
		SupportedLanguages: []model.Language{model.Spanish, model.French, model.Dutch},
	}

	TTSModelGlowTTSMai = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	TTSModelVITSMai = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	// CSS10 dataset models (multiple languages)
	TTSModelVITSCSS10 = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCSS10,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Spanish,
		CurrentLanguage:    model.Spanish,
		SupportedLanguages: []model.Language{model.Spanish, model.French, model.German, model.Dutch, model.Hungarian, model.Finnish},
	}

	TTSModelVITSNeonCSS10 = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCSS10,
		Model:              BaseModelVITSNeon,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Baker dataset models (Chinese)
	TTSModelTacotron2DDCGSTBaker = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBaker,
		Model:              BaseModelTacotron2DDCGST,
		DefaultLanguage:    model.Chinese,
		CurrentLanguage:    model.Chinese,
		SupportedLanguages: []model.Language{model.Chinese},
	}

	// Thorsten dataset models (German)
	TTSModelTacotron2DCAThorsten = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              BaseModelTacotron2DCA,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	TTSModelVITSThorsten = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	TTSModelTacotron2DDCThorsten = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              BaseModelTacotron2DDC,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Kokoro dataset models (Japanese)
	TTSModelTacotron2DDCKokoro = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetKokoro,
		Model:              BaseModelTacotron2DDC,
		DefaultLanguage:    model.Japanese,
		CurrentLanguage:    model.Japanese,
		SupportedLanguages: []model.Language{model.Japanese},
	}

	// Common Voice dataset models
	TTSModelGlowTTSCommonVoice = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCommonVoice,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.Turkish,
		CurrentLanguage:    model.Turkish,
		SupportedLanguages: []model.Language{model.Turkish, model.Belarusian},
	}

	// Mai Female dataset models (Italian)
	TTSModelGlowTTSMaiFemale = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiFemale,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	TTSModelVITSMaiFemale = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiFemale,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian, model.Polish},
	}

	// Mai Male dataset models (Italian)
	TTSModelGlowTTSMaiMale = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiMale,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	TTSModelVITSMaiMale = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiMale,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	// OpenBible dataset models (African languages)
	TTSModelVITSOpenBible = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetOpenBible,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Hausa,
		CurrentLanguage:    model.Hausa,
		SupportedLanguages: []model.Language{model.Ewe, model.Hausa, model.Lin, model.TwiAkuapem, model.TwiAsante, model.Yoruba},
	}

	// Custom dataset models
	TTSModelVITSCustom = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              BaseModelVITS,
		DefaultLanguage:    model.Catalan,
		CurrentLanguage:    model.Catalan,
		SupportedLanguages: []model.Language{model.Catalan, model.Bengali},
	}

	TTSModelGlowTTSCustom = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              BaseModelGlowTTS,
		DefaultLanguage:    model.Persian,
		CurrentLanguage:    model.Persian,
		SupportedLanguages: []model.Language{model.Persian},
	}

	TTSModelVITSMaleCustom = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              BaseModelVITSMale,
		DefaultLanguage:    model.Bengali,
		CurrentLanguage:    model.Bengali,
		SupportedLanguages: []model.Language{model.Bengali},
	}

	TTSModelVITSFemaleCustom = TTSModel{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              BaseModelVITSFemale,
		DefaultLanguage:    model.Bengali,
		CurrentLanguage:    model.Bengali,
		SupportedLanguages: []model.Language{model.Bengali},
	}
)

// ttsModels contains a list of all predefined TTS models.
var TTSModels = model.ModelList[TTSModel]{
	Models: []TTSModel{
		// Multilingual models
		TTSModelXTTSv2,
		TTSModelXTTSv1,
		TTSModelYourTTS,
		TTSModelBark,

		// Common Voice models
		TTSModelVITSCV,

		// model.English models
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
func NewTTSModel(base model.BaseModel, dataset model.Dataset, language model.Language) (model.Identifier, error) {
	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return model.NewModel(model.TypeTTS, base, dataset, language)
}

// GetTTSModels returns a list of all predefined TTS models.
func GetTTSModels() []TTSModel {
	// return slices.Clone(TTSModels.Models)
	return []TTSModel{}
}
