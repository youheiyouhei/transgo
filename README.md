# Transgo CLI

`Transgo` is a command-line interface (CLI) tool written in Go, designed to facilitate text translation using the DeepL API.

## Installation

To install Transgo via Homebrew, run the following command:

```
brew install youheiyouhei/tap/transgo
```

## Usage

### Setting Up API Key

Before using the translation features, you need to set up your DeepL API key:

```
transgo config --set api_key=YOUR_API_KEY
```

Replace `YOUR_API_KEY` with your actual DeepL API key.

### Translate Text

To translate text from a source language to a target language:

```
transgo translate --source=en --target=ja "Hello, world!"
```

### List Supported Languages

To view a list of languages supported for translation:

```
transgo languages
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
