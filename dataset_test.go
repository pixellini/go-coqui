package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatasetString(t *testing.T) {
	assert.Equal(t, "custom", DatasetCustom.String(), "DatasetCustom.String() should return 'custom'")
}

func TestDatasetIsPreset(t *testing.T) {
	assert.True(t, DatasetCustom.isPreset(), "DatasetCustom should be a valid preset dataset")

	emptyDataset := Dataset("")
	assert.False(t, emptyDataset.isPreset(), "Empty dataset should not be a valid preset dataset")
}

func TestGetDatasets(t *testing.T) {
	d := GetDatasets()
	require.NotEmpty(t, d, "GetDatasets should not return an empty slice")
	assert.Equal(t, datasets, d, "GetDatasets should return the predefined datasets slice")
}
