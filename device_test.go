package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeviceString(t *testing.T) {
	assert.Equal(t, "cpu", DeviceCPU.String(), "DeviceCPU.String() should return 'cpu'")
	assert.Equal(t, "cuda", DeviceCUDA.String(), "DeviceCUDA.String() should return 'cuda'")
	assert.Equal(t, "mps", DeviceMPS.String(), "DeviceMPS.String() should return 'mps'")
	assert.Equal(t, "auto", DeviceAuto.String(), "DeviceAuto.String() should return 'auto'")
}

func TestDeviceIsValid(t *testing.T) {
	assert.True(t, DeviceCPU.IsValid(), "DeviceCPU should be valid")
	assert.True(t, DeviceCUDA.IsValid(), "DeviceCUDA should be valid")
	assert.True(t, DeviceMPS.IsValid(), "DeviceMPS should be valid")
	assert.True(t, DeviceAuto.IsValid(), "DeviceAuto should be valid")

	invalidDevice := Device("invalid-device")
	assert.False(t, invalidDevice.IsValid(), "Invalid device should not be valid")

	emptyDevice := Device("")
	assert.False(t, emptyDevice.IsValid(), "Empty device should not be valid")
}

func TestGetDevices(t *testing.T) {
	d := GetDevices()
	require.NotEmpty(t, d, "GetDevices should not return an empty slice")
	assert.Equal(t, devices, d, "GetDevices should return the predefined devices slice")
}

func TestIsRuntimeGOOS(t *testing.T) {
	result := isRuntimeGOOS()
	_ = result
}

func TestIsRuntimeGOARCH(t *testing.T) {
	result := isRuntimeGOARCH()
	_ = result
}

func TestIsCudaAvailable(t *testing.T) {
	result := isCudaAvailable()
	_ = result
}

func TestDetectDevice(t *testing.T) {
	device := detectDevice()
	require.True(t, device.IsValid(), "detectDevice should return a valid device")
	validDetectedDevices := []Device{DeviceCPU, DeviceCUDA, DeviceMPS}
	assert.Contains(t, validDetectedDevices, device, "detectDevice returned unexpected device: %s", device)
}
