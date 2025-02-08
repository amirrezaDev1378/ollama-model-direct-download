---
sidebar_position: 2
---

# Getting Direct Links

In this section you will learn how to use the project to generate direct download link for the projects.

## Finding a model

You first need to find a model from [here](https://ollama.com/library)

Copy the model name (and tag name if you want to download a specific version)

![Docusaurus logo](/img/model-name-and-tag.png)

In this case the model name is qwen2.5-coder and the tag name is 3b

## Generating Links

Use this command to generate direct links

```shell
./omdd get qwen2.5-coder:3b
```

You may omit the tag if you want to get the latest tag (default & standard version) 

```shell
./omdd get qwen2.5-coder
```

The results should look like this:

```text
Fetching direct download link for model: qwen2.5-coder
Manifest download link: https://registry.ollama.ai/v2/library/qwen2.5-coder/manifests/latest
Download links for layers:
1 - https://registry.ollama.ai/v2/library/qwen2.5-coder/blobs/sha256:60e05f2100071479f596b964f89f510f057ce397ea22f2833a0cfe029bfc2463
2 - https://registry.ollama.ai/v2/library/qwen2.5-coder/blobs/sha256:66b9ea09bd5b7099cbb4fc820f31b575c0366fa439b08245566692c6784e281e
3 - https://registry.ollama.ai/v2/library/qwen2.5-coder/blobs/sha256:e94a8ecb9327ded799604a2e478659bc759230fe316c50d686358f932f52776c
4 - https://registry.ollama.ai/v2/library/qwen2.5-coder/blobs/sha256:832dd9e00a68dd83b3c3fb9f5588dad7dcf337a0db50f7d9483f310cd292e92e
5 - https://registry.ollama.ai/v2/library/qwen2.5-coder/blobs/sha256:d9bb33f2786931fea42f50936a2424818aa2f14500638af2f01861eb2c8fb446
Generated download links for model: qwen2.5-coder
Finished successfully!

```

As you can see, we get the links for blobs and the manifest of model.
