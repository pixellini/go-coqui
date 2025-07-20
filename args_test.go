package coqui

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pixellini/go-coqui/model"
)

var MockModel = model.BaseModel("mock-model")

func baseTTS() *TTS {
	return &TTS{}
}

func TestToArgs(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*TTS)
		expected []string
	}{
		{
			name: "CPU, modelPath, no vocoder, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCPU
				tts.modelPath = "/path/to/model"
				tts.model = model.Identifier{
					Category:        "tts",
					CurrentLanguage: model.English,
					Dataset:         model.DatasetLJSpeech,
					Model:           MockModel,
				}
			},
			expected: []string{argDevice, "cpu", argModelPath, "/path/to/model"},
		},
		{
			name: "CUDA, modelName, vocoder, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCUDA
				tts.model = model.Identifier{
					Category:        "tts",
					CurrentLanguage: model.Spanish,
					Dataset:         model.DatasetCV,
					Model:           MockModel,
				}
				tts.vocoder = model.Identifier{
					Model:           MockModel,
					DefaultLanguage: model.Spanish,
					Dataset:         model.DatasetCV,
				}
			},
			expected: []string{
				argDevice, "cuda",
				argModelName, "tts/es/cv/mock-model",
				argUseCuda, "true",
			},
		},
		{
			name: "Auto device resolves, speaker sample, custom model",
			setup: func(tts *TTS) {
				tts.device = model.DeviceAuto
				tts.model = model.Identifier{
					Category:        "tts",
					CurrentLanguage: model.French,
					Dataset:         model.DatasetLJSpeech,
					Model:           MockModel,
					IsCustom:        true,
				}
				tts.speakerIdx = "spk1"
			},
			expected: []string{
				argDevice, "mps",
				argModelName, "tts/fr/ljspeech/mock-model",
				argSpeakerIdx, "spk1",
			},
		},
		{
			name: "Voice cloning, speaker sample, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCPU
				tts.model = model.Identifier{
					Category:             "tts",
					CurrentLanguage:      model.German,
					Dataset:              model.DatasetVCTK,
					Model:                MockModel,
					SupportsVoiceCloning: true,
				}
				tts.speakerSample = "/tmp/clone.wav"
				tts.speakerIdx = "spk2"
			},
			expected: []string{
				argDevice, "cpu",
				argModelName, "tts/de/vctk/mock-model",
				argSpeakerWav, "/tmp/clone.wav",
				argLanguageIdx, "de",
				argSpeakerIdx, "spk2",
			},
		},
		{
			name: "Voice cloning, no speaker sample, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCPU
				tts.model = model.Identifier{
					Category:             "tts",
					CurrentLanguage:      model.Japanese,
					Dataset:              model.DatasetCSS10,
					Model:                MockModel,
					SupportsVoiceCloning: true,
				}
				tts.speakerIdx = "spk3"
			},
			expected: []string{
				argDevice, "cpu",
				argModelName, "tts/ja/css10/mock-model",
				argLanguageIdx, "ja",
				argSpeakerIdx, "spk3",
			},
		},
		{
			name:     "Minimal TTS struct",
			setup:    func(tts *TTS) {},
			expected: []string{argDevice, "", argModelName, "///"},
		},
		{
			name: "Custom model with vocoder (invalid combination)",
			setup: func(tts *TTS) {
				tts.model = model.Identifier{
					Category:        "tts",
					CurrentLanguage: model.English,
					Dataset:         model.DatasetLJSpeech,
					Model:           MockModel,
					IsCustom:        true,
				}
				tts.vocoder = model.Identifier{
					Model:           MockModel,
					DefaultLanguage: model.English,
					Dataset:         model.DatasetLJSpeech,
				}
			},
			expected: []string{argDevice, "", argModelName, "tts/en/ljspeech/mock-model", argSpeakerIdx, ""},
		},
		{
			name: "Missing device",
			setup: func(tts *TTS) {
				tts.model = model.Identifier{
					Category:        "tts",
					CurrentLanguage: model.English,
					Dataset:         model.DatasetLJSpeech,
					Model:           MockModel,
					IsCustom:        true,
				}
			},
			expected: []string{argDevice, "", argModelName, "tts/en/ljspeech/mock-model", argSpeakerIdx, ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tts := baseTTS()
			tt.setup(tts)
			args := toArgs(*tts)
			if !cmp.Equal(args, tt.expected) {
				t.Errorf("Test %q failed:\nDiff:\n%s", tt.name, cmp.Diff(tt.expected, args))
			}
		})
	}
}
