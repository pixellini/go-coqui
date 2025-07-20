package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVoiceConversion(t *testing.T) {
	vc, err := NewVoiceConversion(BaseVoiceConversionFreevc24, DatasetVCTK, English)
	require.NoError(t, err, "NewVoiceConversion should not return an error for valid input")
	assert.Equal(t, modelTypeVoiceConversion, vc.category, "NewVoiceConversion should set category to modelTypeVoiceConversion")
}

func TestGetVoiceConversionModels(t *testing.T) {
	models := GetVoiceConversionModels()
	require.NotEmpty(t, models, "GetVoiceConversionModels should not return an empty slice")
	assert.Equal(t, VoiceConversions.models, models, "GetVoiceConversionModels should return the predefined VoiceConversions slice")
}
