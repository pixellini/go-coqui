package coqui

// ModelType represents the category of model.
type ModelType string

// Model interface defines the common behavior for all model types.
type Model interface {
	// String returns a string representation of the model identifier.
	String() string
	// IsValid checks if the model identifier is valid.
	IsValid() bool
	// IsMultilingual checks if the model supports multiple languages.
	IsMultilingual() bool
	// GetType returns the type of the model (e.g., TTS, Vocoder, Voice Conversion).
	GetModelType() ModelType
	// GetArchitecture returns the architecture of the model (e.g., Wavegrad, MelGAN).
	GetArchitecture() Architecture
	// GetDataset returns the dataset used by the model.
	GetDataset() Dataset
	// GetSupportedLanguages returns the languages supported by the model.
	GetSupportedLanguages() []Language
	// GetDefaultLanguage returns the default language of the model.
	GetDefaultLanguage() Language
}

const (
	// modelTypeTTS represents text-to-speech models.
	modelTypeTTS ModelType = "tts_models"
	// modelTypeVocoder represents vocoder models for audio synthesis.
	modelTypeVocoder ModelType = "vocoder_models"
	// modelTypeVoiceConversion represents voice conversion models.
	modelTypeVoiceConversion ModelType = "voice_conversion_models"
)

// allModelTypes contains all predefined model types.
var allModelTypes = []ModelType{
	modelTypeTTS,
	modelTypeVocoder,
	modelTypeVoiceConversion,
}

// GetAllModelTypes returns a slice of all predefined model types.
func GetAllModelTypes() []ModelType {
	return append([]ModelType(nil), allModelTypes...)
}

// Generic filter functions that work with any Model implementation.
// These are simple, readable, and work with both Model and Vocoder types.

// FilterModelsByType filters any slice of Model by type
func FilterModelsByType[T Model](models []T, modelType ModelType) []T {
	var result []T
	for _, model := range models {
		if model.GetModelType() == modelType {
			result = append(result, model)
		}
	}
	return result
}

// FilterModelsByArchitecture filters any slice of Model by architecture.
func FilterModelsByArchitecture[T Model](models []T, architecture Architecture) []T {
	var result []T
	for _, model := range models {
		if model.GetArchitecture() == architecture {
			result = append(result, model)
		}
	}
	return result
}

// FilterModelsByDataset filters any slice of Model by dataset.
func FilterModelsByDataset[T Model](models []T, dataset Dataset) []T {
	var result []T
	for _, model := range models {
		if model.GetDataset() == dataset {
			result = append(result, model)
		}
	}
	return result
}

// FilterModelsBySupportedLanguages filters models that support any of the specified languages.
func FilterModelsBySupportedLanguages[T Model](models []T, languages []Language) []T {
	var result []T
	for _, model := range models {
		modelSupported := model.GetSupportedLanguages()
		// Check if any of the model's supported languages match any of the requested languages.
		for _, supportedLang := range modelSupported {
			for _, requestedLang := range languages {
				if supportedLang == requestedLang {
					result = append(result, model)
					goto nextModel // Found a match, move to next model.
				}
			}
		}
	nextModel:
	}
	return result
}

// FilterModelsMultilingual returns all models that support multiple languages.
func FilterModelsMultilingual[T Model](models []T) []T {
	var result []T
	for _, model := range models {
		if model.IsMultilingual() {
			result = append(result, model)
		}
	}
	return result
}

// FilterModelsByDefaultLanguage filters models by their default language.
func FilterModelsByDefaultLanguage[T Model](models []T, language Language) []T {
	var result []T
	for _, model := range models {
		if model.GetDefaultLanguage() == language {
			result = append(result, model)
		}
	}
	return result
}
