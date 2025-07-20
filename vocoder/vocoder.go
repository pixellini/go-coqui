package vocoder

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type Vocoder = model.ModelIdentifier

const ( // Vocoder architectures.
	BaseVocoderWavegrad        model.BaseModel = "wavegrad"
	BaseVocoderFullbandMelgan  model.BaseModel = "fullband-melgan"
	BaseVocoderMultibandMelgan model.BaseModel = "multiband-melgan"
	BaseVocoderHifiganV1       model.BaseModel = "hifigan_v1"
	BaseVocoderHifiganV2       model.BaseModel = "hifigan_v2"
	BaseVocoderHifigan         model.BaseModel = "hifigan"
	BaseVocoderUnivnet         model.BaseModel = "univnet"
	BaseVocoderParallelWavegan model.BaseModel = "parallel-wavegan"
)

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by model type for easier discovery.
var (
	// Wavegrad model vocoders.
	VocoderWavegradLibriTTS = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetLibriTTS,
		model:              model.BaseVocoderWavegrad,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: model.supportedLanguages, // Universal vocoder.
	}

	VocoderWavegradEK1 = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetEK1,
		model:              model.BaseVocoderWavegrad,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	VocoderWavegradThorsten = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetThorsten,
		model:              model.BaseVocoderWavegrad,
		defaultLanguage:    model.German,
		currentLanguage:    model.German,
		supportedLanguages: []model.Language{model.German},
	}

	// Fullband MelGAN model vocoders.
	VocoderFullbandMelganLibriTTS = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetLibriTTS,
		model:              model.BaseVocoderFullbandMelgan,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: model.supportedLanguages, // Universal vocoder.
	}

	VocoderFullbandMelganThorsten = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetThorsten,
		model:              model.BaseVocoderFullbandMelgan,
		defaultLanguage:    model.German,
		currentLanguage:    model.German,
		supportedLanguages: []model.Language{model.German},
	}

	// Multiband MelGAN model vocoders.
	VocoderMultibandMelganLJSpeech = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetLJSpeech,
		model:              model.BaseVocoderMultibandMelgan,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	VocoderMultibandMelganMai = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetMai,
		model:              model.BaseVocoderMultibandMelgan,
		defaultLanguage:    model.Ukrainian,
		currentLanguage:    model.Ukrainian,
		supportedLanguages: []model.Language{model.Ukrainian},
	}

	// HiFi-GAN v1 model vocoders.
	VocoderHifiganV1Thorsten = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetThorsten,
		model:              model.BaseVocoderHifiganV1,
		defaultLanguage:    model.German,
		currentLanguage:    model.German,
		supportedLanguages: []model.Language{model.German},
	}

	VocoderHifiganV1Kokoro = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetKokoro,
		model:              model.BaseVocoderHifiganV1,
		defaultLanguage:    model.Japanese,
		currentLanguage:    model.Japanese,
		supportedLanguages: []model.Language{model.Japanese},
	}

	// HiFi-GAN v2 model vocoders.
	VocoderHifiganV2LJSpeech = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetLJSpeech,
		model:              model.BaseVocoderHifiganV2,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2Blizzard2013 = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetBlizzard2013,
		model:              model.BaseVocoderHifiganV2,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2VCTK = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetVCTK,
		model:              model.BaseVocoderHifiganV2,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2Sam = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetSam,
		model:              model.BaseVocoderHifiganV2,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	// HiFi-GAN (generic) model vocoders.
	VocoderHifiganCommonVoiceTurkish = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetCommonVoice,
		model:              model.BaseVocoderHifigan,
		defaultLanguage:    model.Turkish,
		currentLanguage:    model.Turkish,
		supportedLanguages: []model.Language{model.Turkish},
	}

	VocoderHifiganCommonVoiceBelarusian = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetCommonVoice,
		model:              model.BaseVocoderHifigan,
		defaultLanguage:    model.Belarusian,
		currentLanguage:    model.Belarusian,
		supportedLanguages: []model.Language{model.Belarusian},
	}

	// UnivNet model vocoders.
	VocoderUnivnetLJSpeech = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetLJSpeech,
		model:              model.BaseVocoderUnivnet,
		defaultLanguage:    model.English,
		currentLanguage:    model.English,
		supportedLanguages: []model.Language{model.English},
	}

	// Parallel WaveGAN model vocoders.
	VocoderParallelWaveganMai = Vocoder{
		category:           model.modelTypeVocoder,
		dataset:            model.DatasetMai,
		model:              model.BaseVocoderParallelWavegan,
		defaultLanguage:    model.Dutch,
		currentLanguage:    model.Dutch,
		supportedLanguages: []model.Language{model.Dutch},
	}
)

// Vocoders contains a list of all predefined vocoder model identifiers.
var Vocoders = model.ModelList[Vocoder]{
	models: []Vocoder{
		// Wavegrad vocoders.
		VocoderWavegradLibriTTS,
		VocoderWavegradEK1,
		VocoderWavegradThorsten,

		// Fullband MelGAN vocoders.
		VocoderFullbandMelganLibriTTS,
		VocoderFullbandMelganThorsten,

		// Multiband MelGAN vocoders.
		VocoderMultibandMelganLJSpeech,
		VocoderMultibandMelganMai,

		// HiFi-GAN v1 vocoders.
		VocoderHifiganV1Thorsten,
		VocoderHifiganV1Kokoro,

		// HiFi-GAN v2 vocoders.
		VocoderHifiganV2LJSpeech,
		VocoderHifiganV2Blizzard2013,
		VocoderHifiganV2VCTK,
		VocoderHifiganV2Sam,

		// HiFi-GAN (generic) vocoders.
		VocoderHifiganCommonVoiceTurkish,
		VocoderHifiganCommonVoiceBelarusian,

		// UnivNet vocoders.
		VocoderUnivnetLJSpeech,

		// Parallel WaveGAN vocoders.
		VocoderParallelWaveganMai,
	},
}

// NewVocoder creates a new vocoder model identifier.
func NewVocoder(model model.BaseModel, dataset model.Dataset, language model.Language) (Vocoder, error) {
	return model.NewModel(model.modelTypeVocoder, model, dataset, language)
}

// GetVocoders returns a list of all predefined vocoder models.
func GetVocoders() []Vocoder {
	return slices.Clone(Vocoders.models)
}
