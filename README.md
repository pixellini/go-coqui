# 🐸 go-coqui

A Go library for seamlessly converting text files into high-quality audio using [**Coqui TTS**](https://github.com/idiap/coqui-ai-TTS). Easily integrate text-to-speech capabilities into your Go projects with simple, efficient APIs.

<div align="center">
  <img src="https://raw.githubusercontent.com/pixellini/pixellini/main/assets/go-coqui.png" alt="Logo">
</div>


## 🚧 Current Status

⚠️ **Work in Progress** – This project is still under active development. Functionality is incomplete, and things may break or change frequently.

## 🚀 Getting Started

Instructions for installing and using this tool will be added once it's ready for public use.

### 📦 Requirements

To run this project, you'll need the following dependencies installed:

#### **[Coqui TTS](https://github.com/idiap/coqui-ai-TTS)**
Used for generating natural-sounding speech.
[Check their docs for more information](https://coqui-tts.readthedocs.io/en/latest/)

Install via pip3:

```bash
pip3 install coqui-tts
```

> **Upgrading from original TTS?** 
> 
> If you previously had the original `TTS` package installed, you may encounter dependency conflicts. To resolve this, first uninstall the conflicting packages:
> ```bash
> pip3 uninstall TTS coqpit -y
> pip3 install coqui-tts
> ```
> 
> **Still having issues?** If you continue to get import errors or dependency conflicts, try a clean reinstallation:
> ```bash
> pip3 uninstall coqui-tts coqpit coqpit-config -y
> pip3 cache purge
> pip3 install --no-cache-dir coqui-tts
> ```

## Voice Cloning
provide explanation of supplying a 1-3 minute audio of the desired voice to use for applicable TTS Models (XTTS-v2, XTTS-v1, YourTTS, Bark)

## Example Usage
### Creating with a predefined TTS Model
If you are using a model that supports voice cloning (XTTS-v2, XTTS-v1, YourTTS, Bark...).
You will need to supply either a speaker wav file, or a speaker index.
```go
tts, err := coqui.New(
  coqui.WithModel(coqui.ModelXTTSv2),
  coqui.WithSpeaker("./speaker.wav"),
  // Other options...
)
```

If you are not using a model with voice cloning, you may need to supply a speaker index 
```go
tts, err := coqui.New(
  coqui.WithModel(coqui.ModelVITSVCTK),
  coqui.WithSpeakerIndex("p298"),
  // Other options...
)
```

### Creating a new model
You can build your own model from the provided presets.
For example, this will build the model `tts_models/en/vctk/vits` which **Coqui TTS** already knows.
```go
myModel := coqui.NewTTSModel(coqui.English, coqui.DatasetVCTK, coqui.BaseModelVITS)

tts, err := coqui.New(
  coqui.WithModel(myModel),
  coqui.WithSpeakerIndex("p298"),
  coqui.WithOutputDir("./dist"),
  // Other options...
)
```

### Using a local model
This must be a valid file that can be read, otherwise Coqui will panic.
Since this model is custom, it's entirely up to you to ensure the options you pass are valid and will work as expected when synthesising.

```go
tts, err := coqui.New(
  coqui.WithModelPath("path/to/model")
  // Other options...
)
```

### Synthesizing the text to speech
Once you have defined and configured the model that you wish to use, simple use the `Synthesize` method with a text and output file name:
```go
_, err = tts.Synthesize("Hello World!", "output.wav")
if err != nil {
  fmt.Println("Error synthesizing speech:", err)
  return
}
```

For batch synthesis, you can just loop through a list of texts:
```go
texts := []string{"Hello!", "How are you?", "Goodbye!"}
for i, text := range texts {
  output := fmt.Sprintf("output_%d.wav", i)
  _, err := tts.Synthesize(text, output)
  if err != nil {
    fmt.Println("Error synthesizing:", err)
  }
}
```

Using text from a file:
```go
_, err = tts.SynthesizeFromFile("path/to/file.txt", "output.wav")
if err != nil {
  fmt.Println("Error synthesizing speech:", err)
  return
}
```

## ⚖️ Terms of Use & Disclaimer

By using this tool, you agree to the following:

- **Consent Required:** You must only use voice samples for which you have explicit permission from the person whose voice is featured. Do not use this tool to generate speech using the voice of any individual without their informed consent.
- **User Responsibility:** You are solely responsible for ensuring that your use of this tool and any voice samples you provide comply with all applicable laws and regulations, including those relating to copyright, privacy, and the right of publicity.
- **No Liability:** The developers and contributors of this project are not responsible for any misuse of this tool or any legal consequences arising from the unauthorised use of voice samples.