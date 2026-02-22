---
title: microsoft/BitNet
date: 2026-01-24T12:00:00
tags:
  - blog/weblink
  - llm
---
A new paradigm shift from Microsoft called [BitNet](https://github.com/microsoft/BitNet), where the objective is to reduce speed and footprint of LLM models, by storing the LLM weights as 1bit. This is done during training, compared to quantization which fixes the weights post training. Authors report promising benchmarking with *BitNet b1.58 2B* using only **400MB** RAM and 29ms latency. It will be interesting to see where we will see these types of models deployed in future. Targets could be phones, IoT or privacy focused closed offline  systems. For more in-depth details, I highly recommend watching [this](https://www.youtube.com/watch?v=WBm0nyDkVYM) great explanation by [Julia Turc](https://www.patreon.com/JuliaTurc). 

Notes on some applied tricks:
- store multiple values in 8bit int format (requires custom loading these values)
- replaced multiplication with addition/subtraction (faster computation for CPUs)
- normalization layer in order to control for extreme values
- storing weights in ternary format: -1, 0, 1

> Ma, Shuming, et al. “The Era of 1-Bit LLMs: All Large Language Models Are in 1.58 Bits.” arXiv:2402.17764, arXiv, 27 Feb. 2024. _arXiv.org_, https://doi.org/10.48550/arXiv.2402.17764.