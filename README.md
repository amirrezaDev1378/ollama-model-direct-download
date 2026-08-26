# Ollama Model Direct Link Generator and Installer

## Overview

The Ollama Model Direct Link Generator and Installer is a utility designed to streamline the process of obtaining direct
download links for Ollama models and installing them. This tool is intended for developers, researchers, and enthusiasts
interested in Ollama models, providing a straightforward and efficient solution.

---

## [New Tutorial & Documentation](https://amirrezadev1378.github.io/ollama-model-direct-download/)

---

### Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
- [Proxy Configuration](#proxy-configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This command-line interface (CLI) tool generates direct download links for Ollama models and allows automatic installation for locally available model files. Written in Golang, it utilizes the Requests library to fetch the necessary links.

Key functionalities include:

- Generating direct download links for Ollama models quickly.
- Installing locally available Ollama models.

## Note

This program doesn't validate what you have downloaded. It will just copy all the files to the model folder to install
it to Ollama. It's recommended that you verify checksums manually.

## Usage

- Download the binary files based on your OS and processor architecture from
  the [Release page](https://github.com/amirrezaDev1378/ollama-model-direct-download/releases).

### Generate Direct Download Links

- Run the binary file in your terminal using this command:

  `omdd get deepseek-coder-v2:latest`

- Wait for the tool to fetch the download link.
- Download all the fetched files.

### Install Ollama Models

- First, make sure to create a backup of your current models.
  [Where are Ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)
- Store your models and your manifest file in any folder.
  The manifest file must be named `manifest` without any file extension.
- Run the following command:

  `omdd install --model=<model-name> --blobsPath=<downloaded-blobs-path>`

- Replace `<model-name>` with the name of your model and `<downloaded-blobs-path>` with the path to the folder where you
  stored model files and the manifest.
- The tool will now install the model. It may take a while. Don't worry if it seems stuck.
- Once it's finished, you can run the model from Ollama normally with:

  `ollama run <model-name>`

## Proxy Configuration

The program supports HTTP/HTTPS proxies through the standard environment variables. If you have a proxy configured, set
the appropriate environment variables before running `omdd`.

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

If your proxy requires authentication, use:

```text
http://username:password@127.0.0.1:8080
```

For example, on Windows PowerShell:

```powershell
$env:HTTP_PROXY="http://username:password@127.0.0.1:8080"
$env:HTTPS_PROXY="http://username:password@127.0.0.1:8080"

omdd get deepseek-coder-v2:latest
```

### SOCKS5 Proxy

SOCKS5 can be configured using the `SOCKS5_PROXY` environment variable.

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

- Clone the repository to your local machine.
- Install the dependencies using `go mod tidy`.
- Make your changes.
- Run the tests using `go test ./...`.
- Build the binary using `make build`.

## License

MIT License
