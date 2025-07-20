package coqui

import (
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/pixellini/go-coqui/models/tts"
	"github.com/pixellini/go-coqui/models/vocoder"
	"github.com/stretchr/testify/assert"
)

func TestNew_Default(t *testing.T) {
	coqui, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	assert.Equal(t, tts.PresetXTTSv2, coqui.CurrentModel(), "CurrentModel should return the default model")
	assert.Equal(t, vocoder.Model{}, coqui.CurrentVocoder(), "CurrentVocoder should return the default vocoder")
	assert.Equal(t, defaultLanguage, coqui.CurrentModelLanguage(), "CurrentModelLanguage should return the default language")
	assert.Equal(t, "", coqui.CurrentSpeakerSample(), "CurrentSpeakerSample should return the default speaker sample")
	assert.Equal(t, "", coqui.CurrentSpeakerIndex(), "CurrentSpeakerIndex should return the default speaker index")
	assert.Equal(t, defaultOutputDir, coqui.CurrentOutputDir(), "CurrentOutputDir should return the default output directory")
	assert.Equal(t, defaultDevice, coqui.CurrentDevice(), "CurrentDevice should return the default device")
	assert.Equal(t, defaultMaxRetries, coqui.CurrentMaxRetries(), "CurrentMaxRetries should return the default max retries")
}

func TestNewWithModelXttsV2(t *testing.T) {
	coqui, err := NewWithModelXttsV2()
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, coqui, "TTS instance should not be nil")
	assert.Equal(t, tts.PresetXTTSv2, coqui.CurrentModel(), "Current model should be XTTSv2")
}

func TestNewWithModelXttsV1(t *testing.T) {
	coqui, err := NewWithModelXttsV1()
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, coqui, "TTS instance should not be nil")
	assert.Equal(t, tts.PresetXTTSv1, coqui.CurrentModel(), "Current model should be XTTSv1")
}

func TestNewWithModelYourTTS(t *testing.T) {
	coqui, err := NewWithModelYourTTS()
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, coqui, "TTS instance should not be nil")
	assert.Equal(t, tts.PresetYourTTS, coqui.CurrentModel(), "Current model should be YourTTS")
}

func TestNewWithModelBark(t *testing.T) {
	coqui, err := NewWithModelBark()
	assert.NoError(t, err, "NewWithModel should not return an error for valid input")
	assert.NotNil(t, coqui, "TTS instance should not be nil")
	assert.Equal(t, tts.PresetBark, coqui.CurrentModel(), "Current model should be Bark")
}

func TestTTSName(t *testing.T) {
	coqui := &TTS{
		// A model that is not multilingual
		model: tts.PresetVITSFemaleCustom,
	}

	expectedName := "tts_models/bn/custom/vits-female"
	assert.Equal(t, expectedName, coqui.Name(), "TTS Name should match the expected value")
}

func TestTTSName_Multilingual(t *testing.T) {
	coqui, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	expectedName := "tts_models/multilingual/multi-dataset/xtts_v2"
	assert.Equal(t, expectedName, coqui.Name(), "TTS Name should match the expected value")
}

func TestTTSVocoderName(t *testing.T) {
	coqui := &TTS{
		vocoder: vocoder.PresetHifiganV2VCTK,
	}

	expectedVocoderName := "vocoder_models/en/vctk/hifigan_v2"
	assert.Equal(t, expectedVocoderName, coqui.VocoderName(), "TTS Vocoder Name should match the expected value")
}

func TestTTSSetters(t *testing.T) {
	coqui, err := New()
	assert.NoError(t, err, "New should not return an error for valid input")

	newModel := tts.PresetBark
	newModel.CurrentLanguage = model.English

	coqui.SetCurrentIdentifier(newModel)
	assert.Equal(t, newModel, coqui.CurrentModel(), "SetCurrentModel should update the current model")

	newVocoder := vocoder.Model{}
	coqui.SetCurrentVocoder(newVocoder)
	assert.Equal(t, newVocoder, coqui.CurrentVocoder(), "SetCurrentVocoder should update the current vocoder")

	coqui.SetCurrentModelLanguage(model.Portuguese)
	assert.Equal(t, model.Portuguese, coqui.CurrentModelLanguage(), "SetCurrentModelLanguage should update the current model language")

	coqui.SetCurrentSpeakerSample("sample.wav")
	assert.Equal(t, "sample.wav", coqui.CurrentSpeakerSample(), "SetCurrentSpeakerSample should update the current speaker sample")

	coqui.SetCurrentSpeakerIndex("speaker1")
	assert.Equal(t, "speaker1", coqui.CurrentSpeakerIndex(), "SetCurrentSpeakerIndex should update the current speaker index")

	coqui.SetCurrentOutputDir("/tmp/output")
	assert.Equal(t, "/tmp/output", coqui.CurrentOutputDir(), "SetCurrentOutputDir should update the current output directory")

	coqui.SetCurrentDevice(model.DeviceCPU)
	assert.Equal(t, model.DeviceCPU, coqui.CurrentDevice(), "SetCurrentDevice should update the current device")

	coqui.SetCurrentMaxRetries(5)
	assert.Equal(t, 5, coqui.CurrentMaxRetries(), "SetCurrentMaxRetries should update the current max retries")
}
