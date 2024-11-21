---
sidebar_position: 3
---

# Installing A Downloaded Model

In this section you will learn how to install a downloaded model.

## Disclaimer

Please make sure you have a backup of your currently installed models!!!

[Where are ollama models stored?](https://github.com/ollama/ollama/blob/main/docs/faq.md#where-are-models-stored)

## Copy all required file to a directory(preferably the same as the executable)

You first need to copy all downloaded blobs to a folder, keep in mind there are no proper validation of you blobs so
make sure you have only the blobs there!

## Rename the manifest file

You should the manifest file (which may be the tag name of you model or latest) to `manifest`

## Install your model

You can use this command to install model:

```shell
omdd install --model=<your-model-name> --blobsPath=<downlaoded-blobs-relative-path>
```

In this case `<your-model-name>` is the name of your model (this can be anything you like you dont have to put the name
of downloaded model) and `<downlaoded-blobs-relative-path>` is the relative path to your model.

### Examples:

```shell
omdd install --model=MyAwesomeModel --blobsPath=./
```

Or if your models are somewhere else

```shell
omdd install --model=MyAwesomeModel --blobsPath=../../path/to/blobs
```

