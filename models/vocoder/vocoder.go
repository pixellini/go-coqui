package vocoder

import (
	"slices"

	"github.com/pixellini/go-coqui/model"
)

type Vocoder = model.Identifier

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
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLibriTTS,
		Model:              BaseVocoderWavegrad,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(), // Universal vocoder.
	}

	VocoderWavegradEK1 = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetEK1,
		Model:              BaseVocoderWavegrad,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	VocoderWavegradThorsten = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              BaseVocoderWavegrad,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Fullband MelGAN model vocoders.
	VocoderFullbandMelganLibriTTS = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLibriTTS,
		Model:              BaseVocoderFullbandMelgan,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: model.GetSupportedLanguages(), // Universal vocoder.
	}

	VocoderFullbandMelganThorsten = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              BaseVocoderFullbandMelgan,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	// Multiband MelGAN model vocoders.
	VocoderMultibandMelganLJSpeech = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseVocoderMultibandMelgan,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	VocoderMultibandMelganMai = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetMai,
		Model:              BaseVocoderMultibandMelgan,
		DefaultLanguage:    model.Ukrainian,
		CurrentLanguage:    model.Ukrainian,
		SupportedLanguages: []model.Language{model.Ukrainian},
	}

	// HiFi-GAN v1 model vocoders.
	VocoderHifiganV1Thorsten = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetThorsten,
		Model:              BaseVocoderHifiganV1,
		DefaultLanguage:    model.German,
		CurrentLanguage:    model.German,
		SupportedLanguages: []model.Language{model.German},
	}

	VocoderHifiganV1Kokoro = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetKokoro,
		Model:              BaseVocoderHifiganV1,
		DefaultLanguage:    model.Japanese,
		CurrentLanguage:    model.Japanese,
		SupportedLanguages: []model.Language{model.Japanese},
	}

	// HiFi-GAN v2 model vocoders.
	VocoderHifiganV2LJSpeech = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseVocoderHifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2Blizzard2013 = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetBlizzard2013,
		Model:              BaseVocoderHifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2VCTK = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetVCTK,
		Model:              BaseVocoderHifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	VocoderHifiganV2Sam = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetSam,
		Model:              BaseVocoderHifiganV2,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// HiFi-GAN (generic) model vocoders.
	VocoderHifiganCommonVoiceTurkish = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetCommonVoice,
		Model:              BaseVocoderHifigan,
		DefaultLanguage:    model.Turkish,
		CurrentLanguage:    model.Turkish,
		SupportedLanguages: []model.Language{model.Turkish},
	}

	VocoderHifiganCommonVoiceBelarusian = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetCommonVoice,
		Model:              BaseVocoderHifigan,
		DefaultLanguage:    model.Belarusian,
		CurrentLanguage:    model.Belarusian,
		SupportedLanguages: []model.Language{model.Belarusian},
	}

	// UnivNet model vocoders.
	VocoderUnivnetLJSpeech = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetLJSpeech,
		Model:              BaseVocoderUnivnet,
		DefaultLanguage:    model.English,
		CurrentLanguage:    model.English,
		SupportedLanguages: []model.Language{model.English},
	}

	// Parallel WaveGAN model vocoders.
	VocoderParallelWaveganMai = Vocoder{
		Category:           model.TypeVocoder,
		Dataset:            model.DatasetMai,
		Model:              BaseVocoderParallelWavegan,
		DefaultLanguage:    model.Dutch,
		CurrentLanguage:    model.Dutch,
		SupportedLanguages: []model.Language{model.Dutch},
	}
)

// Vocoders contains a list of all predefined vocoder model identifiers.
var Vocoders = model.ModelList[Vocoder]{
	Models: []Vocoder{
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
func NewVocoder(base model.BaseModel, dataset model.Dataset, language model.Language) (Vocoder, error) {
	return model.NewModel(model.TypeVocoder, base, dataset, language)
}

// GetVocoders returns a list of all predefined vocoder models.
func GetVocoders() []Vocoder {
	return slices.Clone(Vocoders.Models)
}
