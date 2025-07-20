package voiceconversion

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVoiceConversion(t *testing.T) {
	vc, err := NewVoiceConversion(BaseVoiceConversionFreevc24, model.DatasetVCTK, model.English)
	require.NoError(t, err, "NewVoiceConversion should not return an error for valid input")
	assert.Equal(t, model.TypeVoiceConversion, vc.Category, "NewVoiceConversion should set category to modelTypeVoiceConversion")
}

func TestGetVoiceConversionModels(t *testing.T) {
	models := GetVoiceConversionModels()
	require.NotEmpty(t, models, "GetVoiceConversionModels should not return an empty slice")
	assert.Equal(t, VoiceConversions.Models, models, "GetVoiceConversionModels should return the predefined VoiceConversions slice")
}
