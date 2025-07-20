package coqui

import (
	"os"
	"testing"

	"github.com/pixellini/go-coqui/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var MockDataset = model.Dataset("mock-dataset")
var MockBaseModel = model.BaseModel("mock-base-model")
var MockOutputDir = "/tmp/output"
var MockModelId = model.Identifier{
	Category:             model.TypeTTS,
	Dataset:              MockDataset,
	Model:                MockBaseModel,
	DefaultLanguage:      model.English,
	CurrentLanguage:      model.English,
	SupportedLanguages:   []model.Language{model.English, model.German, model.French},
	SupportsVoiceCloning: true,
}

func TestOptions(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test-model-*.ckpt")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(tmpFile.Name()) })

	tests := []struct {
		name   string
		option Option
		check  func(*testing.T, *TTS)
		pre    func(*TTS)
	}{
		{
			name:   "WithModelId",
			option: WithModelId(MockModelId),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, MockModelId, tts.model, "WithModelId should set the model field")
			},
		},
		{
			name:   "WithModelPath",
			option: WithModelPath(tmpFile.Name()),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, tmpFile.Name(), tts.modelPath, "WithModelPath should set the modelPath field")
			},
		},
		{
			name:   "WithModelLanguage",
			option: WithModelLanguage(model.English),
			pre: func(tts *TTS) {
				tts.model = MockModelId
			},
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, model.English, tts.model.CurrentLanguage, "WithModelLanguage should set the model's currentLanguage field")
			},
		},
		{
			name:   "WithVocoder",
			option: WithVocoder(MockModelId),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, MockModelId, tts.vocoder, "WithVocoder should set the vocoder field")
			},
		},
		{
			name:   "WithVocoderLanguage",
			option: WithVocoderLanguage(model.English),
			pre: func(tts *TTS) {
				tts.vocoder = MockModelId
			},
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, model.English, tts.vocoder.CurrentLanguage, "WithVocoderLanguage should set the vocoder's currentLanguage field")
			},
		},
		{
			name:   "WithSpeaker",
			option: WithSpeaker("speaker001"),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, "speaker001", tts.speakerIdx, "WithSpeaker should set the speakerIdx field")
			},
		},
		{
			name:   "WithSpeakerSample",
			option: WithSpeakerSample("/path/to/sample.wav"),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, "/path/to/sample.wav", tts.speakerSample, "WithSpeakerSample should set the speakerSample field")
			},
		},
		{
			name:   "WithSpeakerIndex",
			option: WithSpeakerIndex("0"),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, "0", tts.speakerIdx, "WithSpeakerIndex should set the speakerIdx field")
			},
		},
		{
			name:   "WithOutputDir",
			option: WithOutputDir(MockOutputDir),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, MockOutputDir, tts.outputDir, "WithOutputDir should set the outputDir field")
			},
		},
		{
			name:   "WithDevice",
			option: WithDevice(model.DeviceCPU),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, model.DeviceCPU, tts.device, "WithDevice should set the device field")
			},
		},
		{
			name:   "WithMaxRetries",
			option: WithMaxRetries(3),
			check: func(t *testing.T, tts *TTS) {
				assert.Equal(t, 3, tts.maxRetries, "WithMaxRetries should set the maxRetries field")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tts := &TTS{}
			if tt.pre != nil {
				tt.pre(tts)
			}
			err := tt.option.apply(tts)
			require.NoError(t, err)
			tt.check(t, tts)
		})
	}
}
