---
layout: weekly
title: "Weekly news #10"
date: 2021-03-07
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Brave buys a search engine, promises no tracking, no profiling – and may even offer a paid-for, no-ad version](https://www.theregister.com/2021/03/03/brave_buys_a_search_engine)
* [R Conferences for 2021](https://rviews.rstudio.com/2021/03/03/2021-r-conferences/)
* [Google Fonts](https://fonts.google.com/icons)
* [Episode 1: Talking open source, bioinformatics and maple syrup with Phil Ewels of SciLifeLab](https://www.youtube.com/watch?v=Y619gTTSz3I)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [ArchR is a scalable software package for integrative single-cell chromatin accessibility analysis](https://www.nature.com/articles/s41588-021-00790-6)

_Granja et al., Nat Genet 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#scatac-seq</a>

_The advent of single-cell chromatin accessibility profiling has accelerated the ability to map gene regulatory landscapes but has outpaced the development of scalable software to rapidly extract biological meaning from these data. Here we present a software suite for single-cell analysis of regulatory chromatin in R (ArchR; https://www.archrproject.com/) that enables fast and comprehensive analysis of single-cell chromatin accessibility data. ArchR provides an intuitive, user-focused interface for complex single-cell analyses, including doublet removal, single-cell clustering and cell type identification, unified peak set generation, cellular trajectory identification, DNA element-to-gene linkage, transcription factor footprinting, mRNA expression level prediction from chromatin accessibility and multi-omic integration with single-cell RNA sequencing (scRNA-seq). Enabling the analysis of over 1.2 million single cells within 8 h on a standard Unix laptop, ArchR is a comprehensive software suite for end-to-end analysis of single-cell chromatin accessibility that will accelerate the understanding of gene regulation at the resolution of individual cells._

### [Using prototyping to choose a bioinformatics workflow management system](https://journals.plos.org/ploscompbiol/article?id=10.1371%2Fjournal.pcbi.1008622)

_Jackson et al., Plos Comp Biol 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>

_Workflow management systems represent, manage, and execute multistep computational analyses and offer many benefits to bioinformaticians. They provide a common language for describing analysis workflows, contributing to reproducibility and to building libraries of reusable components. They can support both incremental build and re-entrancy—the ability to selectively re-execute parts of a workflow in the presence of additional inputs or changes in configuration and to resume execution from where a workflow previously stopped. Many workflow management systems enhance portability by supporting the use of containers, high-performance computing (HPC) systems, and clouds. Most importantly, workflow management systems allow bioinformaticians to delegate how their workflows are run to the workflow management system and its developers. This frees the bioinformaticians to focus on what these workflows should do, on their data analyses, and on their science. RiboViz is a package to extract biological insight from ribosome profiling data to help advance understanding of protein synthesis. At the heart of RiboViz is an analysis workflow, implemented in a Python script. To conform to best practices for scientific computing which recommend the use of build tools to automate workflows and to reuse code instead of rewriting it, the authors reimplemented this workflow within a workflow management system. To select a workflow management system, a rapid survey of available systems was undertaken, and candidates were shortlisted: Snakemake, cwltool, Toil, and Nextflow. Each candidate was evaluated by quickly prototyping a subset of the RiboViz workflow, and Nextflow was chosen. The selection process took 10 person-days, a small cost for the assurance that Nextflow satisfied the authors’ requirements. The use of prototyping can offer a low-cost way of making a more informed selection of software to use within projects, rather than relying solely upon reviews and recommendations by others._

### [Normalization of single-cell RNA-seq counts by log(x + 1)* or log(1 + x)*](https://academic.oup.com/bioinformatics/advance-article/doi/10.1093/bioinformatics/btab085/6155989)

_Booeshaghi et al., Bioinformatics 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#normalization</a>

_Single-cell RNA-seq technologies have been successfully employed over the past decade to generate many high resolution cell atlases. These have proved invaluable in recent efforts aimed at understanding the cell type specificity of host genes involved in SARS-CoV-2 infections. While single-cell atlases are based on well-sampled highly-expressed genes, many of the genes of interest for understanding SARS-CoV-2 can be expressed at very low levels. Common assumptions underlying standard single-cell analyses don’t hold when examining low-expressed genes, with the result that standard workflows can produce misleading results._

### [Accurate and efficient detection of gene fusions from RNA sequencing data](https://genome.cshlp.org/content/31/3/448.full)

_Uhrig et al., Genome Research 2021_

<a href="#" class="badge badge-primary">#fusion-gene</a>

_The identification of gene fusions from RNA sequencing data is a routine task in cancer research and precision oncology. However, despite the availability of many computational tools, fusion detection remains challenging. Existing methods suffer from poor prediction accuracy and are computationally demanding. We developed Arriba, a novel fusion detection algorithm with high sensitivity and short runtime. When applied to a large collection of published pancreatic cancer samples (n = 803), Arriba identified a variety of driver fusions, many of which affected druggable proteins, including ALK, BRAF, FGFR2, NRG1, NTRK1, NTRK3, RET, and ROS1. The fusions were significantly associated with KRAS wild-type tumors and involved proteins stimulating the MAPK signaling pathway, suggesting that they substitute for activating mutations in KRAS. In addition, we confirmed the transforming potential of two novel fusions, RRBP1-RAF1 and RASGRP1-ATP1A1, in cellular assays. These results show Arriba's utility in both basic cancer research and clinical translation._

### [Fast searches of large collections of single-cell data using scfind](https://www.nature.com/articles/s41592-021-01076-9)

_Lee et al., Nat Methods 2021_

<a href="#" class="badge badge-primary">#searching</a>

_Single-cell technologies have made it possible to profile millions of cells, but for these resources to be useful they must be easy to query and access. To facilitate interactive and intuitive access to single-cell data we have developed scfind, a single-cell analysis tool that facilitates fast search of biologically or clinically relevant marker genes in cell atlases. Using transcriptome data from six mouse cell atlases, we show how scfind can be used to evaluate marker genes, perform in silico gating, and identify both cell-type-specific and housekeeping genes. Moreover, we have developed a subquery optimization routine to ensure that long and complex queries return meaningful results. To make scfind more user friendly, we use indices of PubMed abstracts and techniques from natural language processing to allow for arbitrary queries. Finally, we show how scfind can be used for multi-omics analyses by combining single-cell ATAC-seq data with transcriptome data._

---

## <i class="far fa-keyboard"></i> Programming

### [jsonschema](https://rseng.github.io/rseng/software/jsonschema)

If you ever written a validation schema for XML, you can do the same with `JSON`.

---

## <i class="fas fa-toolbox"></i> Tools

### [MyHeritage](https://www.myheritage.com/)

Animate the faces in your family photos with amazing technology. Experience your family history like never before!

### [Flameshot](https://flameshot.org/)

Powerful yet simple to use screenshot software.

### [trailofbits/graphtage](https://github.com/trailofbits/graphtage)

A semantic diff utility and library for tree-like files such as JSON, JSON5, XML, HTML, YAML, and CSV.

### [jevakallio/git-notify](https://github.com/jevakallio/git-notify)

🙉 📣 Communicate important updates to your team via git commit messages

### [Zero Data App](https://0data.app/)

Own your data, all of it.

### [google-research-datasets/wit](https://github.com/google-research-datasets/wit)

WIT (Wikipedia-based Image Text) Dataset is a large multimodal multilingual dataset comprising 37M+ image-text sets with 11M+ unique images across 100+ languages.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [My Approach to Getting Dramatically Better as a Programmer](https://malisper.me/my-approach-to-getting-dramatically-better-as-a-programmer/)

### [Atlas: Our journey from a Python monolith to a managed platform](https://dropbox.tech/infrastructure/atlas--our-journey-from-a-python-monolith-to-a-managed-platform)

### [PAIRED: A New Multi-agent Approach for Adversarial Environment Generation](https://ai.googleblog.com/2021/03/paired-new-multi-agent-approach-for.html)

---

## <i class="fas fa-rss"></i> Others

### [Pycro-Manager: open-source software for customized and reproducible microscope control](https://www.nature.com/articles/s41592-021-01087-6)
