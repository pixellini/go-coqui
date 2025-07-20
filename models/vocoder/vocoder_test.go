package vocoder

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	v, err := New(Hifigan, model.DatasetVCTK, model.English)
	require.NoError(t, err, "New should not return an error for valid input")
	assert.Equal(t, model.TypeVocoder, v.Category, "New should set category to modelTypeVocoder")
}

func TestGetPresets(t *testing.T) {
	models := GetPresets()
	require.NotEmpty(t, models, "GetPresets should not return an empty slice")
	assert.Equal(t, presets.Models, models, "GetPresets should return the predefined presets slice")
}
