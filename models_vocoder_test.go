package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVocoder(t *testing.T) {
	v, err := NewVocoder(English, DatasetVCTK, BaseVoiceConversionFreevc24)
	require.NoError(t, err, "NewVocoder should not return an error for valid input")
	assert.Equal(t, modelTypeVocoder, v.category, "NewVocoder should set category to modelTypeVocoder")
}

func TestGetVocoders(t *testing.T) {
	models := GetVocoders()
	require.NotEmpty(t, models, "GetVocoders should not return an empty slice")
	assert.Equal(t, Vocoders.models, models, "GetVocoders should return the predefined Vocoders slice")
}
