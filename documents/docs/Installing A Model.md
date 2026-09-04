---
sidebar_position: 3
---

# Installing A Downloaded Model

This section explains how to install a model you have downloaded.

## Disclaimer

Make sure you have a backup of your current models.
[Where are ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)

## Prepare your files

First, copy all the downloaded blobs into one folder. We recommend putting them in the same folder as the executable. The tool does not validate the files, so check that only the correct blobs are in this folder.

## Rename the manifest file

You need to rename your manifest file. It must be named `manifest` exactly, without any file extension.

## Install your model

Run this command to install the model:

```shell
omdd install --model=<model-name> --blobsPath=<downloaded-blobs-path>
```

Replace `<model-name>` with whatever name you want to give your model. Replace `<downloaded-blobs-path>` with the relative path to the folder containing your model files.

### Examples:

If your files are in the same folder:

```shell
omdd install --model=MyAwesomeModel --blobsPath=./
```

If your files are stored in a different folder:

```shell
omdd install --model=MyAwesomeModel --blobsPath=../../path/to/blobs
```

## Run your model

That is it! You have installed your model. You can now run it in Ollama with this command:

```shell
ollama run <model-name>
```
