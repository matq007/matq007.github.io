---
layout: weekly
title: "Weekly news #17"
date: 2021-04-24
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Ubuntu 21.04 is out](https://ubuntu.com//blog/ubuntu-21-04-is-here)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Machine-learning-based dynamic-importance sampling for adaptive multiscale simulations](https://www.nature.com/articles/s42256-021-00327-w)

_Bhatia et al., Nat Mach Intell (2021)_

<a href="#" class="badge badge-primary">#machine-learning</a>
<a href="#" class="badge badge-primary">#sampling</a>

_Multiscale simulations are a well-accepted way to bridge the length and time scales required for scientific studies with the solution accuracy achievable through available computational resources. Traditional approaches either solve a coarse model with selective refinement or coerce a detailed model into faster sampling, both of which have limitations. Here, we present a paradigm of adaptive, multiscale simulations that couple different scales using a dynamic-importance sampling approach. Our method uses machine learning to dynamically and exhaustively sample the phase space explored by a macro model using microscale simulations and enables an automatic feedback from the micro to the macro scale, leading to a self-healing multiscale simulation. As a result, our approach delivers macro length and time scales, but with the effective precision of the micro scale. Our approach is arbitrarily scalable as well as transferable to many different types of simulations. Our method made possible a multiscale scientific campaign of unprecedented scale to understand the interactions of RAS proteins with a plasma membrane in the context of cancer research running over several days on Sierra, which is currently the second-most-powerful supercomputer in the world._

### [Iterative single-cell multi-omic integration using online learning](https://www.nature.com/articles/s41587-021-00867-x)

_Gao et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#integration</a>
<a href="#" class="badge badge-primary">#multimodal</a>

_Integrating large single-cell gene expression, chromatin accessibility and DNA methylation datasets requires general and scalable computational approaches. Here we describe online integrative non-negative matrix factorization (iNMF), an algorithm for integrating large, diverse and continually arriving single-cell datasets. Our approach scales to arbitrarily large numbers of cells using fixed memory, iteratively incorporates new datasets as they are generated and allows many users to simultaneously analyze a single copy of a large dataset by streaming it over the internet. Iterative data addition can also be used to map new data to a reference dataset. Comparisons with previous methods indicate that the improvements in efficiency do not sacrifice dataset alignment and cluster preservation performance. We demonstrate the effectiveness of online iNMF by integrating more than 1 million cells on a standard laptop, integrating large single-cell RNA sequencing and spatial transcriptomic datasets, and iteratively constructing a single-cell multi-omic atlas of the mouse motor cortex._

### [Integration of multiomics data with graph convolutional networks to identify new cancer genes and their associated molecular mechanisms](https://www.nature.com/articles/s42256-021-00325-y)

_Schulte-Sasse et al., Nat Mach Intell (2021)_

<a href="#" class="badge badge-primary">#ai</a>
<a href="#" class="badge badge-primary">#multimodal</a>
<a href="#" class="badge badge-primary">#cancer</a>

_The increase in available high-throughput molecular data creates computational challenges for the identification of cancer genes. Genetic as well as non-genetic causes contribute to tumorigenesis, and this necessitates the development of predictive models to effectively integrate different data modalities while being interpretable. We introduce EMOGI, an explainable machine learning method based on graph convolutional networks to predict cancer genes by combining multiomics pan-cancer data—such as mutations, copy number changes, DNA methylation and gene expression—together with protein–protein interaction (PPI) networks. EMOGI was on average more accurate than other methods across different PPI networks and datasets. We used layer-wise relevance propagation to stratify genes according to whether their classification was driven by the interactome or any of the omics levels, and to identify important modules in the PPI network. We propose 165 novel cancer genes that do not necessarily harbour recurrent alterations but interact with known cancer genes, and we show that they correspond to essential genes from loss-of-function screens. We believe that our method can open new avenues in precision oncology and be applied to predict biomarkers for other complex diseases._

### [Compressed sensing for highly efficient imaging transcriptomics](https://www.nature.com/articles/s41587-021-00883-x)

_Cleary et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#imaging</a>

_Recent methods for spatial imaging of tissue samples can identify up to ~100 individual proteins1,2,3 or RNAs4,5,6,7,8,9,10 at single-cell resolution. However, the number of proteins or genes that can be studied in these approaches is limited by long imaging times. Here we introduce Composite In Situ Imaging (CISI), a method that leverages structure in gene expression across both cells and tissues to limit the number of imaging cycles needed to obtain spatially resolved gene expression maps. CISI defines gene modules that can be detected using composite measurements from imaging probes for subsets of genes. The data are then decompressed to recover expression values for individual genes. CISI further reduces imaging time by not relying on spot-level resolution, enabling lower magnification acquisition, and is overall about 500-fold more efficient than current methods. Applying CISI to 12 mouse brain sections, we accurately recovered the spatial abundance of 37 individual genes from 11 composite measurements covering 180 mm2 and 476,276 cells._

### [Modeling transcriptional regulation using gene regulatory networks based on multi-omics data sources](https://bmcbioinformatics.biomedcentral.com/articles/10.1186/s12859-021-04126-3)

_Patel et al., BMC Bioinformatics (2021)_

<a href="#" class="badge badge-primary">#grn</a>

_Transcriptional regulation is complex, requiring multiple cis (local) and trans acting mechanisms working in concert to drive gene expression, with disruption of these processes linked to multiple diseases. Previous computational attempts to understand the influence of regulatory mechanisms on gene expression have used prediction models containing input features derived from cis regulatory factors. However, local chromatin looping and trans-acting mechanisms are known to also influence transcriptional regulation, and their inclusion may improve model accuracy and interpretation. In this study, we create a general model of transcription factor influence on gene expression by incorporating both cis and trans gene regulatory features. We describe a computational framework to model gene expression for GM12878 and K562 cell lines. This framework weights the impact of transcription factor-based regulatory data using multi-omics gene regulatory networks to account for both cis and trans acting mechanisms, and measures of the local chromatin context. These prediction models perform significantly better compared to models containing cis-regulatory features alone. Models that additionally integrate long distance chromatin interactions (or chromatin looping) between distal transcription factor binding regions and gene promoters also show improved accuracy. As a demonstration of their utility, effect estimates from these models were used to weight cis-regulatory rare variants for sequence kernel association test analyses of gene expression. Our models generate refined effect estimates for the influence of individual transcription factors on gene expression, allowing characterization of their roles across the genome. This work also provides a framework for integrating multiple data types into a single model of transcriptional regulation._

---

## <i class="far fa-keyboard"></i> Programming

### [Clever Hack Finds Mystery CPU Instructions](https://www.eejournal.com/article/clever-hack-finds-mystery-cpu-instructions/)

Enterprising Programmer Uses x86 Microcode to Reveal Itself

### [A Better Model for Stacked (GitHub) Pull Requests](https://timothyandrew.dev/blog/git-stack/)

How to make a stacked PR.

### [Basic Music Theory in ~200 Lines of Python](https://www.mvanga.com/blog/basic-music-theory-in-200-lines-of-python)

This article explains the very basics of Western music theory in around 200 lines of Python.

### [Evolving Reinforcement Learning Algorithms](https://ai.googleblog.com/2021/04/evolving-reinforcement-learning.html)

The authors show that it’s possible to learn new, analytically interpretable and generalizable RL algorithms by using a graph representation and applying optimization techniques from the AutoML community.

---

## <i class="fas fa-toolbox"></i> Tools

### [Pytlicek/sheet2dict](https://github.com/Pytlicek/sheet2dict)

Simple XLSX and CSV to dictionary converter.

### [Zellij: a Rusty terminal workspace releases a beta](https://zellij.dev/news/beta/#)

A terminal workspace with batteries included.

### [Litmaps](https://www.litmaps.co/)

Litmaps combines interactive citation maps, modern search tools, and regular email updates, to create the best research discovery experience ever.

### [Starboard](https://starboard.gg/)

The shareable in-browser notebook.

### [gruntwork-io/git-xargs](https://github.com/gruntwork-io/git-xargs)

git-xargs is a command-line tool (CLI) for making updates across multiple Github repositories with a single command.

### [MarkShow](https://mark.show/)

Markdown to Slideshow (Powered by Reveal.js).

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [An Alternative to the Correlation Coefficient That Works For Numeric and Categorical Variables](https://rviews.rstudio.com/2021/04/15/an-alternative-to-the-correlation-coefficient-that-works-for-numeric-and-categorical-variables/)

by Dr. Rama Ramakrishnan.

### [March 2021: "Top 40" New CRAN Packages](https://rviews.rstudio.com/2021/04/22/march-2021-top-40-new-cran-packages/)

by Joseph Rickert.

---

## <i class="fas fa-rss"></i> Others

### [Pokemon or Big Data](https://pixelastic.github.io/pokemonorbigdata/)

Guess from the name if it's a Pokemon or Big Data platform/tool.

---

## <i class="fas fa-rss"></i> Long read

* [Feynman: I am burned out and I'll never accomplish anything (1985)](https://www.asc.ohio-state.edu/kilcup.1/262/feynman.html?repostindays=413)
* [How often do people actually copy and paste from Stack Overflow? Now we know.](https://stackoverflow.blog/2021/04/19/how-often-do-people-actually-copy-and-paste-from-stack-overflow-now-we-know/)
* [Signal's CEO Just Hacked the Cops' Favorite Phone Cracking Tool and Became a Legend](https://gizmodo.com/signals-ceo-just-hacked-the-cops-favorite-phone-crackin-1846733412)
* [Pythagorean Mean Rank Metrics](https://cthoyt.com/2021/04/19/pythagorean-mean-ranks.html)
* [Getting Into FAANG as an Aspiring Developer](https://cvcompiler.com/blog/getting-into-faang-as-an-aspiring-developer)