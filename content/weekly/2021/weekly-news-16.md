---
layout: weekly
title: "Weekly news #16"
date: 2021-04-17
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [COSMIC to Arrive in June Release of Pop!_OS 21.04](https://blog.system76.com/post/648371526931038208/cosmic-to-arrive-in-june-release-of-popos-2104)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Single-cell CUT&Tag profiles histone modifications and transcription factors in complex tissues](https://www.nature.com/articles/s41587-021-00869-9)

_Bartosovic et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#cut-and-tag</a>
<a href="#" class="badge badge-primary">#single-cell</a>

_In contrast to single-cell approaches for measuring gene expression and DNA accessibility, single-cell methods for analyzing histone modifications are limited by low sensitivity and throughput. Here, we combine the CUT&Tag technology, developed to measure bulk histone modifications, with droplet-based single-cell library preparation to produce high-quality single-cell data on chromatin modifications. We apply single-cell CUT&Tag (scCUT&Tag) to tens of thousands of cells of the mouse central nervous system and probe histone modifications characteristic of active promoters, enhancers and gene bodies (H3K4me3, H3K27ac and H3K36me3) and inactive regions (H3K27me3). These scCUT&Tag profiles were sufficient to determine cell identity and deconvolute regulatory principles such as promoter bivalency, spreading of H3K4me3 and promoter–enhancer connectivity. We also used scCUT&Tag to investigate the single-cell chromatin occupancy of transcription factor OLIG2 and the cohesin complex component RAD21. Our results indicate that analysis of histone modifications and transcription factor occupancy at single-cell resolution provides unique insights into epigenomic landscapes in the central nervous system._

### [Integration of multiomics data with graph convolutional networks to identify new cancer genes and their associated molecular mechanisms](https://www.nature.com/articles/s42256-021-00325-y)

_Schulte-Sasse et al., Nat Mach Intell (2021)_

<a href="#" class="badge badge-primary">#cancer</a>
<a href="#" class="badge badge-primary">#prediction</a>
<a href="#" class="badge badge-primary">#machine-learning</a>

_The increase in available high-throughput molecular data creates computational challenges for the identification of cancer genes. Genetic as well as non-genetic causes contribute to tumorigenesis, and this necessitates the development of predictive models to effectively integrate different data modalities while being interpretable. We introduce EMOGI, an explainable machine learning method based on graph convolutional networks to predict cancer genes by combining multiomics pan-cancer data—such as mutations, copy number changes, DNA methylation and gene expression—together with protein–protein interaction (PPI) networks. EMOGI was on average more accurate than other methods across different PPI networks and datasets. We used layer-wise relevance propagation to stratify genes according to whether their classification was driven by the interactome or any of the omics levels, and to identify important modules in the PPI network. We propose 165 novel cancer genes that do not necessarily harbour recurrent alterations but interact with known cancer genes, and we show that they correspond to essential genes from loss-of-function screens. We believe that our method can open new avenues in precision oncology and be applied to predict biomarkers for other complex diseases._

### [An automated framework for efficiently designing deep convolutional neural networks in genomics](https://www.nature.com/articles/s42256-021-00316-z)

_Zhang et al., Nat Mach Intell (2021)_

<a href="#" class="badge badge-primary">#genomics</a>
<a href="#" class="badge badge-primary">#machine-learning</a>

_Convolutional neural networks (CNNs) have become a standard for analysis of biological sequences. Tuning of network architectures is essential for a CNN’s performance, yet it requires substantial knowledge of machine learning and commitment of time and effort. This process thus imposes a major barrier to broad and effective application of modern deep learning in genomics. Here we present Automated Modelling for Biological Evidence-based Research (AMBER), a fully automated framework to efficiently design and apply CNNs for genomic sequences. AMBER designs optimal models for user-specified biological questions through the state-of-the-art neural architecture search (NAS). We applied AMBER to the task of modelling genomic regulatory features and demonstrated that the predictions of the AMBER-designed model are significantly more accurate than the equivalent baseline non-NAS models and match or even exceed published expert-designed models. Interpretation of AMBER architecture search revealed its design principles of utilizing the full space of computational operations for accurately modelling genomic sequences. Furthermore, we illustrated the use of AMBER to accurately discover functional genomic variants in allele-specific binding and disease heritability enrichment. AMBER provides an efficient automated method for designing accurate deep learning models in genomics._

### [Common pitfalls and recommendations for using machine learning to detect and prognosticate for COVID-19 using chest radiographs and CT scans](https://www.nature.com/articles/s42256-021-00307-0)

_Roberts et al., Nat Mach Intell (2021)_

<a href="#" class="badge badge-primary">#covid-19</a>
<a href="#" class="badge badge-primary">#machine-learning</a>

_Machine learning methods offer great promise for fast and accurate detection and prognostication of coronavirus disease 2019 (COVID-19) from standard-of-care chest radiographs (CXR) and chest computed tomography (CT) images. Many articles have been published in 2020 describing new machine learning-based models for both of these tasks, but it is unclear which are of potential clinical utility. In this systematic review, we consider all published papers and preprints, for the period from 1 January 2020 to 3 October 2020, which describe new machine learning models for the diagnosis or prognosis of COVID-19 from CXR or CT images. All manuscripts uploaded to bioRxiv, medRxiv and arXiv along with all entries in EMBASE and MEDLINE in this timeframe are considered. Our search identified 2,212 studies, of which 415 were included after initial screening and, after quality screening, 62 studies were included in this systematic review. Our review finds that none of the models identified are of potential clinical use due to methodological flaws and/or underlying biases. This is a major weakness, given the urgency with which validated COVID-19 models are needed. To address this, we give many recommendations which, if followed, will solve these issues and lead to higher-quality model development and well-documented manuscripts._

### [Simultaneous trimodal single-cell measurement of transcripts, epitopes, and chromatin accessibility using TEA-seq](https://elifesciences.org/articles/63632)

_Swanson et al., eLife (2021)_

<a href="#" class="badge badge-primary">#sequencing</a>

_Single-cell measurements of cellular characteristics have been instrumental in understanding the heterogeneous pathways that drive differentiation, cellular responses to signals, and human disease. Recent advances have allowed paired capture of protein abundance and transcriptomic state, but a lack of epigenetic information in these assays has left a missing link to gene regulation. Using the heterogeneous mixture of cells in human peripheral blood as a test case, we developed a novel scATAC-seq workflow that increases signal-to-noise and allows paired measurement of cell surface markers and chromatin accessibility: integrated cellular indexing of chromatin landscape and epitopes, called ICICLE-seq. We extended this approach using a droplet-based multiomics platform to develop a trimodal assay that simultaneously measures transcriptomics (scRNA-seq), epitopes, and chromatin accessibility (scATAC-seq) from thousands of single cells, which we term TEA-seq. Together, these multimodal single-cell assays provide a novel toolkit to identify type-specific gene regulation and expression grounded in phenotypically defined cell types._

---

## <i class="fas fa-toolbox"></i> Tools

### [smallhadroncollider/taskell](https://github.com/smallhadroncollider/taskell)

Command-line Kanban board/task manager with support for Trello boards and GitHub projects.

### [✨ Open source, experimental, and tiny tools roundup](https://tinytools.directory/)

This is a list of small, free, or experimental tools that might be useful in building your game/website/interactive project. Although I’ve included `standards`, this list has a focus on artful tools and toys that are as fun to use as they are functional.

The goal of this list is to enable making entirely outside of closed production ecosystems or walled software gardens.

### [The open source - Calendly alternative](https://calendso.com/)

Self-hosted or hosted by us. White-label by design. API-driven and ready to be deployed on your own domain. You are in control of your events and data.

### [typeup](https://skuz.xyz/typeup/)

A markup language that gets out of your way.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Downstream analysis for scRNAseq](https://www.youtube.com/watch?v=grgjWmaEtxE)

This technical seminar was presented by Matthew Salaciak, an MSc student in the Johnson Lab (Lady Davis Institute) on April 8th, 2021

### [6 open source tools and tips to securing a Linux server for beginners](https://opensource.com/article/21/4/securing-linux-servers)

Use open source tools to protect your Linux environment from breaches.

### [yengoteam/awesome-gha-snippets](https://github.com/yengoteam/awesome-gha-snippets)

🤯 A list of useful snippets and tips for GitHub Actions (GHA).

### [Work with GitHub Actions in your terminal with GitHub CLI](https://github.blog/2021-04-15-work-with-github-actions-in-your-terminal-with-github-cli/)

`gh` brings GitHub to the command line by helping developers manage pull requests, issues, gists, and much more. As of 1.9.0, even more of GitHub is available in your terminal: GitHub Actions.

---

## <i class="fas fa-rss"></i> Long read

* [The Brain ‘Rotates’ Memories to Save Them From New Sensations](https://www.quantamagazine.org/the-brain-rotates-memories-to-save-them-from-new-sensations-20210415/)
* [The amoral nonsense of Orchid’s embryo selection](https://liorpachter.wordpress.com/2021/04/12/the-amoral-nonsense-of-orchids-embryo-selection/)
* [A 'Worst Nightmare' Cyberattack: The Untold Story Of The SolarWinds Hack](https://www.npr.org/2021/04/16/985439655/a-worst-nightmare-cyberattack-the-untold-story-of-the-solarwinds-hack)
* [Rust is for Professionals](https://gregoryszorc.com/blog/2021/04/13/rust-is-for-professionals/)
* [Unifying the CUDA Python Ecosystem](https://developer.nvidia.com/blog/unifying-the-cuda-python-ecosystem/)

## <i class="fab fa-youtube"></i> Less is more: Why our brains struggle to subtract

<br />
<iframe width="560" height="315"
  src="https://www.youtube-nocookie.com/embed/1y32OpI2_LM"
  title="YouTube video player" frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen>
</iframe>
