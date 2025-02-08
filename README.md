# Ollama Model Direct Link Generator and Installer

## Overview

The Ollama Model Direct Link Generator and Installer is a utility designed to streamline the process of obtaining direct
download links for Ollama models and installing them. This tool is intended for developers, researchers, and enthusiasts
interested in Ollama models, providing a straightforward and efficient solution.
_________

## [New Tutorial & Documentation](https://amirrezadev1378.github.io/ollama-model-direct-download/)

___________
### Table of Contents

- [Introduction](#Introduction)
- [Usage](#Usage)
- [Contributing](#Contributing)
- [License](#License)

## Introduction

This command-line interface (CLI) tool generates direct download links for Ollama models and allows automatic installation for locally available model files. Written in Golang, it utilizes the Requests library to fetch the necessary links.

Key functionalities include:

- Generating direct download links for Ollama models quickly.
- Installing locally available Ollama models.

## Note

This program doesn't validate what you have downloaded. I will just copy all the files to the model folder to install it to Ollama. Its recommended that you verify checksums manually.

## Usage

- Download the binary files based on your OS and processor architecture
  from [Release page](https://github.com/amirrezaDev1378/ollama-model-direct-download/releases).
- ### Generate Direct Download Links
    - Run the binary file in your terminal Using this command `omdd get deepseek-coder-v2:latest`.
    - Wait for the tool to fetch the download link.
    - Download all the fetched files.
- ### Install Ollama Models
    - First make sure to create a backup of your current models.
      [Where are ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)
    - Store your models and your manifest file in any folder.
      The manifest file must be named 'manifest' without any file extension.
    - Run the following command:
    - `omdd install --model=<model-name> --blobsPath=<downloaded-blobs-path>`.
    - Replace `<model-name>` with the name of your model and `<downloaded-blobs-path>` with the path to the folder where you stored model files & manifest.
    - The tool will now install the model. It may take a while. Don't worry if it seems stuck.
    - Once its finished, you can run the model from Ollama normally with command `ollama run <model-name>`

## Contributing

- Clone the repository to your local machine.
- Install the dependencies using `go mod tidy`.
- Make your changes.
- Run the tests using `go test ./...`.
- Build the binary using `make build`.

## License

MIT License