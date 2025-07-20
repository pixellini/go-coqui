package vocoder

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVocoder(t *testing.T) {
	v, err := NewVocoder(BaseVocoderHifigan, model.DatasetVCTK, model.English)
	require.NoError(t, err, "NewVocoder should not return an error for valid input")
	assert.Equal(t, model.TypeVocoder, v.Category, "NewVocoder should set category to modelTypeVocoder")
}

func TestGetVocoders(t *testing.T) {
	models := GetVocoders()
	require.NotEmpty(t, models, "GetVocoders should not return an empty slice")
	assert.Equal(t, Vocoders.Models, models, "GetVocoders should return the predefined Vocoders slice")
}
