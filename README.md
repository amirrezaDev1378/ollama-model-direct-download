# Ollama Model Direct Link Generator and Installer

## Overview

The Ollama Model Direct Link Generator and Installer is a tool that helps you get direct download links for Ollama models and install them. It is useful for anyone who wants a fast and simple way to download and manage Ollama models.

## [New Tutorial & Documentation](https://amirrezadev1378.github.io/ollama-model-direct-download/)

### Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Proxy Configuration](#proxy-configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This CLI tool generates direct download links for Ollama models and installs the local model files for you. It is written in Golang and uses the Requests library to get the links.

Features:

- Gets direct download links for Ollama models fast.
- Installs local Ollama model files.

## Note

This program does not validate your downloads. It copies the files to the model folder to install them to Ollama. You should verify the file checksums yourself.

## Usage

Download the binary files for your OS and processor architecture from the [Release page](https://github.com/amirrezaDev1378/ollama-model-direct-download/releases).

### Generate Direct Download Links

Run the binary file in your terminal:

`omdd get deepseek-coder-v2:latest`

Wait for the tool to get the download links. Then, download all the files.

### Install Ollama Models

First, create a backup of your current models.
[Where are Ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)

Store your downloaded models and the manifest file in a folder.
The manifest file must be named `manifest` exactly, with no file extension.

Run this command:

`omdd install --model=<model-name> --blobsPath=<downloaded-blobs-path>`

Replace `<model-name>` with the name you want for your model and `<downloaded-blobs-path>` with the path to the folder where you stored the model files and the manifest.

The tool will install the model. It can take a while. Don't worry if it looks stuck.
When it is done, you can run the model from Ollama:

`ollama run <model-name>`

## Proxy Configuration

The program supports HTTP and HTTPS proxies using standard environment variables. If you use a proxy, set these variables before running `omdd`.

### HTTP/HTTPS Proxy

#### Windows PowerShell

```powershell
$env:HTTP_PROXY="http://127.0.0.1:8080"
$env:HTTPS_PROXY="http://127.0.0.1:8080"

omdd get deepseek-coder-v2:latest
```

#### Windows CMD

```cmd
set HTTP_PROXY=http://127.0.0.1:8080
set HTTPS_PROXY=http://127.0.0.1:8080

omdd get deepseek-coder-v2:latest
```

#### Linux / macOS

```bash
export HTTP_PROXY="http://127.0.0.1:8080"
export HTTPS_PROXY="http://127.0.0.1:8080"

omdd get deepseek-coder-v2:latest
```

If your proxy needs authentication, use this format:

```text
http://username:password@127.0.0.1:8080
```

For example, in Windows PowerShell:

```powershell
$env:HTTP_PROXY="http://username:password@127.0.0.1:8080"
$env:HTTPS_PROXY="http://username:password@127.0.0.1:8080"

omdd get deepseek-coder-v2:latest
```

### SOCKS5 Proxy

You can configure SOCKS5 using the `SOCKS5_PROXY` environment variable.

#### Windows PowerShell

```powershell
$env:SOCKS5_PROXY="socks5://127.0.0.1:1080"

omdd get deepseek-coder-v2:latest
```

With authentication:

```powershell
$env:SOCKS5_PROXY="socks5://username:password@127.0.0.1:1080"

omdd get deepseek-coder-v2:latest
```

#### Windows CMD

```cmd
set SOCKS5_PROXY=socks5://127.0.0.1:1080

omdd get deepseek-coder-v2:latest
```

#### Linux / macOS

```bash
export SOCKS5_PROXY="socks5://127.0.0.1:1080"

omdd get deepseek-coder-v2:latest
```

## Contributing

- Clone the repository to your machine.
- Install the dependencies with `go mod tidy`.
- Make your changes.
- Run the tests with `go test ./...`.
- Build the binary with `make build`.

## License

MIT License
