package tts

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type Model = model.Identifier

const (
	// Multilingual architectures.
	XTTSv2  model.BaseModel = "xtts_v2"
	XTTSv1  model.BaseModel = "xtts_v1.1"
	YourTTS model.BaseModel = "your_tts"
	Bark    model.BaseModel = "bark"

	// Single language architectures.
	VITS         model.BaseModel = "vits"
	VITSNeon     model.BaseModel = "vits--neon"
	VITSNeonDash model.BaseModel = "vits-neon"
	VITSMale     model.BaseModel = "vits-male"
	VITSFemale   model.BaseModel = "vits-female"

	// Tacotron variants.
	Tacotron2       model.BaseModel = "tacotron2"
	Tacotron2DDC    model.BaseModel = "tacotron2-DDC"
	Tacotron2DDCPh  model.BaseModel = "tacotron2-DDC_ph"
	Tacotron2DCA    model.BaseModel = "tacotron2-DCA"
	Tacotron2DDCGST model.BaseModel = "tacotron2-DDC-GST"

	// Other architectures.
	GlowTTS            model.BaseModel = "glow-tts"
	FastPitch          model.BaseModel = "fast_pitch"
	SpeedySpeech       model.BaseModel = "speedy-speech"
	Overflow           model.BaseModel = "overflow"
	NeuralHMM          model.BaseModel = "neural_hmm"
	Tortoise           model.BaseModel = "tortoise-v2"
	CapacitronT2C50    model.BaseModel = "capacitron-t2-c50"
	CapacitronT2C150v2 model.BaseModel = "capacitron-t2-c150_v2"
	Jenny              model.BaseModel = "jenny"
)

// Presets of predefined TTS models for easy access that is currently available in the Coqui TTS library.
var (
	// Multilingual models (support all languages)
	PresetXTTSv2 = Model{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                XTTSv2,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	PresetXTTSv1 = Model{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                XTTSv1,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	PresetYourTTS = Model{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                YourTTS,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	PresetBark = Model{
		Category:             model.TypeTTS,
		Dataset:              model.DatasetMultiDataset,
		Model:                Bark,
		DefaultLanguage:      model.English,
		CurrentLanguage:      model.English,
		SupportedLanguages:   model.GetSupportedLanguages(),
		SupportsVoiceCloning: true,
	}

	// Common Voice (CV) dataset models
	PresetVITSCV = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCV,
		Model:              VITS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.Bulgarian, model.Czech, model.Danish, model.Estonian, model.Irish, model.Greek, model.Croatian, model.Lithuanian, model.Latvian, model.Maltese, model.Portuguese, model.Romanian, model.Slovak, model.Slovenian, model.Swedish},
	}

	// model.English dataset models
	PresetTacotron2EK1 = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetEK1,
		Model:              Tacotron2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// LJSpeech dataset models
	PresetTacotron2DDCLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              Tacotron2DDC,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetTacotron2DDCPhLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              Tacotron2DDCPh,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetGlowTTSLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              GlowTTS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetSpeedySpeechLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              SpeedySpeech,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetTacotron2DCALJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              Tacotron2DCA,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetVITSLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              VITS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetVITSNeonLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              VITSNeon,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetFastPitchLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              FastPitch,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetOverflowLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              Overflow,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetNeuralHMMLJSpeech = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetLJSpeech,
		Model:              NeuralHMM,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// VCTK dataset models
	PresetVITSVCTK = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetVCTK,
		Model:              VITS,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetFastPitchVCTK = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetVCTK,
		Model:              FastPitch,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Sam dataset models
	PresetTacotron2DDCSam = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetSam,
		Model:              Tacotron2DDC,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Blizzard2013 dataset models
	PresetCapacitronT2C50Blizzard = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBlizzard2013,
		Model:              CapacitronT2C50,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetCapacitronT2C150v2Blizzard = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBlizzard2013,
		Model:              CapacitronT2C150v2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Multi-dataset model.English models
	PresetTortoiseV2 = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMultiDataset,
		Model:              Tortoise,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Jenny dataset models
	PresetJenny = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetJenny,
		Model:              Jenny,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Mai dataset models (multiple languages)
	PresetTacotron2DDCMai = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              Tacotron2DDC,
		DefaultLanguage:    model.Spanish,
		CurrentLanguage:    model.Spanish,
		SupportedLanguages: []model.Language{model.Spanish, model.French, model.Dutch},
	}

	PresetGlowTTSMai = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              GlowTTS,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	PresetVITSMai = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMai,
		Model:              VITS,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	// CSS10 dataset models (multiple languages)
	PresetVITSCSS10 = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCSS10,
		Model:              VITS,
		DefaultLanguage:    model.Spanish,
		CurrentLanguage:    model.Spanish,
		SupportedLanguages: []model.Language{model.Spanish, model.French, model.German, model.Dutch, model.Hungarian, model.Finnish},
	}

	PresetVITSNeonCSS10 = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCSS10,
		Model:              VITSNeon,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Baker dataset models (Chinese)
	PresetTacotron2DDCGSTBaker = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetBaker,
		Model:              Tacotron2DDCGST,
		DefaultLanguage:    model.Chinese,
		CurrentLanguage:    model.Chinese,
		SupportedLanguages: []model.Language{model.Chinese},
	}

	// Thorsten dataset models (German)
	PresetTacotron2DCAThorsten = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              Tacotron2DCA,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	PresetVITSThorsten = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              VITS,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	PresetTacotron2DDCThorsten = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetThorsten,
		Model:              Tacotron2DDC,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Kokoro dataset models (Japanese)
	PresetTacotron2DDCKokoro = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetKokoro,
		Model:              Tacotron2DDC,
		DefaultLanguage:    model.Japanese,
		CurrentLanguage:    model.Japanese,
		SupportedLanguages: []model.Language{model.Japanese},
	}

	// Common Voice dataset models
	PresetGlowTTSCommonVoice = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCommonVoice,
		Model:              GlowTTS,
		DefaultLanguage:    model.Turkish,
		CurrentLanguage:    model.Turkish,
		SupportedLanguages: []model.Language{model.Turkish, model.Belarusian},
	}

	// Mai Female dataset models (Italian)
	PresetGlowTTSMaiFemale = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiFemale,
		Model:              GlowTTS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	PresetVITSMaiFemale = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiFemale,
		Model:              VITS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian, model.Polish},
	}

	// Mai Male dataset models (Italian)
	PresetGlowTTSMaiMale = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiMale,
		Model:              GlowTTS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	PresetVITSMaiMale = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetMaiMale,
		Model:              VITS,
		DefaultLanguage:    model.Italian,
		CurrentLanguage:    model.Italian,
		SupportedLanguages: []model.Language{model.Italian},
	}

	// OpenBible dataset models (African languages)
	PresetVITSOpenBible = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetOpenBible,
		Model:              VITS,
		DefaultLanguage:    model.Hausa,
		CurrentLanguage:    model.Hausa,
		SupportedLanguages: []model.Language{model.Ewe, model.Hausa, model.Lin, model.TwiAkuapem, model.TwiAsante, model.Yoruba},
	}

	// Custom dataset models
	PresetVITSCustom = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              VITS,
		DefaultLanguage:    model.Catalan,
		CurrentLanguage:    model.Catalan,
		SupportedLanguages: []model.Language{model.Catalan, model.Bengali},
	}

	PresetGlowTTSCustom = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              GlowTTS,
		DefaultLanguage:    model.Persian,
		CurrentLanguage:    model.Persian,
		SupportedLanguages: []model.Language{model.Persian},
	}

	PresetVITSMaleCustom = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              VITSMale,
		DefaultLanguage:    model.Bengali,
		CurrentLanguage:    model.Bengali,
		SupportedLanguages: []model.Language{model.Bengali},
	}

	PresetVITSFemaleCustom = Model{
		Category:           model.TypeTTS,
		Dataset:            model.DatasetCustom,
		Model:              VITSFemale,
		DefaultLanguage:    model.Bengali,
		CurrentLanguage:    model.Bengali,
		SupportedLanguages: []model.Language{model.Bengali},
	}
)

// presets contains a list of all predefined TTS models.
var presets = model.ModelList[Model]{
	Models: []Model{
		// Multilingual models
		PresetXTTSv2,
		PresetXTTSv1,
		PresetYourTTS,
		PresetBark,

		// Common Voice models
		PresetVITSCV,

		// model.English models
		PresetTacotron2EK1,
		PresetTacotron2DDCLJSpeech,
		PresetTacotron2DDCPhLJSpeech,
		PresetGlowTTSLJSpeech,
		PresetSpeedySpeechLJSpeech,
		PresetTacotron2DCALJSpeech,
		PresetVITSLJSpeech,
		PresetVITSNeonLJSpeech,
		PresetFastPitchLJSpeech,
		PresetOverflowLJSpeech,
		PresetNeuralHMMLJSpeech,
		PresetVITSVCTK,
		PresetFastPitchVCTK,
		PresetTacotron2DDCSam,
		PresetCapacitronT2C50Blizzard,
		PresetCapacitronT2C150v2Blizzard,
		PresetTortoiseV2,
		PresetJenny,

		// Multi-language models
		PresetTacotron2DDCMai,
		PresetGlowTTSMai,
		PresetVITSMai,
		PresetVITSCSS10,
		PresetVITSNeonCSS10,

		// Language-specific models
		PresetTacotron2DDCGSTBaker,
		PresetTacotron2DCAThorsten,
		PresetVITSThorsten,
		PresetTacotron2DDCThorsten,
		PresetTacotron2DDCKokoro,
		PresetGlowTTSCommonVoice,
		PresetGlowTTSMaiFemale,
		PresetVITSMaiFemale,
		PresetGlowTTSMaiMale,
		PresetVITSMaiMale,
		PresetVITSOpenBible,
		PresetVITSCustom,
		PresetGlowTTSCustom,
		PresetVITSMaleCustom,
		PresetVITSFemaleCustom,
	},
}

// New creates a new custom TTS model identifier.
func New(base model.BaseModel, dataset model.Dataset, language model.Language) (model.Identifier, error) {
	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return model.NewModel(model.TypeTTS, base, dataset, language)
}

// GetPresets returns a list of all predefined TTS models.
func GetPresets() []Model {
	return slices.Clone(presets.Models)
}
