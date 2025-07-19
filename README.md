# 🐸 go-coqui

A Go library for seamlessly converting text files into high-quality audio using Coqui TTS. Easily integrate text-to-speech capabilities into your Go projects with simple, efficient APIs.

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
Used for generating natural-sounding speech  

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
