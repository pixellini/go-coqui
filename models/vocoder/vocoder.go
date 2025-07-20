package vocoder

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type Model = model.Identifier

const ( // Vocoder architectures.
	Wavegrad        model.BaseModel = "wavegrad"
	FullbandMelgan  model.BaseModel = "fullband-melgan"
	MultibandMelgan model.BaseModel = "multiband-melgan"
	HifiganV1       model.BaseModel = "hifigan_v1"
	HifiganV2       model.BaseModel = "hifigan_v2"
	Hifigan         model.BaseModel = "hifigan"
	Univnet         model.BaseModel = "univnet"
	ParallelWavegan model.BaseModel = "parallel-wavegan"
)

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by model type for easier discovery.
var (
	// Wavegrad model vocoders.
	PresetWavegradLibriTTS = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLibriTTS,
		Model:              Wavegrad,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(), // Universal vocoder.
	}

	PresetWavegradEK1 = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetEK1,
		Model:              Wavegrad,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetWavegradThorsten = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              Wavegrad,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Fullband MelGAN model vocoders.
	PresetFullbandMelganLibriTTS = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLibriTTS,
		Model:              FullbandMelgan,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(), // Universal vocoder.
	}

	PresetFullbandMelganThorsten = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              FullbandMelgan,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Multiband MelGAN model vocoders.
	PresetMultibandMelganLJSpeech = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              MultibandMelgan,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetMultibandMelganMai = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetMai,
		Model:              MultibandMelgan,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	// HiFi-GAN v1 model vocoders.
	PresetHifiganV1Thorsten = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              HifiganV1,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	PresetHifiganV1Kokoro = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetKokoro,
		Model:              HifiganV1,
		DefaultLanguage:    model.Japanese,
		CurrentLanguage:    model.Japanese,
		SupportedLanguages: []model.Language{model.Japanese},
	}

	// HiFi-GAN v2 model vocoders.
	PresetHifiganV2LJSpeech = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              HifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetHifiganV2Blizzard2013 = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetBlizzard2013,
		Model:              HifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetHifiganV2VCTK = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetVCTK,
		Model:              HifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	PresetHifiganV2Sam = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetSam,
		Model:              HifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// HiFi-GAN (generic) model vocoders.
	PresetHifiganCommonVoiceTurkish = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetCommonVoice,
		Model:              Hifigan,
		DefaultLanguage:    model.Turkish,
		CurrentLanguage:    model.Turkish,
		SupportedLanguages: []model.Language{model.Turkish},
	}

	PresetHifiganCommonVoiceBelarusian = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetCommonVoice,
		Model:              Hifigan,
		DefaultLanguage:    model.Belarusian,
		CurrentLanguage:    model.Belarusian,
		SupportedLanguages: []model.Language{model.Belarusian},
	}

	// UnivNet model vocoders.
	PresetUnivnetLJSpeech = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              Univnet,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Parallel WaveGAN model vocoders.
	PresetParallelWaveganMai = Model{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetMai,
		Model:              ParallelWavegan,
		DefaultLanguage:    model.Dutch,
		CurrentLanguage:    model.Dutch,
		SupportedLanguages: []model.Language{model.Dutch},
	}
)

// presets contains a list of all predefined vocoder model identifiers.
var presets = model.ModelList[Model]{
	Models: []Model{
		// Wavegrad vocoders.
		PresetWavegradLibriTTS,
		PresetWavegradEK1,
		PresetWavegradThorsten,

		// Fullband MelGAN vocoders.
		PresetFullbandMelganLibriTTS,
		PresetFullbandMelganThorsten,

		// Multiband MelGAN vocoders.
		PresetMultibandMelganLJSpeech,
		PresetMultibandMelganMai,

		// HiFi-GAN v1 vocoders.
		PresetHifiganV1Thorsten,
		PresetHifiganV1Kokoro,

		// HiFi-GAN v2 vocoders.
		PresetHifiganV2LJSpeech,
		PresetHifiganV2Blizzard2013,
		PresetHifiganV2VCTK,
		PresetHifiganV2Sam,

		// HiFi-GAN (generic) vocoders.
		PresetHifiganCommonVoiceTurkish,
		PresetHifiganCommonVoiceBelarusian,

		// UnivNet vocoders.
		PresetUnivnetLJSpeech,

		// Parallel WaveGAN vocoders.
		PresetParallelWaveganMai,
	},
}

// New creates a new vocoder model identifier.
func New(base model.BaseModel, dataset model.Dataset, language model.Language) (Model, error) {
	return model.NewModel(model.TypeVocoder, base, dataset, language)
}

// GetPresets returns a list of all predefined vocoder models.
func GetPresets() []Model {
	return slices.Clone(presets.Models)
}
