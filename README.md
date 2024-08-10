# Ollama Model Direct Link Generator and Installer

## Overview

The Ollama Model Direct Link Generator and Installer is a utility designed to streamline the process of obtaining direct
download links for Ollama models and installing them. This tool is intended for developers, researchers, and enthusiasts
interested in Ollama models, providing a straightforward and efficient solution.

### Table of Contents

- [Introduction](#Introduction)
- [Usage](#Usage)
- [Contributing](#Contributing)
- [To-Do Tasks](#TODO-Tasks)
- [License](#License)
- [Known Issues](#Known-Issues)

## Introduction

This command-line interface (CLI) tool generates direct download links for Ollama models and provides installation
instructions. Written in Golang, it utilizes the Requests library to fetch the necessary links.
You can use this tool to:

Key functionalities include:

- Generating direct download links for Ollama models quickly.
- Installing locally available Ollama models.

## Usage

- Download the binary files based on your OS and processor architecture from Release page.
- ### Generate Direct Download Links
    - Run the binary file in your terminal Using this command `omdd get deepseek-coder-v2:latest`.
    - Wait for the tool to fetch the download link.
- ### Install Ollama Models
    - Store your models and your manifest file (can be named latest or your model parameter length e.g. 16b ) in a new
      folder.
    - Run the following command:
    - `omdd install --model=<your-model-name> --blobsPath=<downlaoded-blobs-relative-path>`.
    - Replace `<your-model-name>` with the name of your model and `<downlaoded-blobs-relative-path>` with the path to
      the downloaded blobs.
    - The tool will install the model and provide you with the installation path.

## Contributing

- Clone the repository to your local machine.
- Install the dependencies using `go mod tidy`.
- Make your changes.
- Run the tests using `go test ./...`.
- Build the binary using `go build -o omdd cmd/omdd/main.go`.

## TODO Tasks

List any pending tasks or features you plan to add. For example:

[ ] Improve error handling.
[ ] Add more unit tests.

## Known Issues

We won't validate what you have downloaded any copy all the files to the model folder.

## License

MIT License


