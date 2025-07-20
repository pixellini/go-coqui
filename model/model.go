package model

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
	// GetType returns the type of the model (e.g., TTS, Vocoder, Voice Conversion).
	GetType() Type
	// GetModel returns the architecture of the model (e.g., Wavegrad, MelGAN).
	GetBaseModel() BaseModel
	// GetDataset returns the dataset used by the model.
	GetDataset() Dataset
	// GetDefaultLanguage returns the default language of the model.
	GetDefaultLanguage() Language
	// GetSupportedLanguages returns the languages supported by the model.
	GetSupportedLanguages() []Language
}

// Identifier represents a comprehensive model identifier.
type Identifier struct {
	// category of the model (e.g., TTS, Vocoder, Voice Conversion).
	Category Type
	// dataset used by the model (e.g., LJSpeech, VCTK, Common Voice).
	Dataset Dataset
	// architecture of the model (e.g., Tacotron2, VITS, GlowTTS).
	Model BaseModel
	// defaultLanguage is the primary language for this model.
	// There isn't any documentation on which language is the default for each model,
	// so I assume the default language is English, or the first language in the supportedLanguages list.
	DefaultLanguage Language
	// supportedLanguages lists all languages this model supports.
	SupportedLanguages []Language
	// currentLanguage is the language currently set for this model.
	CurrentLanguage Language
	// supportsVoiceCloning indicates if the model supports voice cloning by providing a speaker sample.
	SupportsVoiceCloning bool
	// isCustom Indicates if this is a custom model not predefined in the library
	IsCustom bool
}

// NewModel creates a new custom Model Identifier.
// This is useful for models that are not predefined in the Coqui TTS library.
func NewModel(t Type, m BaseModel, d Dataset, l Language) (Identifier, error) {
	if t == "" {
		return Identifier{}, fmt.Errorf("model type cannot be empty")
	}
	if l == "" {
		return Identifier{}, fmt.Errorf("language cannot be empty")
	}
	if d == "" {
		return Identifier{}, fmt.Errorf("dataset cannot be empty")
	}
	if m == "" {
		return Identifier{}, fmt.Errorf("model architecture cannot be empty")
	}

	if !slices.Contains(types, t) {
		return Identifier{}, fmt.Errorf("unsupported model type: %s", t)
	}
	if !l.IsSupported() {
		return Identifier{}, fmt.Errorf("unsupported language: %s", l)
	}
	if !d.isPreset() {
		return Identifier{}, fmt.Errorf("unsupported dataset: %s", d)
	}

	var supportedLanguages = []Language{l}
	if l == Universal || l == Multilingual {
		// If the language is Universal or Multilingual, we assume it supports all languages.
		supportedLanguages = GetSupportedLanguages()
	}

	// TODO: Need a method to point to a custom model via a path.
	// IDEA: It would be cool to add a method that splits the Model Name from a string like "tts_models/en/ek1/tacotron2" into its components.
	return Identifier{
		Category:             t,
		Dataset:              d,
		Model:                m,
		DefaultLanguage:      l,
		CurrentLanguage:      l,
		SupportedLanguages:   supportedLanguages,
		SupportsVoiceCloning: false, // Default to false, can be set later if needed.
		IsCustom:             true,
	}, nil
}

// Name returns a string representation of the model identifier.
// It formats the model identifier as "category/dataset/language/model".
func (id Identifier) Name() string {
	return fmt.Sprintf("%s/%s/%s/%s", id.GetType(), id.GetCurrentLanguage(), id.GetDataset(), id.GetBaseModel())
}

// NameList returns a list of string representations of the model identifier for each supported language.
func (id Identifier) NameList() []string {
	var names []string
	for _, lang := range id.SupportedLanguages {
		name := fmt.Sprintf("%s/%s/%s/%s", id.GetType(), id.GetDataset(), lang, id.GetBaseModel())
		names = append(names, name)
	}
	return names
}

// IsValid checks if the model identifier is valid.
func (id Identifier) IsValid() bool {
	return id.Validate() == nil
}

// Validate checks if the model identifier is valid and returns an error explaining why it's invalid.
func (id Identifier) Validate() error {
	// Assume it's valid if it's a custom model provided by the user.
	if id.IsCustom {
		return nil
	}
	if !slices.Contains(types, id.Category) {
		return fmt.Errorf("unsupported model type: %s", id.Category)
	}
	if id.Model == "" {
		return fmt.Errorf("architecture cannot be empty for model type: %s", id.Category)
	}
	if id.Dataset == "" {
		return fmt.Errorf("dataset cannot be empty for model type: %s", id.Category)
	}
	if len(id.SupportedLanguages) == 0 {
		return fmt.Errorf("supported languages cannot be empty for model type: %s", id.Category)
	}

	return nil
}

// IsMultilingual returns true if the effective model supports multiple languages.
func (id Identifier) IsMultilingual() bool {
	return len(id.SupportedLanguages) > 1
}

// SupportsLanguage checks if the model supports the specified language.
func (id Identifier) SupportsLanguage(lang Language) bool {
	return slices.Contains(id.SupportedLanguages, lang)
}

// SupportsVoiceCloning checks if the model supports voice cloning by providing a speaker sample.
func (id Identifier) SupportsCloning() bool {
	if id.IsCustom {
		return true // Custom models are assumed to support voice cloning.
	}
	return id.SupportsVoiceCloning
}

// GetType returns the model type.
func (id Identifier) GetType() Type {
	return id.Category
}

// GetBaseModel returns the model architecture.
func (id Identifier) GetBaseModel() BaseModel {
	return id.Model
}

// GetDataset returns the model dataset.
func (id Identifier) GetDataset() Dataset {
	return id.Dataset
}

// GetCurrentLanguage returns the current language set for the model.
func (id Identifier) GetCurrentLanguage() Language {
	return id.CurrentLanguage
}

// GetDefaultLanguage returns the default language of the model.
// There isn't any documentation on which language is the default for each model,
// so I assume the default language is English, or the first language in the supportedLanguages list.
func (id Identifier) GetDefaultLanguage() Language {
	return id.DefaultLanguage
}

// GetSupportedLanguages returns the supported languages.
func (id Identifier) GetSupportedLanguages() []Language {
	return id.SupportedLanguages
}

type ModelList[T Model] struct {
	Models []T
}

// FilterByBaseModel filters any slice of Model by architecture.
func (m *ModelList[T]) FilterByBaseModel(baseModel BaseModel) *ModelList[T] {
	var result []T
	for _, model := range m.Models {
		if model.GetBaseModel() == baseModel {
			result = append(result, model)
		}
	}
	return &ModelList[T]{Models: result}
}

// FilterByDataset filters any slice of Model by dataset.
func (m *ModelList[T]) FilterByDataset(dataset Dataset) []T {
	var result []T
	for _, model := range m.Models {
		if model.GetDataset() == dataset {
			result = append(result, model)
		}
	}
	return result
}

// FilterBySupportedLanguages filters models that support any of the specified languages.
func (m *ModelList[T]) FilterBySupportedLanguages(languages []Language) []T {
	var result []T
	for _, model := range m.Models {
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
	for _, model := range m.Models {
		if model.IsMultilingual() {
			result = append(result, model)
		}
	}
	return result
}

// FilterByDefaultLanguage filters models by their default language.
func (m *ModelList[T]) FilterByDefaultLanguage(language Language) []T {
	var result []T
	for _, model := range m.Models {
		if model.GetDefaultLanguage() == language {
			result = append(result, model)
		}
	}
	return result
}
