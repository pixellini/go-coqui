package tts

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	ttsModel, err := New(Bark, model.DatasetVCTK, model.English)
	require.NoError(t, err, "NewTTSModel should not return an error for valid input")
	assert.Equal(t, model.TypeTTS, ttsModel.Category, "NewTTSModel should set category to modelTypeTTS")
}

func TestGetPresets(t *testing.T) {
	models := GetPresets()
	require.NotEmpty(t, models, "GetPresets should not return an empty slice")
	assert.Equal(t, presets.Models, models, "GetPresets should return the predefined TTSModels slice")
}
