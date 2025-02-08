---
sidebar_position: 3
---

# Installing A Downloaded Model

In this section you will learn how to install a downloaded model.

## Disclaimer

Please make sure you have a backup of your currently installed models.
[Where are ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)

## Copy all required file to a directory (preferably the same as the executable)

You first need to copy all downloaded blobs to a folder. Keep in mind there are no proper validation of files, so
make sure you have only the blobs there!

## Rename the manifest file

The manifest file must be named `manifest` without any file extension.

## Install your model

You can use this command to install the model:

```shell
omdd install --model=<model-name> --blobsPath=<downloaded-blobs-path>
```

In this case `<model-name>` is the name of your model (you can give it any name you wish) and `<downloaded-blobs-path>` is the relative path to your model.

### Examples:

```shell
omdd install --model=MyAwesomeModel --blobsPath=./
```

Or if your models are stored somewhere else

```shell
omdd install --model=MyAwesomeModel --blobsPath=../../path/to/blobs
```

## Run your model

Success! You've successfully installed your model. You can now run it normally using Ollama with command:
```shell
ollama run <model-name>
```

Happy AI inference!