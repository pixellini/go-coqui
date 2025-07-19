package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetModelTypes(t *testing.T) {
	types := GetModelTypes()
	require.NotEmpty(t, types, "GetModelTypes should not return an empty slice")
	assert.Equal(t, modelTypes, types, "GetModelTypes should return the predefined modelTypes slice")
}
