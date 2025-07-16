package coqui

import (
	"slices"
	"strings"
)

// Language represents a supported language for TTS synthesis.
// Uses ISO 639-1 two-letter language codes (e.g., "en", "es", "fr").
// Most models don't accept regional variations (e.g., "en-US" vs "en-GB") because they are already built in. However, there are some models that do support regional variations.
// TTS models with "multilingual" support will extract the regional variation from the supplied speaker file.
type Language string

const (
	// Major languages
	English    Language = "en"
	Spanish    Language = "es"
	French     Language = "fr"
	German     Language = "de"
	Italian    Language = "it"
	Portuguese Language = "pt"
	Dutch      Language = "nl"
	Chinese    Language = "zh"
	Japanese   Language = "ja"

	// European languages
	Polish    Language = "pl"
	Turkish   Language = "tr"
	Russian   Language = "ru"
	Czech     Language = "cs"
	Ukrainian Language = "uk"
	Hungarian Language = "hu"
	Korean    Language = "ko"
	Arabic    Language = "ar"

	// Nordic languages
	Danish  Language = "da"
	Finnish Language = "fi"
	Swedish Language = "sv"

	// Baltic languages
	Estonian   Language = "et"
	Latvian    Language = "lv"
	Lithuanian Language = "lt"

	// Slavic languages
	Bulgarian Language = "bg"
	Croatian  Language = "hr"
	Slovak    Language = "sk"
	Slovenian Language = "sl"
	Romanian  Language = "ro"

	// Other European languages
	Greek   Language = "el"
	Irish   Language = "ga"
	Maltese Language = "mt"
	Catalan Language = "ca"

	// Asian languages
	Bengali Language = "bn"
	Persian Language = "fa"

	// African languages
	Ewe    Language = "ewe"
	Hausa  Language = "hau"
	Lin    Language = "lin"
	Yoruba Language = "yor"

	// Ghanaian Twi variants
	TwiAkuapem Language = "tw_akuapem"
	TwiAsante  Language = "tw_asante"

	// Eastern European
	Belarusian Language = "be"

	// Other
	Universal Language = "universal" // Represents a model that supports all languages
)

// allSupportedLanguages contains the full list of languages supported by the available Coqui TTS models.
// NOTE: Language support varies by model.
var allSupportedLanguages = []Language{
	English,
	Spanish,
	French,
	German,
	Italian,
	Portuguese,
	Dutch,
	Chinese,
	Japanese,
	Polish,
	Turkish,
	Russian,
	Czech,
	Ukrainian,
	Hungarian,
	Korean,
	Arabic,
	Danish,
	Finnish,
	Swedish,
	Estonian,
	Latvian,
	Lithuanian,
	Bulgarian,
	Croatian,
	Slovak,
	Slovenian,
	Romanian,
	Greek,
	Irish,
	Maltese,
	Catalan,
	Bengali,
	Persian,
	Ewe,
	Hausa,
	Lin,
	Yoruba,
	TwiAkuapem,
	TwiAsante,
	Belarusian,
}

// String returns the ISO 639-1 language code as a string.
func (l Language) String() string {
	return string(l)
}

// IsValid checks if the language is supported by Coqui TTS.
// Returns true for all languages in the allSupportedLanguages list.
func (l Language) IsValid() bool {
	return slices.Contains(allSupportedLanguages, l)
}

// ParseLanguage parses a language string and returns the corresponding Language.
// Accepts formats like "en-US", "en", "es-ES" and extracts the language code.
// Returns English as the default for unsupported or invalid languages.
// This function is useful for converting user input, configuration values, or extracted language values (like from an EPUB file) into a valid Language type.
func ParseLanguage(s string) Language {
	// TODO: There may be an exception to some languages that require specific handling such as "zh-CN" for Chinese.

	// Extract language code (before the "-")
	if idx := strings.Index(s, "-"); idx != -1 {
		s = s[:idx]
	}

	lang := Language(strings.ToLower(s))

	// Validate it's a supported language
	if lang.IsValid() {
		return lang
	}

	return English
}

// MustParseLanguage parses a language string and panics if invalid.
// Use this when you need to ensure the language is valid at initialisation time.
func MustParseLanguage(s string) Language {
	lang := ParseLanguage(s)
	if !lang.IsValid() {
		panic("Coqui TTS does not support language: " + s)
	}
	return lang
}

// GetAllSupportedLanguages returns a copy of all supported languages.
func GetAllSupportedLanguages() []Language {
	return append([]Language(nil), allSupportedLanguages...)
}
