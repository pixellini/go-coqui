package tts

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTTSModel(t *testing.T) {
	ttsModel, err := NewTTSModel(BaseModelBark, model.DatasetVCTK, model.English)
	require.NoError(t, err, "NewTTSModel should not return an error for valid input")
	assert.Equal(t, model.TypeTTS, ttsModel.Category, "NewTTSModel should set category to modelTypeTTS")
}

func TestGetTTSModels(t *testing.T) {
	models := GetTTSModels()
	require.NotEmpty(t, models, "GetTTSModels should not return an empty slice")
	assert.Equal(t, TTSModels.Models, models, "GetTTSModels should return the predefined TTSModels slice")
}
