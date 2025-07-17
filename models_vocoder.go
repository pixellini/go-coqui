package coqui

import (
	"fmt"
	"slices"
)

type VocoderModel struct {
	modelType          ModelType
	dataset            Dataset
	architecture       Architecture
	supportedLanguages []Language
	defaultLanguage    Language
}

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by architecture type for easier discovery.
var (
	// Wavegrad architecture vocoders.
	VocoderWavegradLibriTTS = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetLibriTTS,
		architecture:       ArchWavegrad,
		supportedLanguages: allSupportedLanguages, // Universal vocoder.
		defaultLanguage:    English,
	}

	VocoderWavegradEK1 = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetEK1,
		architecture:       ArchWavegrad,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	VocoderWavegradThorsten = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchWavegrad,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Fullband MelGAN architecture vocoders.
	VocoderFullbandMelganLibriTTS = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetLibriTTS,
		architecture:       ArchFullbandMelgan,
		supportedLanguages: allSupportedLanguages, // Universal vocoder.
		defaultLanguage:    English,
	}

	VocoderFullbandMelganThorsten = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchFullbandMelgan,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	// Multiband MelGAN architecture vocoders.
	VocoderMultibandMelganLJSpeech = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchMultibandMelgan,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	VocoderMultibandMelganMai = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetMai,
		architecture:       ArchMultibandMelgan,
		supportedLanguages: []Language{Ukrainian},
		defaultLanguage:    Ukrainian,
	}

	// HiFi-GAN v1 architecture vocoders.
	VocoderHifiganV1Thorsten = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchHifiganV1,
		supportedLanguages: []Language{German},
		defaultLanguage:    German,
	}

	VocoderHifiganV1Kokoro = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetKokoro,
		architecture:       ArchHifiganV1,
		supportedLanguages: []Language{Japanese},
		defaultLanguage:    Japanese,
	}

	// HiFi-GAN v2 architecture vocoders.
	VocoderHifiganV2LJSpeech = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchHifiganV2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	VocoderHifiganV2Blizzard2013 = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchHifiganV2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	VocoderHifiganV2VCTK = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetVCTK,
		architecture:       ArchHifiganV2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	VocoderHifiganV2Sam = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetSam,
		architecture:       ArchHifiganV2,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// HiFi-GAN (generic) architecture vocoders.
	VocoderHifiganCommonVoiceTurkish = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetCommonVoice,
		architecture:       ArchHifigan,
		supportedLanguages: []Language{Turkish},
		defaultLanguage:    Turkish,
	}

	VocoderHifiganCommonVoiceBelarusian = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetCommonVoice,
		architecture:       ArchHifigan,
		supportedLanguages: []Language{Belarusian},
		defaultLanguage:    Belarusian,
	}

	// UnivNet architecture vocoders.
	VocoderUnivnetLJSpeech = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchUnivnet,
		supportedLanguages: []Language{English},
		defaultLanguage:    English,
	}

	// Parallel WaveGAN architecture vocoders.
	VocoderParallelWaveganMai = VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            DatasetMai,
		architecture:       ArchParallelWavegan,
		supportedLanguages: []Language{Dutch},
		defaultLanguage:    Dutch,
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

// NewVocoderModel creates a new vocoder model identifier.
func NewVocoderModel(language Language, dataset Dataset, architecture Architecture) *VocoderModel {
	return &VocoderModel{
		modelType:          ModelTypeVocoder,
		dataset:            dataset,
		architecture:       architecture,
		supportedLanguages: []Language{Language(language)},
		defaultLanguage:    Language(language),
	}
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (v VocoderModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", v.GetModelType(), v.GetDataset(), v.GetDefaultLanguage(), v.GetArchitecture())
}

// IsValid checks if the model identifier is valid.
func (v VocoderModel) IsValid() bool {
	containsValidType := slices.Contains(GetAllModelTypes(), v.modelType)
	// Check if the model has a valid type, dataset, architecture, and supported languages.
	return containsValidType && v.modelType != "" && v.dataset != "" && v.architecture != "" && len(v.supportedLanguages) > 0
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (v VocoderModel) IsMultilingual() bool {
	return len(v.supportedLanguages) > 1
}

// GetModelType returns the model type.
func (v VocoderModel) GetModelType() ModelType {
	return v.modelType
}

// GetArchitecture returns the model architecture.
func (v VocoderModel) GetArchitecture() Architecture {
	return v.architecture
}

// GetDataset returns the model dataset.
func (v VocoderModel) GetDataset() Dataset {
	return v.dataset
}

// GetSupportedLanguages returns the supported languages.
func (v VocoderModel) GetSupportedLanguages() []Language {
	return v.supportedLanguages
}

// GetDefaultLanguage returns the default language.
func (v VocoderModel) GetDefaultLanguage() Language {
	return v.defaultLanguage
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
