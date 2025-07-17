package coqui

import "slices"

// Dataset represents a predefined dataset identifier used in TTS synthesis.
type Dataset string

const (
	// Universal/Multilingual datasets.

	// These datasets are designed to support multiple languages and are often used for training multilingual models.
	DatasetLibriTTS Dataset = "libri-tts"
	// DatasetMultiDataset is a placeholder for datasets that combine multiple datasets or are used in a multi-dataset context.
	DatasetMultiDataset Dataset = "multi-dataset"

	// English datasets.

	// These datasets are primarily focused on English language synthesis and are commonly used for training English TTS models.
	DatasetLJSpeech Dataset = "ljspeech"
	// DatasetVCTK is a dataset containing recordings of multiple speakers, often used for voice cloning and speaker adaptation.
	DatasetVCTK Dataset = "vctk"
	// DatasetEK1 is a dataset that includes recordings from the EK1 corpus, often used for English TTS synthesis.
	DatasetEK1 Dataset = "ek1"
	// DatasetSam is a dataset that includes recordings from the Sam corpus, often used for English TTS synthesis.
	DatasetSam Dataset = "sam"
	// DatasetBlizzard2013 is a dataset from the Blizzard Challenge 2013, which includes recordings for TTS synthesis.
	DatasetBlizzard2013 Dataset = "blizzard2013"
	// DatasetJenny is a dataset that includes recordings from the Jenny corpus, often used for English TTS synthesis.
	DatasetJenny Dataset = "jenny"

	// Language-specific datasets.

	// These datasets are focused on specific languages or regions, often used for training TTS models in those languages.
	DatasetMai Dataset = "mai"
	// DatasetCSS10 is a dataset that includes recordings from 10 different languages, often used for multilingual TTS synthesis.
	DatasetCSS10 Dataset = "css10"
	// DatasetCV is a dataset from Mozilla's Common Voice project, which includes recordings in multiple languages.
	DatasetCV Dataset = "cv"
	// DatasetCommonVoice is an alternate identifier for the Common Voice dataset, used in some contexts.
	DatasetCommonVoice Dataset = "common-voice"
	// DatasetThorsten is a dataset that includes recordings from the Thorsten corpus, often used for German TTS synthesis.
	DatasetThorsten Dataset = "thorsten"
	// DatasetBaker is a dataset that includes recordings from the Baker corpus, often used for Chinese TTS synthesis.
	DatasetBaker Dataset = "baker"
	// DatasetKokoro is a dataset that includes recordings from the Kokoro corpus, often used for Japanese TTS synthesis.
	DatasetKokoro Dataset = "kokoro"
	// DatasetOpenBible is a dataset that includes recordings in various African languages, often used for TTS synthesis in those languages.
	DatasetOpenBible Dataset = "openbible"
	// DatasetCustom is a placeholder for custom datasets that users may create or use for specific TTS tasks.
	DatasetCustom Dataset = "custom"

	// Specific dataset variants.

	// DatasetMaiFemale is a variant of the Mai dataset.
	DatasetMaiFemale Dataset = "mai_female"
	// DatasetMaiMale is a variant of the Mai dataset.
	DatasetMaiMale Dataset = "mai_male"
)

// AllDatasets contains a list of all predefined dataset identifiers.
var allDatasets = []Dataset{
	// Universal/Multilingual datasets.
	DatasetLibriTTS,
	DatasetMultiDataset,

	// English datasets.
	DatasetLJSpeech,
	DatasetVCTK,
	DatasetEK1,
	DatasetSam,
	DatasetBlizzard2013,
	DatasetJenny,

	// Language-specific datasets.
	DatasetMai,
	DatasetCSS10,
	DatasetCV,
	DatasetThorsten,
	DatasetBaker,
	DatasetKokoro,
	DatasetOpenBible,
	DatasetCustom,

	// Specific dataset variants.
	DatasetMaiFemale,
	DatasetMaiMale,
}

// String returns the string representation of the Dataset
func (d Dataset) String() string {
	return string(d)
}

// IsValid checks if the dataset is one of the predefined datasets.
func (d Dataset) IsValid() bool {
	return slices.Contains(allDatasets, d)
}

// GetAllDatasets returns a list of all predefined dataset identifiers.
func GetAllDatasets() []Dataset {
	return slices.Clone(allDatasets)
}
