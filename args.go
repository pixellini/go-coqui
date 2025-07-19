package coqui

// TODO: I'll remove some of these if I don't need them later.
// For now, I just want an easy reference to the CLI arguments.
const (
	// [TEXT] Text to generate speech.
	argText = "--text"
	// [MODEL_INFO_BY_IDX] model info using query format: <model_type>/<model_query_idx>
	argModelInfoByIdx = "--model_info_by_idx"
	// [MODEL_INFO_BY_NAME] model info using query format: <model_type>/<language>/<dataset>/<model_name>
	argModelInfoByName = "--model_info_by_name"
	// [MODEL_NAME] Name of one of the pre-trained TTS models in format <language>/<dataset>/<model_name>
	argModelName = "--model_name"
	// [VOCODER_NAME] Name of one of the pre-trained  vocoder models in format <language>/<dataset>/<model_name>
	argVocoderName = "--vocoder_name"
	// [CONFIG_PATH] Path to model config file.
	argConfigPath = "--config_path"
	// [MODEL_PATH] Path to model file.
	argModelPath = "--model_path"
	// [OUT_PATH] Output wav file path.
	argOutPath = "--out_path"
	// Run model on CUDA.
	argUseCuda = "--use_cuda"
	// [DEVICE] Device to run model on.
	argDevice = "--device"
	// [VOCODER_PATH] Path to vocoder model file. If it is not defined model uses GL as vocoder. Please make sure that you installed vocoder library before (WaveRNN).
	argVocoderPath = "--vocoder_path"
	// [VOCODER_CONFIG_PATH] Path to vocoder model config file.
	argVocoderConfigPath = "--vocoder_config_path"
	// [ENCODER_PATH] Path to speaker encoder model file.
	argEncoderPath = "--encoder_path"
	// [ENCODER_CONFIG_PATH] Path to speaker encoder config file.
	argEncoderConfigPath = "--encoder_config_path"
	// stdout the generated TTS wav file for shell pipe.
	argPipeOut = "--pipe_out"
	// [SPEAKERS_FILE_PATH] JSON file for multi-speaker model.
	argSpeakersFilePath = "--speakers_file_path"
	// [LANGUAGE_IDS_FILE_PATH] JSON file for multi-lingual model.
	argLanguageIdsFilePath = "--language_ids_file_path"
	// [SPEAKER_IDX] Target speaker ID for a multi-speaker TTS model.
	argSpeakerIdx = "--speaker_idx"
	// [LANGUAGE_IDX] Target language ID for a multi-lingual TTS model.
	argLanguageIdx = "--language_idx"
	// [SPEAKER_WAV [SPEAKER_WAV ...]]wav file(s) to condition a multi-speaker TTS model with a Speaker Encoder. You can give multiple file paths. The d_vectors is computed as their average.
	argSpeakerWav = "--speaker_wav"
	// [GST_STYLE] Wav path file for GST style reference.
	argGstStyle = "--gst_style"
	// [CAPACITRON_STYLE_WAV] Wav path file for Capacitron prosody reference.
	argCapacitronStyleWav = "--capacitron_style_wav"
	// [CAPACITRON_STYLE_TEXT] Transcription of the reference.
	argCapacitronStyleText = "--capacitron_style_text"
	// List available speaker ids for the defined multi-speaker model.
	argListSpeakerIdxs = "--list_speaker_idxs"
	// List available language ids for the defined multi-lingual model.
	argListLanguageIdxs = "--list_language_idxs"
	// [REFERENCE_WAV] Reference wav file to convert in the voice of the speaker_idx or speaker_wav
	argReferenceWav = "--reference_wav"
	// [REFERENCE_SPEAKER_IDX] speaker ID of the reference_wav speaker (If not provided the embedding will be computed using the Speaker Encoder).
	argReferenceSpeakerIdx = "--reference_speaker_idx"
	// Show a progress bar for the model download.
	argProgressBar = "--progress_bar"
	// Disable the progress bar for the model download.
	argNoProgressBar = "--no-progress_bar"
	// [SOURCE_WAV] Original audio file to convert into the voice of the target_wav
	argSourceWav = "--source_wav"
	// [TARGET_WAV ...] Audio file(s) of the target voice into which to convert the source_wav
	argTargetWav = "--target_wav"
	// [VOICE_DIR] Voice dir for tortoise model
	argVoiceDir = "--voice_dir"
)
