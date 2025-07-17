package coqui

import (
	"fmt"
	"os/exec"
	"runtime"
	"slices"
)

// Device represents the compute device for TTS synthesis.
// Supported devices include CPU, CUDA (NVIDIA GPU), MPS (Apple Silicon), and auto-detection.
// TODO: Add support for other devices like ipu, opengl, etc. in the future.
type Device string

// Runtime represents system runtime characteristics.
// Used internally for device detection and platform-specific optimisations.
type DeviceRuntime string

const (
	// DeviceAuto enables automatic device detection based on available hardware.
	DeviceAuto Device = "auto"
	// DeviceCPU forces CPU-only synthesis (slowest but most compatible).
	DeviceCPU Device = "cpu"
	// DeviceCUDA enables NVIDIA GPU acceleration (requires CUDA installation).
	DeviceCUDA Device = "cuda"
	// DeviceMPS enables Apple Silicon GPU acceleration (macOS only).
	DeviceMPS Device = "mps"

	// DeviceRuntimeDarwin represents the macOS operating system.
	DeviceRuntimeDarwin DeviceRuntime = "darwin"
	// DeviceRuntimeARM64 represents ARM64 processor architecture (Apple Silicon).
	DeviceRuntimeARM64 DeviceRuntime = "arm64"
)

// AllDevices contains all predefined device types supported by Coqui TTS.
var allDevices = []Device{
	DeviceAuto,
	DeviceCPU,
	DeviceCUDA,
	DeviceMPS,
}

// AllDeviceRuntimes contains all predefined device runtime types supported by Coqui TTS.
var allDeviceRuntimes = []DeviceRuntime{
	DeviceRuntimeDarwin,
	DeviceRuntimeARM64,
}

// String returns the string representation of the Device.
func (d Device) String() string {
	return string(d)
}

// IsValid checks if the device type is supported.
func (d Device) IsValid() bool {
	return slices.Contains(allDevices, d)
}

// String returns the string representation of the Runtime.
func (r DeviceRuntime) String() string {
	return string(r)
}

// IsValid checks if the runtime type is supported.
func (r DeviceRuntime) IsValid() bool {
	return slices.Contains(allDeviceRuntimes, r)
}

// DetectDevice automatically selects the best available compute device.
// Priority order: CUDA (if available) > MPS (macOS ARM64) > CPU (fallback).
// NOTE: I don't know if we should even have this function.
// It might be better/easier to let the user explicitly set the device.
// When I implement more devices, this function could become more complex.
// I'll make that decision later.
func DetectDevice(d Device) Device {
	if isCudaAvailable() {
		fmt.Println("CUDA available, using CUDA device.")
		return DeviceCUDA
	}

	if runtime.GOOS == string(DeviceRuntimeDarwin) && runtime.GOARCH == string(DeviceRuntimeARM64) {
		fmt.Println("ARM64 architecture detected, using MPS.")
		return DeviceMPS
	}

	// Default to "cpu" if no other device is available.
	fmt.Println("No GPU detected or CUDA not available, using CPU device.")
	return DeviceCPU
}

// isCudaAvailable checks if NVIDIA GPU and drivers are available.
// This performs a basic check by looking for the nvidia-smi command.
// Note: This does not guarantee CUDA is properly configured for TTS.
func isCudaAvailable() bool {
	_, err := exec.LookPath("nvidia-smi")
	return err == nil
}
