package model

import (
	"os/exec"
	"runtime"
	"slices"
)

// Device represents the compute device for TTS synthesis.
// Supported devices include CPU, CUDA (NVIDIA GPU), MPS (Apple Silicon), and auto-detection.
// TODO: Add support for other devices like ipu, opengl, etc. in the future.
type Device string

const (
	// DeviceAuto enables automatic device detection based on available hardware.
	DeviceAuto Device = "auto"
	// DeviceCPU forces CPU-only synthesis (slowest but most compatible).
	DeviceCPU Device = "cpu"
	// DeviceCUDA enables NVIDIA GPU acceleration (requires CUDA installation).
	DeviceCUDA Device = "cuda"
	// DeviceMPS enables Apple Silicon GPU acceleration (macOS only).
	DeviceMPS Device = "mps"
)

// AllDevices contains all predefined device types supported by Coqui TTS.
var devices = []Device{
	DeviceAuto,
	DeviceCPU,
	DeviceCUDA,
	DeviceMPS,
}

// String returns the string representation of the Device.
func (d Device) String() string {
	return string(d)
}

// IsValid checks if the device type is supported.
func (d Device) IsValid() bool {
	return slices.Contains(devices, d)
}

// GetDevices returns a list of all predefined device types supported by Coqui TTS.
func GetDevices() []Device {
	return slices.Clone(devices)
}

// isRuntimeGOOS checks if the current operating system is macOS.
func isRuntimeGOOS() bool {
	return runtime.GOOS == "darwin"
}

// isRuntimeGOARCH checks if the current model is ARM64 (Apple Silicon).
func isRuntimeGOARCH() bool {
	return runtime.GOARCH == "arm64"
}

// isCudaAvailable checks if NVIDIA GPU and drivers are available.
// This performs a basic check by looking for the nvidia-smi command.
// Note: This does not guarantee CUDA is properly configured for TTS.
func isCudaAvailable() bool {
	_, err := exec.LookPath("nvidia-smi")
	return err == nil
}

// detectDevice automatically selects the best available compute device.
// Priority order: CUDA (if available) > MPS (macOS ARM64) > CPU (fallback).
// NOTE: I don't know if we should even have this function.
// It might be better/easier to let the user explicitly set the device.
// When I implement more devices, this function could become more complex.
// I'll make that decision later.
func DetectDevice() Device {
	if isCudaAvailable() {
		return DeviceCUDA
	}

	if isRuntimeGOOS() && isRuntimeGOARCH() {
		return DeviceMPS
	}

	// Default to "cpu" if no other device is available.
	return DeviceCPU
}
