package voiceconversion

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	vc, err := New(Freevc24, model.DatasetVCTK, model.English)
	require.NoError(t, err, "NewVoiceConversion should not return an error for valid input")
	assert.Equal(t, model.TypeVoiceConversion, vc.Category, "NewVoiceConversion should set category to modelTypeVoiceConversion")
}

func TestGetPresets(t *testing.T) {
	models := GetPresets()
	require.NotEmpty(t, models, "GetPresets should not return an empty slice")
	assert.Equal(t, presets.Models, models, "GetPresets should return the predefined VoiceConversions slice")
}
