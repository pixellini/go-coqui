package coqui

import "slices"

type Vocoder = ModelIdentifier

const ( // Vocoder architectures.
	BaseVocoderWavegrad        BaseModel = "wavegrad"
	BaseVocoderFullbandMelgan  BaseModel = "fullband-melgan"
	BaseVocoderMultibandMelgan BaseModel = "multiband-melgan"
	BaseVocoderHifiganV1       BaseModel = "hifigan_v1"
	BaseVocoderHifiganV2       BaseModel = "hifigan_v2"
	BaseVocoderHifigan         BaseModel = "hifigan"
	BaseVocoderUnivnet         BaseModel = "univnet"
	BaseVocoderParallelWavegan BaseModel = "parallel-wavegan"
)

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by model type for easier discovery.
var (
	// Wavegrad model vocoders.
	VocoderWavegradLibriTTS = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetLibriTTS,
		model:              BaseVocoderWavegrad,
		defaultLanguage:    English,
		supportedLanguages: supportedLanguages, // Universal vocoder.
	}

	VocoderWavegradEK1 = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetEK1,
		model:              BaseVocoderWavegrad,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderWavegradThorsten = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetThorsten,
		model:              BaseVocoderWavegrad,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Fullband MelGAN model vocoders.
	VocoderFullbandMelganLibriTTS = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetLibriTTS,
		model:              BaseVocoderFullbandMelgan,
		defaultLanguage:    English,
		supportedLanguages: supportedLanguages, // Universal vocoder.
	}

	VocoderFullbandMelganThorsten = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetThorsten,
		model:              BaseVocoderFullbandMelgan,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Multiband MelGAN model vocoders.
	VocoderMultibandMelganLJSpeech = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		model:              BaseVocoderMultibandMelgan,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderMultibandMelganMai = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetMai,
		model:              BaseVocoderMultibandMelgan,
		defaultLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	// HiFi-GAN v1 model vocoders.
	VocoderHifiganV1Thorsten = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetThorsten,
		model:              BaseVocoderHifiganV1,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	VocoderHifiganV1Kokoro = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetKokoro,
		model:              BaseVocoderHifiganV1,
		defaultLanguage:    Japanese,
		supportedLanguages: []Language{Japanese},
	}

	// HiFi-GAN v2 model vocoders.
	VocoderHifiganV2LJSpeech = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		model:              BaseVocoderHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2Blizzard2013 = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetBlizzard2013,
		model:              BaseVocoderHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2VCTK = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetVCTK,
		model:              BaseVocoderHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2Sam = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetSam,
		model:              BaseVocoderHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// HiFi-GAN (generic) model vocoders.
	VocoderHifiganCommonVoiceTurkish = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetCommonVoice,
		model:              BaseVocoderHifigan,
		defaultLanguage:    Turkish,
		supportedLanguages: []Language{Turkish},
	}

	VocoderHifiganCommonVoiceBelarusian = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetCommonVoice,
		model:              BaseVocoderHifigan,
		defaultLanguage:    Belarusian,
		supportedLanguages: []Language{Belarusian},
	}

	// UnivNet model vocoders.
	VocoderUnivnetLJSpeech = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		model:              BaseVocoderUnivnet,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Parallel WaveGAN model vocoders.
	VocoderParallelWaveganMai = Vocoder{
		category:           modelTypeVocoder,
		dataset:            DatasetMai,
		model:              BaseVocoderParallelWavegan,
		defaultLanguage:    Dutch,
		supportedLanguages: []Language{Dutch},
	}
)

// Vocoders contains a list of all predefined vocoder model identifiers.
var Vocoders = ModelList[Vocoder]{
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
func NewVocoder(language Language, dataset Dataset, model BaseModel) (Vocoder, error) {
	return NewModel(modelTypeVocoder, language, dataset, model)
}

// GetVocoders returns a list of all predefined vocoder models.
func GetVocoders() []Vocoder {
	return slices.Clone(Vocoders.models)
}
