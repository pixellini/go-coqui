package coqui

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pixellini/go-coqui/model"
)

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
				tts.model = model.ModelIdentifier{
					category:        "tts",
					currentLanguage: model.English,
					dataset:         model.DatasetLJSpeech,
					model:           model.BaseModelTacotron2DDC,
				}
			},
			expected: []string{argDevice, "cpu", argModelPath, "/path/to/model"},
		},
		{
			name: "CUDA, modelName, vocoder, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCUDA
				tts.model = model.ModelIdentifier{
					category:        "tts",
					currentLanguage: model.Spanish,
					dataset:         model.DatasetCV,
					model:           model.BaseModelVITS,
				}
				tts.vocoder = model.ModelIdentifier{
					model:           model.BaseModelVITS,
					defaultLanguage: model.Spanish,
					dataset:         model.DatasetCV,
				}
			},
			expected: []string{
				argDevice, "cuda",
				argModelName, "tts/es/cv/vits",
				argUseCuda, "true",
			},
		},
		{
			name: "Auto device resolves, speaker sample, custom model",
			setup: func(tts *TTS) {
				tts.device = model.DeviceAuto
				tts.model = model.ModelIdentifier{
					category:        "tts",
					currentLanguage: model.French,
					dataset:         model.DatasetLJSpeech,
					model:           model.BaseModelTacotron2DDC,
					isCustom:        true,
				}
				tts.speakerIdx = "spk1"
			},
			expected: []string{
				argDevice, "mps",
				argModelName, "tts/fr/ljspeech/tacotron2-DDC",
				argSpeakerIdx, "spk1",
			},
		},
		{
			name: "Voice cloning, speaker sample, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCPU
				tts.model = model.ModelIdentifier{
					category:             "tts",
					currentLanguage:      model.German,
					dataset:              model.DatasetVCTK,
					model:                model.BaseModelVITS,
					supportsVoiceCloning: true,
				}
				tts.speakerSample = "/tmp/clone.wav"
				tts.speakerIdx = "spk2"
			},
			expected: []string{
				argDevice, "cpu",
				argModelName, "tts/de/vctk/vits",
				argSpeakerWav, "/tmp/clone.wav",
				argLanguageIdx, "de",
				argSpeakerIdx, "spk2",
			},
		},
		{
			name: "Voice cloning, no speaker sample, not custom",
			setup: func(tts *TTS) {
				tts.device = model.DeviceCPU
				tts.model = model.ModelIdentifier{
					category:             "tts",
					currentLanguage:      model.Japanese,
					dataset:              model.DatasetCSS10,
					model:                model.BaseModelVITS,
					supportsVoiceCloning: true,
				}
				tts.speakerIdx = "spk3"
			},
			expected: []string{
				argDevice, "cpu",
				argModelName, "tts/ja/css10/vits",
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
				tts.model = model.ModelIdentifier{
					category:        "tts",
					currentLanguage: model.English,
					dataset:         model.DatasetLJSpeech,
					model:           model.BaseModelTacotron2DDC,
					isCustom:        true,
				}
				tts.vocoder = model.ModelIdentifier{
					model:           model.BaseModelVITS,
					defaultLanguage: model.English,
					dataset:         model.DatasetLJSpeech,
				}
			},
			expected: []string{argDevice, "", argModelName, "tts/en/ljspeech/tacotron2-DDC", argSpeakerIdx, ""},
		},
		{
			name: "Missing device",
			setup: func(tts *TTS) {
				tts.model = model.ModelIdentifier{
					category:        "tts",
					currentLanguage: model.English,
					dataset:         model.DatasetLJSpeech,
					model:           model.BaseModelTacotron2DDC,
					isCustom:        true,
				}
			},
			expected: []string{argDevice, "", argModelName, "tts/en/ljspeech/tacotron2-DDC", argSpeakerIdx, ""},
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
