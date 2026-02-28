---
title: Pachter et al., 2026 bioRxiv
date: 2026-02-23T08:00
tags:
  - blog/weblink
  - science
---
The work of _Pachter et al._, focuses on porting well established [edgeR](https://bioconductor.org/packages/release/bioc/html/edgeR.html) package to Python (edgePython) using Claude Opus 4.5, 4.6 and Codex. As described in the paper, the edgeR is a non trivial package supporting various functions used for differential expression analysis in genomics research. Due to Python's popularity in data science, having edgeR natively available in Python ecosystem, would be highly beneficial for the scientific community. Although, there exists alternative solutions of combining/converting R and Python objects interchangeably[^1], they can be cumbersome.

We have already seen an interest from the single-cell community where [DESeq2](https://bioconductor.org/packages/release/bioc/html/DESeq2.html) was rewritten from scratch to [PyDESeq2](https://github.com/scverse/PyDESeq2), although it does not produce exact results as the original. That being said, there were few things I enjoyed when reading this article.

1. Transparency of usage of LLM models.
2. Using existing tests in edgeR to verify correctness. I do not think asking LLM model to write tests against its own implementation would be a good idea in practice.

I am curious to see if more tools in future will be ported to other more optimized programming languages such as Rust with bindings for interpreters (i.e. [pyo3](https://github.com/PyO3/pyo3)). It seems like these types of tasks are well suited for the LLMs where the objective is not to invent ([gcc compiler](https://www.anthropic.com/engineering/building-c-compiler)) but to rather "translate" from one language to another.

> Lior Pachter, **Differential analysis of genomics count data with edge**. bioRxiv 2026.02.16.706223; doi: https://doi.org/10.64898/2026.02.16.706223

[^1]: https://github.com/theislab/zellkonverter