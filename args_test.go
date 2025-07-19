package coqui

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
				tts.device = DeviceCPU
				tts.modelPath = "/path/to/model"
				tts.model = ModelIdentifier{
					category:        "tts",
					currentLanguage: English,
					dataset:         DatasetLJSpeech,
					model:           BaseModelTacotron2DDC,
				}
			},
			expected: []string{argDevice, "cpu", argModelPath, "/path/to/model"},
		},
		{
			name: "CUDA, modelName, vocoder, not custom",
			setup: func(tts *TTS) {
				tts.device = DeviceCUDA
				tts.model = ModelIdentifier{
					category:        "tts",
					currentLanguage: Spanish,
					dataset:         DatasetCV,
					model:           BaseModelVITS,
				}
				tts.vocoder = ModelIdentifier{
					model:           BaseModelVITS,
					defaultLanguage: Spanish,
					dataset:         DatasetCV,
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
				tts.device = DeviceAuto
				tts.model = ModelIdentifier{
					category:        "tts",
					currentLanguage: French,
					dataset:         DatasetLJSpeech,
					model:           BaseModelTacotron2DDC,
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
				tts.device = DeviceCPU
				tts.model = ModelIdentifier{
					category:             "tts",
					currentLanguage:      German,
					dataset:              DatasetVCTK,
					model:                BaseModelVITS,
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
				tts.device = DeviceCPU
				tts.model = ModelIdentifier{
					category:             "tts",
					currentLanguage:      Japanese,
					dataset:              DatasetCSS10,
					model:                BaseModelVITS,
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
				tts.model = ModelIdentifier{
					category:        "tts",
					currentLanguage: English,
					dataset:         DatasetLJSpeech,
					model:           BaseModelTacotron2DDC,
					isCustom:        true,
				}
				tts.vocoder = ModelIdentifier{
					model:           BaseModelVITS,
					defaultLanguage: English,
					dataset:         DatasetLJSpeech,
				}
			},
			expected: []string{argDevice, "", argModelName, "tts/en/ljspeech/tacotron2-DDC", argSpeakerIdx, ""},
		},
		{
			name: "Missing device",
			setup: func(tts *TTS) {
				tts.model = ModelIdentifier{
					category:        "tts",
					currentLanguage: English,
					dataset:         DatasetLJSpeech,
					model:           BaseModelTacotron2DDC,
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
