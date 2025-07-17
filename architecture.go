package coqui

import "slices"

// Architecture represents the underlying model architecture.
type Architecture string

const (
	// Multilingual architectures.
	ArchXTTSv2  Architecture = "xtts_v2"
	ArchXTTSv1  Architecture = "xtts_v1.1"
	ArchYourTTS Architecture = "your_tts"
	ArchBark    Architecture = "bark"

	// Single language architectures.
	ArchVITS         Architecture = "vits"
	ArchVITSNeon     Architecture = "vits--neon"
	ArchVITSNeonDash Architecture = "vits-neon"
	ArchVITSMale     Architecture = "vits-male"
	ArchVITSFemale   Architecture = "vits-female"

	// Tacotron variants.
	ArchTacotron2       Architecture = "tacotron2"
	ArchTacotron2DDC    Architecture = "tacotron2-DDC"
	ArchTacotron2DDCPh  Architecture = "tacotron2-DDC_ph"
	ArchTacotron2DCA    Architecture = "tacotron2-DCA"
	ArchTacotron2DDCGST Architecture = "tacotron2-DDC-GST"

	// Other architectures.
	ArchGlowTTS            Architecture = "glow-tts"
	ArchFastPitch          Architecture = "fast_pitch"
	ArchSpeedySpeech       Architecture = "speedy-speech"
	ArchOverflow           Architecture = "overflow"
	ArchNeuralHMM          Architecture = "neural_hmm"
	ArchTortoise           Architecture = "tortoise-v2"
	ArchCapacitronT2C50    Architecture = "capacitron-t2-c50"
	ArchCapacitronT2C150v2 Architecture = "capacitron-t2-c150_v2"
	ArchJenny              Architecture = "jenny"

	// Vocoder architectures.
	ArchWavegrad        Architecture = "wavegrad"
	ArchFullbandMelgan  Architecture = "fullband-melgan"
	ArchMultibandMelgan Architecture = "multiband-melgan"
	ArchHifiganV1       Architecture = "hifigan_v1"
	ArchHifiganV2       Architecture = "hifigan_v2"
	ArchHifigan         Architecture = "hifigan"
	ArchUnivnet         Architecture = "univnet"
	ArchParallelWavegan Architecture = "parallel-wavegan"

	// Voice conversion architectures.
	ArchFreevc24 Architecture = "freevc24"
)

var allArchitectures = []Architecture{
	// Multilingual architectures.
	ArchXTTSv2,
	ArchXTTSv1,
	ArchYourTTS,
	ArchBark,

	// Single language architectures.
	ArchVITS,
	ArchVITSNeon,
	ArchVITSNeonDash,
	ArchVITSMale,
	ArchVITSFemale,

	// Tacotron variants.
	ArchTacotron2,
	ArchTacotron2DDC,
	ArchTacotron2DDCPh,
	ArchTacotron2DCA,
	ArchTacotron2DDCGST,

	// Other architectures.
	ArchGlowTTS,
	ArchFastPitch,
	ArchSpeedySpeech,
	ArchOverflow,
	ArchNeuralHMM,
	ArchTortoise,
	ArchCapacitronT2C50,
	ArchCapacitronT2C150v2,
	ArchJenny,

	// Vocoder architectures.
	ArchWavegrad,
	ArchFullbandMelgan,
	ArchMultibandMelgan,
	ArchHifiganV1,
	ArchHifiganV2,
	ArchHifigan,
	ArchUnivnet,
	ArchParallelWavegan,

	// Voice conversion architectures.
	ArchFreevc24,
}

// Strings returns the architecture name as a string.
func (a Architecture) String() string {
	return string(a)
}

// IsValid checks if the architecture is one of the predefined architectures.
func (a Architecture) IsValid() bool {
	return slices.Contains(allArchitectures, a)
}

// GetAllModelArchitectures returns a list of all predefined architectures.
func GetAllModelArchitectures() []Architecture {
	return slices.Clone(allArchitectures)
}
