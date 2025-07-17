package coqui

import (
	"fmt"
	"slices"
)

// I think we can completely remove this struct and use Model instead.
// They are basically the same thing and I have to keep repeating the same methods.
type VocoderModel struct {
	modelType          ModelType
	dataset            Dataset
	architecture       Architecture
	defaultLanguage    Language
	supportedLanguages []Language
}

// Predefined Vocoder Models for use with TTS synthesis.
// Vocoders convert mel-spectrograms to audio waveforms.
// Variables organized by architecture type for easier discovery.
var (
	// Wavegrad architecture vocoders.
	VocoderWavegradLibriTTS = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetLibriTTS,
		architecture:       ArchWavegrad,
		defaultLanguage:    English,
		supportedLanguages: allSupportedLanguages, // Universal vocoder.
	}

	VocoderWavegradEK1 = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetEK1,
		architecture:       ArchWavegrad,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderWavegradThorsten = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchWavegrad,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Fullband MelGAN architecture vocoders.
	VocoderFullbandMelganLibriTTS = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetLibriTTS,
		architecture:       ArchFullbandMelgan,
		defaultLanguage:    English,
		supportedLanguages: allSupportedLanguages, // Universal vocoder.
	}

	VocoderFullbandMelganThorsten = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchFullbandMelgan,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	// Multiband MelGAN architecture vocoders.
	VocoderMultibandMelganLJSpeech = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchMultibandMelgan,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderMultibandMelganMai = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetMai,
		architecture:       ArchMultibandMelgan,
		defaultLanguage:    Ukrainian,
		supportedLanguages: []Language{Ukrainian},
	}

	// HiFi-GAN v1 architecture vocoders.
	VocoderHifiganV1Thorsten = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetThorsten,
		architecture:       ArchHifiganV1,
		defaultLanguage:    German,
		supportedLanguages: []Language{German},
	}

	VocoderHifiganV1Kokoro = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetKokoro,
		architecture:       ArchHifiganV1,
		defaultLanguage:    Japanese,
		supportedLanguages: []Language{Japanese},
	}

	// HiFi-GAN v2 architecture vocoders.
	VocoderHifiganV2LJSpeech = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2Blizzard2013 = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetBlizzard2013,
		architecture:       ArchHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2VCTK = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetVCTK,
		architecture:       ArchHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	VocoderHifiganV2Sam = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetSam,
		architecture:       ArchHifiganV2,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// HiFi-GAN (generic) architecture vocoders.
	VocoderHifiganCommonVoiceTurkish = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetCommonVoice,
		architecture:       ArchHifigan,
		defaultLanguage:    Turkish,
		supportedLanguages: []Language{Turkish},
	}

	VocoderHifiganCommonVoiceBelarusian = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetCommonVoice,
		architecture:       ArchHifigan,
		defaultLanguage:    Belarusian,
		supportedLanguages: []Language{Belarusian},
	}

	// UnivNet architecture vocoders.
	VocoderUnivnetLJSpeech = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetLJSpeech,
		architecture:       ArchUnivnet,
		defaultLanguage:    English,
		supportedLanguages: []Language{English},
	}

	// Parallel WaveGAN architecture vocoders.
	VocoderParallelWaveganMai = VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            DatasetMai,
		architecture:       ArchParallelWavegan,
		defaultLanguage:    Dutch,
		supportedLanguages: []Language{Dutch},
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
func NewVocoder(language Language, dataset Dataset, architecture Architecture) VocoderModel {
	return VocoderModel{
		modelType:          modelTypeVocoder,
		dataset:            dataset,
		architecture:       architecture,
		defaultLanguage:    Language(language),
		supportedLanguages: []Language{Language(language)},
	}
}

// String returns a string representation of the model identifier.
// It will only print the default language.
// Will need to be updated if we want to include all supported languages.
func (v VocoderModel) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", v.GetModelType(), v.GetDataset(), v.GetDefaultLanguage(), v.GetArchitecture())
}

// NameList returns a list of string representations of the model identifier for each supported language.
func (v VocoderModel) NameList() []string {
	var names []string
	for _, lang := range v.supportedLanguages {
		name := fmt.Sprintf("%s/%s/%s/%s", v.GetModelType(), v.GetDataset(), lang, v.GetArchitecture())
		names = append(names, name)
	}
	return names
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

// GetDefaultLanguage returns the default language.
func (v VocoderModel) GetDefaultLanguage() Language {
	return v.defaultLanguage
}

// GetSupportedLanguages returns the supported languages.
func (v VocoderModel) GetSupportedLanguages() []Language {
	return v.supportedLanguages
}

// GetAllVocoders returns a copy of the list of all predefined vocoder models.
func GetAllVocoders() []VocoderModel {
	return slices.Clone(allVocoders)
}

// GetVocodersByArchitecture returns all vocoder models that use the specified architecture.
func GetVocodersByArchitecture(arch Architecture) []VocoderModel {
	return filterModelsByArchitecture(allVocoders, arch)
}

// GetVocodersByDataset returns all vocoder models of the specified dataset.
func GetVocodersByDataset(dataset Dataset) []VocoderModel {
	return filterModelsByDataset(allVocoders, dataset)
}

// GetVocodersByLanguage returns all vocoder models that support the specified language.
func GetVocodersByLanguage(language Language) []VocoderModel {
	return filterModelsBySupportedLanguages(allVocoders, []Language{language})
}

// GetVocodersBySupportedLanguages returns all vocoder models that support any of the specified languages.
func GetVocodersBySupportedLanguages(languages []Language) []VocoderModel {
	return filterModelsBySupportedLanguages(allVocoders, languages)
}

// GetVocodersByDefaultLanguage returns all vocoder models with the specified default language.
func GetVocodersByDefaultLanguage(language Language) []VocoderModel {
	return filterModelsByDefaultLanguage(allVocoders, language)
}

// GetVocodersMultilingual returns all vocoder models that support multiple languages.
func GetVocodersMultilingual() []VocoderModel {
	return filterModelsMultilingual(allVocoders)
}
