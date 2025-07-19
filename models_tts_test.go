package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTTSModel(t *testing.T) {
	v, err := NewTTSModel(BaseVoiceConversionFreevc24, DatasetVCTK, English)
	require.NoError(t, err, "NewTTSModel should not return an error for valid input")
	assert.Equal(t, modelTypeTTS, v.category, "NewTTSModel should set category to modelTypeTTS")
}

func TestGetTTSModels(t *testing.T) {
	models := GetTTSModels()
	require.NotEmpty(t, models, "GetTTSModels should not return an empty slice")
	assert.Equal(t, TTSModels.models, models, "GetTTSModels should return the predefined TTSModels slice")
}
