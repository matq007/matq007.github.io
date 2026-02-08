---
title: Trying to run local LLM on linux server
date: 2026-02-08T12:00:00
tags:
  - blog/notes
  - llm
---
I've been testing a local LLM on my trusty MacBook Pro M4 16GB using [ollama](https://ollama.com). However, this configuration is limited to only smaller models with quantization usually around 4 bits. Luckily enough, I have access to a linux server with some GPUs to play with. When it comes to running LLM locally with CUDA, based on my reading the options are either ollama, [vLLM](https://vllm.ai/) or [llama.cpp](https://github.com/ggml-org/llama.cpp). 

I wanted to try something new so I started first with vllm and boy oh boy, installing this was a nightmare. The latest wheels for vLLM ([v0.15.1](https://github.com/vllm-project/vllm/releases/tag/v0.15.1)) are compiled for CUDA 13 (I have only access to CUDA 12 on the HPC). I started first with recommended approach using [uv](https://docs.astral.sh/uv/) and `--torch-backend=auto`. After a while I am hitting the first barrier also described in [issue #7785](https://github.com/vllm-project/vllm/issues/7785).

```log
'xxxx/.cache/uv/builds-v0/.tmpMou2on/bin/ninja' '--version'

        failed with:

         no such file or directory
```

Ok, `ninja` is part of the [requirements](https://github.com/vllm-project/vllm/blob/785cf28ffffeddeecf79018074a222b7c5938f9c/requirements/build.txt#L3), so this shouldn't be an issue. No problem, I have ninja available as a module on the HPC, I'll just load it and restart. To no avail, I run `uv cache clean`, restart the process. I'm getting nowhere, time to build from source. So I git clone, setup the venv and run `uv pip install --editable .`.  This time a new problem occurs ([issue #23062](https://github.com/vllm-project/vllm/issues/23062))

```log
Pytorch version 2.9.1 expected for CUDA build, saw 2.4.0 instead.
```

Alright, wrong Pytorch version let me install it manually first `uv pip install torch==2.9.1` and restart. No change, same issue. Ok, probably it downgrades back, so I read the documentation further and find I have to run `python use_existing_torch.py` when enforcing custom Pytorch version. Re-run again and this time

```log
Pytorch version 2.4.0 expected for CUDA build, saw 2.9.1 instead.
```

What the *** is happening here. I've spend few more hours on this and then decided that unless I upgrade to CUDA 13 this is going nowhere. So I abandon this approach for and switch my focus to llama.cpp.

After this experience I wasn't sure I would be successful with llama.cpp, but after I was determined to make this work. Surprisingly the only issue I encountered was that I was using old version of gcc (v9.x) when compiling from source. With simple upgrade to gcc v11, I've successfully managed to compile the package for CUDA with

```bash
$ module load cuda/12.8 cudnn/9.18.0.77_cuda12 cmake/4.0.3 gcc/11.2.0
$ git clone https://github.com/ggml-org/llama.cpp
$ cd llama.cpp
$ cmake -B build -DGGML_CUDA=ON
$ cmake --build build --config Release -- -j 30
```

I've tested it using the [new WebUI](https://github.com/ggml-org/llama.cpp/discussions/16938) and it worked without any issues

```bash
export CUDA_VISIBLE_DEVICES=2,3
llama-server -m glm-4.7-flash-claude-4.5-opus.f16.gguf --port 8080 --ctx-size 0 --jinja -ub 2048 -b 2048
```

My next plan is to connect it to [opencode](https://opencode.ai) and experiment more with the agentic mode. 

I'm still not quite sure what I did wrong with the vLLM setup, but at this point llama.cpp seems to work great. If you've hit the same issues as I did during installation and found a way how to resolve it, do let me know by submitting [an issue here](https://github.com/matq007/matq007.github.io/issues).