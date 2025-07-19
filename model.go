package coqui

import (
	"fmt"
	"slices"
)

type BaseModel string

// Model interface defines the common behavior for all model types.
type Model interface {
	Name() string
	// NameList returns a list of string representations of the model identifier for each supported language.
	NameList() []string
	// IsValid checks if the model identifier is valid.
	IsValid() bool
	// Validate checks if the model identifier is valid and returns an error explaining why it's invalid.
	Validate() error
	// IsMultilingual checks if the model supports multiple languages.
	IsMultilingual() bool
	// GetCategory returns the type of the model (e.g., TTS, Vocoder, Voice Conversion).
	GetCategory() Category
	// GetModel returns the architecture of the model (e.g., Wavegrad, MelGAN).
	GetBaseModel() BaseModel
	// GetDataset returns the dataset used by the model.
	GetDataset() Dataset
	// GetDefaultLanguage returns the default language of the model.
	GetDefaultLanguage() Language
	// GetSupportedLanguages returns the languages supported by the model.
	GetSupportedLanguages() []Language
}

// ModelIdentifier represents a comprehensive model identifier.
type ModelIdentifier struct {
	// category of the model (e.g., TTS, Vocoder, Voice Conversion).
	category Category
	// dataset used by the model (e.g., LJSpeech, VCTK, Common Voice).
	dataset Dataset
	// architecture of the model (e.g., Tacotron2, VITS, GlowTTS).
	model BaseModel
	// defaultLanguage is the primary language for this model.
	// There isn't any documentation on which language is the default for each model,
	// so I assume the default language is English, or the first language in the supportedLanguages list.
	defaultLanguage Language
	// supportedLanguages lists all languages this model supports.
	supportedLanguages []Language
	// currentLanguage is the language currently set for this model.
	currentLanguage Language
	// supportsVoiceCloning indicates if the model supports voice cloning by providing a speaker sample.
	supportsVoiceCloning bool
	// isCustom Indicates if this is a custom model not predefined in the library
	isCustom bool
}

// NewModal creates a new custom Model Identifier.
// This is useful for models that are not predefined in the Coqui TTS library.
func NewModal(modelType Category, language Language, dataset Dataset, model BaseModel) (ModelIdentifier, error) {
	if modelType == "" {
		return ModelIdentifier{}, fmt.Errorf("model type cannot be empty")
	}
	if language == "" {
		return ModelIdentifier{}, fmt.Errorf("language cannot be empty")
	}
	if dataset == "" {
		return ModelIdentifier{}, fmt.Errorf("dataset cannot be empty")
	}
	if model == "" {
		return ModelIdentifier{}, fmt.Errorf("model architecture cannot be empty")
	}

	var supportedLanguages = []Language{language}
	if language == Universal || language == Multilingual {
		// If the language is Universal or Multilingual, we assume it supports all languages.
		supportedLanguages = GetSupportedLanguages()
	}

	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return ModelIdentifier{
		category:             modelType,
		dataset:              dataset,
		model:                model,
		defaultLanguage:      language,
		currentLanguage:      language,
		supportedLanguages:   supportedLanguages,
		supportsVoiceCloning: false, // Default to false, can be set later if needed.
		isCustom:             true,
	}, nil
}

// Name returns a string representation of the model identifier.
// It formats the model identifier as "category/dataset/language/model".
func (m ModelIdentifier) Name() string {
	return fmt.Sprintf("%s/%s/%s/%s", m.GetCategory(), m.GetCurrentLanguage(), m.GetDataset(), m.GetBaseModel())
}

// NameList returns a list of string representations of the model identifier for each supported language.
func (m ModelIdentifier) NameList() []string {
	var names []string
	for _, lang := range m.supportedLanguages {
		name := fmt.Sprintf("%s/%s/%s/%s", m.GetCategory(), m.GetDataset(), lang, m.GetBaseModel())
		names = append(names, name)
	}
	return names
}

// IsValid checks if the model identifier is valid.
func (m ModelIdentifier) IsValid() bool {
	return m.Validate() == nil
}

// Validate checks if the model identifier is valid and returns an error explaining why it's invalid.
func (m ModelIdentifier) Validate() error {
	// Assume it's valid if it's a custom model provided by the user.
	if m.isCustom {
		return nil
	}
	if !slices.Contains(modelTypes, m.category) {
		return fmt.Errorf("unsupported model type: %s", m.category)
	}
	if m.model == "" {
		return fmt.Errorf("architecture cannot be empty for model type: %s", m.category)
	}
	if m.dataset == "" {
		return fmt.Errorf("dataset cannot be empty for model type: %s", m.category)
	}
	if len(m.supportedLanguages) == 0 {
		return fmt.Errorf("supported languages cannot be empty for model type: %s", m.category)
	}

	return nil
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (m ModelIdentifier) IsMultilingual() bool {
	return len(m.supportedLanguages) > 1
}

// SupportsLanguage checks if the model supports the specified language.
func (m ModelIdentifier) SupportsLanguage(lang Language) bool {
	return slices.Contains(m.supportedLanguages, lang)
}

// SupportsVoiceCloning checks if the model supports voice cloning by providing a speaker sample.
func (m ModelIdentifier) SupportsVoiceCloning() bool {
	if m.isCustom {
		return true // Custom models are assumed to support voice cloning.
	}
	return m.supportsVoiceCloning
}

// GetCategory returns the model type.
func (m ModelIdentifier) GetCategory() Category {
	return m.category
}

// GetBaseModel returns the model architecture.
func (m ModelIdentifier) GetBaseModel() BaseModel {
	return m.model
}

// GetDataset returns the model dataset.
func (m ModelIdentifier) GetDataset() Dataset {
	return m.dataset
}

// GetCurrentLanguage returns the current language set for the model.
func (m ModelIdentifier) GetCurrentLanguage() Language {
	return m.currentLanguage
}

// GetDefaultLanguage returns the default language of the model.
// There isn't any documentation on which language is the default for each model,
// so I assume the default language is English, or the first language in the supportedLanguages list.
func (m ModelIdentifier) GetDefaultLanguage() Language {
	return m.defaultLanguage
}

// GetSupportedLanguages returns the supported languages.
func (m ModelIdentifier) GetSupportedLanguages() []Language {
	return m.supportedLanguages
}

type ModelList[T Model] struct {
	models []T
}

// FilterByBaseModel filters any slice of Model by architecture.
func (m *ModelList[T]) FilterByBaseModel(baseModel BaseModel) *ModelList[T] {
	var result []T
	for _, model := range m.models {
		if model.GetBaseModel() == baseModel {
			result = append(result, model)
		}
	}
	return &ModelList[T]{models: result}
}

// FilterByDataset filters any slice of Model by dataset.
func (m *ModelList[T]) FilterByDataset(dataset Dataset) []T {
	var result []T
	for _, model := range m.models {
		if model.GetDataset() == dataset {
			result = append(result, model)
		}
	}
	return result
}

// FilterBySupportedLanguages filters models that support any of the specified languages.
func (m *ModelList[T]) FilterBySupportedLanguages(languages []Language) []T {
	var result []T
	for _, model := range m.models {
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

// FilterByMultilingual returns all models that support multiple languages.
func (m *ModelList[T]) FilterByMultilingual() []T {
	var result []T
	for _, model := range m.models {
		if model.IsMultilingual() {
			result = append(result, model)
		}
	}
	return result
}

// FilterByDefaultLanguage filters models by their default language.
func (m *ModelList[T]) FilterByDefaultLanguage(language Language) []T {
	var result []T
	for _, model := range m.models {
		if model.GetDefaultLanguage() == language {
			result = append(result, model)
		}
	}
	return result
}
