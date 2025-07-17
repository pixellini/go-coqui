package coqui

import (
	"fmt"
	"slices"
)

type VocoderModel struct {
	Type               ModelType
	Dataset            Dataset
	Architecture       Architecture
	SupportedLanguages []Language
	DefaultLanguage    Language
}

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by architecture type for easier discovery.
var (
	// Wavegrad architecture vocoders.
	VocoderWavegradLibriTTS = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetLibriTTS,
		Architecture:       ArchWavegrad,
		SupportedLanguages: allSupportedLanguages, // Universal vocoder.
		DefaultLanguage:    English,
	}

	VocoderWavegradEK1 = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetEK1,
		Architecture:       ArchWavegrad,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	VocoderWavegradThorsten = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetThorsten,
		Architecture:       ArchWavegrad,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	// Fullband MelGAN architecture vocoders.
	VocoderFullbandMelganLibriTTS = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetLibriTTS,
		Architecture:       ArchFullbandMelgan,
		SupportedLanguages: allSupportedLanguages, // Universal vocoder.
		DefaultLanguage:    English,
	}

	VocoderFullbandMelganThorsten = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetThorsten,
		Architecture:       ArchFullbandMelgan,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	// Multiband MelGAN architecture vocoders.
	VocoderMultibandMelganLJSpeech = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchMultibandMelgan,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	VocoderMultibandMelganMai = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetMai,
		Architecture:       ArchMultibandMelgan,
		SupportedLanguages: []Language{Ukrainian},
		DefaultLanguage:    Ukrainian,
	}

	// HiFi-GAN v1 architecture vocoders.
	VocoderHifiganV1Thorsten = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetThorsten,
		Architecture:       ArchHifiganV1,
		SupportedLanguages: []Language{German},
		DefaultLanguage:    German,
	}

	VocoderHifiganV1Kokoro = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetKokoro,
		Architecture:       ArchHifiganV1,
		SupportedLanguages: []Language{Japanese},
		DefaultLanguage:    Japanese,
	}

	// HiFi-GAN v2 architecture vocoders.
	VocoderHifiganV2LJSpeech = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchHifiganV2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	VocoderHifiganV2Blizzard2013 = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetBlizzard2013,
		Architecture:       ArchHifiganV2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	VocoderHifiganV2VCTK = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetVCTK,
		Architecture:       ArchHifiganV2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	VocoderHifiganV2Sam = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetSam,
		Architecture:       ArchHifiganV2,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// HiFi-GAN (generic) architecture vocoders.
	VocoderHifiganCommonVoiceTurkish = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetCommonVoice,
		Architecture:       ArchHifigan,
		SupportedLanguages: []Language{Turkish},
		DefaultLanguage:    Turkish,
	}

	VocoderHifiganCommonVoiceBelarusian = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetCommonVoice,
		Architecture:       ArchHifigan,
		SupportedLanguages: []Language{Belarusian},
		DefaultLanguage:    Belarusian,
	}

	// UnivNet architecture vocoders.
	VocoderUnivnetLJSpeech = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetLJSpeech,
		Architecture:       ArchUnivnet,
		SupportedLanguages: []Language{English},
		DefaultLanguage:    English,
	}

	// Parallel WaveGAN architecture vocoders.
	VocoderParallelWaveganMai = VocoderModel{
		Type:               ModelTypeVocoder,
		Dataset:            DatasetMai,
		Architecture:       ArchParallelWavegan,
		SupportedLanguages: []Language{Dutch},
		DefaultLanguage:    Dutch,
	}
)

// allVocoders contains a list of all predefined vocoder model identifiers.
var allVocoders = []VocoderModel{
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
}

// GetType returns the model type.
func (v VocoderModel) GetType() ModelType {
	return v.Type
}

// GetArchitecture returns the model architecture.
func (v VocoderModel) GetArchitecture() Architecture {
	return v.Architecture
}

// GetDataset returns the model dataset.
func (v VocoderModel) GetDataset() Dataset {
	return v.Dataset
}

// GetSupportedLanguages returns the supported languages.
func (v VocoderModel) GetSupportedLanguages() []Language {
	return v.SupportedLanguages
}

// GetDefaultLanguage returns the default language.
func (v VocoderModel) GetDefaultLanguage() Language {
	return v.DefaultLanguage
}

// IsValid checks if the model identifier is valid.
func (v VocoderModel) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), v.Type)
	// Check if the model has a valid type, dataset, architecture, and supported languages.
	return containsValidType && v.Type != "" && v.Dataset != "" && v.Architecture != "" && len(v.SupportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (v VocoderModel) IsMultilingual() bool {
	return len(v.SupportedLanguages) > 1
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (v VocoderModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", v.GetType(), v.GetDataset(), v.GetDefaultLanguage(), v.GetArchitecture())
}

// NewVocoderModel creates a new vocoder model identifier.
func NewVocoderModel(language, dataset string, architecture Architecture) VocoderModel {
	return VocoderModel{}
}

// GetVocoderModelsByArchitecture returns all vocoder models that use the specified architecture.
func GetVocoderModelsByArchitecture(arch Architecture) []VocoderModel {
	return FilterModelsByArchitecture(allVocoders, arch)
}

// GetVocoderModelsByDataset returns all vocoder models of the specified dataset.
func GetVocoderModelsByDataset(dataset Dataset) []VocoderModel {
	return FilterModelsByDataset(allVocoders, dataset)
}

// GetVocoderModelsByLanguage returns all vocoder models that support the specified language.
func GetVocoderModelsByLanguage(language Language) []VocoderModel {
	return FilterModelsBySupportedLanguages(allVocoders, []Language{language})
}

// GetVocoderModelsBySupportedLanguages returns all vocoder models that support any of the specified languages.
func GetVocoderModelsBySupportedLanguages(languages []Language) []VocoderModel {
	return FilterModelsBySupportedLanguages(allVocoders, languages)
}

// GetVocoderModelsByDefaultLanguage returns all vocoder models with the specified default language.
func GetVocoderModelsByDefaultLanguage(language Language) []VocoderModel {
	return FilterModelsByDefaultLanguage(allVocoders, language)
}

// GetVocoderModelsMultilingual returns all vocoder models that support multiple languages.
func GetVocoderModelsMultilingual() []VocoderModel {
	return FilterModelsMultilingual(allVocoders)
}
