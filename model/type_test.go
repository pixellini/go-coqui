package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetModelTypes(t *testing.T) {
	tp := GetTypes()
	require.NotEmpty(t, tp, "GetModelTypes should not return an empty slice")
	assert.Equal(t, types, tp, "GetModelTypes should return the predefined modelTypes slice")
}
