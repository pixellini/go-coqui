package coqui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Default(t *testing.T) {
	tts, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	assert.Equal(t, TTSModelXTTSv2, tts.CurrentModel(), "CurrentModel should return the default model")
	assert.Equal(t, Vocoder{}, tts.CurrentVocoder(), "CurrentVocoder should return the default vocoder")
	assert.Equal(t, defaultLanguage, tts.CurrentModelLanguage(), "CurrentModelLanguage should return the default language")
	assert.Equal(t, "", tts.CurrentSpeakerSample(), "CurrentSpeakerSample should return the default speaker sample")
	assert.Equal(t, "", tts.CurrentSpeakerIndex(), "CurrentSpeakerIndex should return the default speaker index")
	assert.Equal(t, defaultOutputDir, tts.CurrentOutputDir(), "CurrentOutputDir should return the default output directory")
	assert.Equal(t, defaultDevice, tts.CurrentDevice(), "CurrentDevice should return the default device")
	assert.Equal(t, defaultMaxRetries, tts.CurrentMaxRetries(), "CurrentMaxRetries should return the default max retries")
}

func TestNewWithModelXttsV2(t *testing.T) {
	samplePath := "sample.wav"
	tts, err := NewWithModelXttsV2(samplePath)
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, tts, "TTS instance should not be nil")
	assert.Equal(t, TTSModelXTTSv2, tts.CurrentModel(), "Current model should be XTTSv2")
}

func TestNewWithModelXttsV1(t *testing.T) {
	samplePath := "sample.wav"
	tts, err := NewWithModelXttsV1(samplePath)
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, tts, "TTS instance should not be nil")
	assert.Equal(t, TTSModelXTTSv1, tts.CurrentModel(), "Current model should be XTTSv1")
}

func TestNewWithModelYourTTS(t *testing.T) {
	samplePath := "sample.wav"
	tts, err := NewWithModelYourTTS(samplePath)
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, tts, "TTS instance should not be nil")
	assert.Equal(t, TTSModelYourTTS, tts.CurrentModel(), "Current model should be YourTTS")
}

func TestNewWithModelBark(t *testing.T) {
	samplePath := "sample.wav"
	tts, err := NewWithModelBark(samplePath)
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, tts, "TTS instance should not be nil")
	assert.Equal(t, TTSModelBark, tts.CurrentModel(), "Current model should be Bark")
}

func TestTTSName(t *testing.T) {
	tts := &TTS{
		// A model that is not multilingual
		model: TTSModelVITSFemaleCustom,
	}

	expectedName := "tts_models/bn/custom/vits-female"
	assert.Equal(t, expectedName, tts.Name(), "TTS Name should match the expected value")
}

func TestTTSName_Multilingual(t *testing.T) {
	tts, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	expectedName := "tts_models/multilingual/multi-dataset/xtts_v2"
	assert.Equal(t, expectedName, tts.Name(), "TTS Name should match the expected value")
}

func TestTTSVocoderName(t *testing.T) {
	tts := &TTS{
		vocoder: VocoderHifiganV2VCTK,
	}

	expectedVocoderName := "vocoder_models/en/vctk/hifigan_v2"
	assert.Equal(t, expectedVocoderName, tts.VocoderName(), "TTS Vocoder Name should match the expected value")
}

func TestTTSSetters(t *testing.T) {
	tts, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	newModel := TTSModelBark
	newModel.currentLanguage = English

	tts.SetCurrentModelIdentifier(newModel)
	assert.Equal(t, newModel, tts.CurrentModel(), "SetCurrentModel should update the current model")

	newVocoder := Vocoder{}
	tts.SetCurrentVocoder(newVocoder)
	assert.Equal(t, newVocoder, tts.CurrentVocoder(), "SetCurrentVocoder should update the current vocoder")

	tts.SetCurrentModelLanguage(Portuguese)
	assert.Equal(t, Portuguese, tts.CurrentModelLanguage(), "SetCurrentModelLanguage should update the current model language")

	tts.SetCurrentSpeakerSample("sample.wav")
	assert.Equal(t, "sample.wav", tts.CurrentSpeakerSample(), "SetCurrentSpeakerSample should update the current speaker sample")

	tts.SetCurrentSpeakerIndex("speaker1")
	assert.Equal(t, "speaker1", tts.CurrentSpeakerIndex(), "SetCurrentSpeakerIndex should update the current speaker index")

	tts.SetCurrentOutputDir("/tmp/output")
	assert.Equal(t, "/tmp/output", tts.CurrentOutputDir(), "SetCurrentOutputDir should update the current output directory")

	tts.SetCurrentDevice(DeviceCPU)
	assert.Equal(t, DeviceCPU, tts.CurrentDevice(), "SetCurrentDevice should update the current device")

	tts.SetCurrentMaxRetries(5)
	assert.Equal(t, 5, tts.CurrentMaxRetries(), "SetCurrentMaxRetries should update the current max retries")
}
