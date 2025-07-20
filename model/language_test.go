package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLanguageString(t *testing.T) {
	assert.Equal(t, "en", English.String(), "English.String() should return 'en'")
	assert.Equal(t, "es", Spanish.String(), "Spanish.String() should return 'es'")
	assert.Equal(t, "fr", French.String(), "French.String() should return 'fr'")
	assert.Equal(t, "tw_akuapem", TwiAkuapem.String(), "TwiAkuapem.String() should return 'tw_akuapem'")
}

func TestLanguageIsSupported(t *testing.T) {
	assert.True(t, English.IsSupported(), "English should be supported")
	assert.True(t, Spanish.IsSupported(), "Spanish should be supported")
	assert.True(t, Universal.IsSupported(), "Universal should be supported")
	assert.True(t, Multilingual.IsSupported(), "Multilingual should be supported")

	invalidLang := Language("invalid")
	assert.False(t, invalidLang.IsSupported(), "Invalid language should not be supported")

	emptyLang := Language("")
	assert.False(t, emptyLang.IsSupported(), "Empty language should not be supported")
}

func TestParseLanguage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Language
		hasError bool
	}{
		{
			name:     "simple English",
			input:    "en",
			expected: English,
			hasError: false,
		},
		{
			name:     "English with region",
			input:    "en-US",
			expected: English,
			hasError: false,
		},
		{
			name:     "Spanish with region",
			input:    "es-ES",
			expected: Spanish,
			hasError: false,
		},
		{
			name:     "uppercase input",
			input:    "EN",
			expected: English,
			hasError: false,
		},
		{
			name:     "mixed case with region",
			input:    "En-US",
			expected: English,
			hasError: false,
		},
		{
			name:     "French",
			input:    "fr",
			expected: French,
			hasError: false,
		},
		{
			name:     "German with region",
			input:    "de-DE",
			expected: German,
			hasError: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: Language(""),
			hasError: true,
		},
		{
			name:     "unsupported language",
			input:    "xyz",
			expected: Language(""),
			hasError: true,
		},
		{
			name:     "unsupported with region",
			input:    "xyz-ABC",
			expected: Language(""),
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseLanguage(tt.input)

			if tt.hasError {
				assert.Error(t, err, "return error")
			} else {
				assert.NoError(t, err, "not return error")
				assert.Equal(t, tt.expected, result, "return %v", tt.expected)
			}
		})
	}
}

func TestMustParseLanguage(t *testing.T) {
	result := MustParseLanguage("en")
	assert.Equal(t, English, result, "should return English")

	result = MustParseLanguage("es-ES")
	assert.Equal(t, Spanish, result, "should return Spanish")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseLanguage should panic for invalid input")
		}
	}()
	_ = MustParseLanguage("invalid")
}

func TestGetSupportedLanguages(t *testing.T) {
	languages := GetSupportedLanguages()
	require.NotEmpty(t, languages, "Supported languages should not be empty")
	assert.Equal(t, supportedLanguages, languages, "GetSupportedLanguages should return the predefined supportedLanguages slice")
}
