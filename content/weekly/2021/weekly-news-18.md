---
layout: weekly
draft: true
title: "Weekly news #18"
date: 2021-05-01
tags: [weekly]
---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [DRscDB: A single-cell RNA-seq resource for data mining and data comparison across species](https://www.sciencedirect.com/science/article/pii/S2001037021001331)

_Hu et al., Comp and Struct Biotech Journal (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#database</a>

_With the advent of single-cell RNA sequencing (scRNA-seq) technologies, there has been a spike in studies involving scRNA-seq of several tissues across diverse species including Drosophila. Although a few databases exist for users to query genes of interest within the scRNA-seq studies, search tools that enable users to find orthologous genes and their cell type-specific expression patterns across species are limited. Here, we built a new search database, DRscDB (https://www.flyrnai.org/tools/single_cell/web/), to address this need. DRscDB serves as a comprehensive repository for published scRNA-seq datasets for Drosophila and relevant datasets from human and other model organisms. DRscDB is based on manual curation of Drosophila scRNA-seq studies of various tissue types and their corresponding analogous tissues in vertebrates including zebrafish, mouse, and human. Of note, our search database provides most of the literature-derived marker genes, thus preserving the original analysis of the published scRNA-seq datasets. Finally, DRscDB serves as a web-based user interface that allows users to mine gene expression data from scRNA-seq studies and perform cell cluster enrichment analyses pertaining to various scRNA-seq studies, both within and across species._

### [A single-embryo, single-cell time-resolved model for mouse gastrulation](https://www.cell.com/cell/fulltext/S0092-8674(21)00439-6)

_Mittnenzweig et al., Cell Press (2021)_

<a href="#" class="badge badge-primary">#mouse</a>
<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#gastrulation</a>

_Mouse embryonic development is a canonical model system for studying mammalian cell fate acquisition. Recently, single-cell atlases comprehensively charted embryonic transcriptional landscapes, yet inference of the coordinated dynamics of cells over such atlases remains challenging. Here, we introduce a temporal model for mouse gastrulation, consisting of data from 153 individually sampled embryos spanning 36 h of molecular diversification. Using algorithms and precise timing, we infer differentiation flows and lineage specification dynamics over the embryonic transcriptional manifold. Rapid transcriptional bifurcations characterize the commitment of early specialized node and blood cells. However, for most lineages, we observe combinatorial multi-furcation dynamics rather than hierarchical transcriptional transitions. In the mesoderm, dozens of transcription factors combinatorially regulate multifurcations, as we exemplify using time-matched chimeric embryos of Foxc1/Foxc2 mutants. Our study rejects the notion of differentiation being governed by a series of binary choices, providing an alternative quantitative model for cell fate acquisition._

### [Bayesian inference of gene expression states from single-cell RNA-seq data](https://www.nature.com/articles/s41587-021-00875-x)

_Breda et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#gene-expression</a>

_Despite substantial progress in single-cell RNA-seq (scRNA-seq) data analysis methods, there is still little agreement on how to best normalize such data. Starting from the basic requirements that inferred expression states should correct for both biological and measurement sampling noise and that changes in expression should be measured in terms of fold changes, we here derive a Bayesian normalization procedure called Sanity (SAmpling-Noise-corrected Inference of Transcription activitY) from first principles. Sanity estimates expression values and associated error bars directly from raw unique molecular identifier (UMI) counts without any tunable parameters. Using simulated and real scRNA-seq datasets, we show that Sanity outperforms other normalization methods on downstream tasks, such as finding nearest-neighbor cells and clustering cells into subtypes. Moreover, we show that by systematically overestimating the expression variability of genes with low expression and by introducing spurious correlations through mapping the data to a lower-dimensional representation, other methods yield severely distorted pictures of the data._

### [Gene signature extraction and cell identity recognition at the single-cell level with Cell-ID](https://www.nature.com/articles/s41587-021-00896-6)

_Cortal et al., Nat Biotechnol (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#clustering</a>

_Because of the stochasticity associated with high-throughput single-cell sequencing, current methods for exploring cell-type diversity rely on clustering-based computational approaches in which heterogeneity is characterized at cell subpopulation rather than at full single-cell resolution. Here we present Cell-ID, a clustering-free multivariate statistical method for the robust extraction of per-cell gene signatures from single-cell sequencing data. We applied Cell-ID to data from multiple human and mouse samples, including blood cells, pancreatic islets and airway, intestinal and olfactory epithelium, as well as to comprehensive mouse cell atlas datasets. We demonstrate that Cell-ID signatures are reproducible across different donors, tissues of origin, species and single-cell omics technologies, and can be used for automatic cell-type annotation and cell matching across datasets. Cell-ID improves biological interpretation at individual cell level, enabling discovery of previously uncharacterized rare cell types or cell states. Cell-ID is distributed as an open-source R software package._

### [PCprophet: a framework for protein complex prediction and differential analysis using proteomic data]()

_Fossati et al., Nat Methods (2021)_

<a href="#" class="badge badge-primary">#scrna-seq</a>

_Despite the availability of methods for analyzing protein complexes, systematic analysis of complexes under multiple conditions remains challenging. Approaches based on biochemical fractionation of intact, native complexes and correlation of protein profiles have shown promise. However, most approaches for interpreting cofractionation datasets to yield complex composition and rearrangements between samples depend considerably on protein–protein interaction inference. We introduce PCprophet, a toolkit built on size exclusion chromatography–sequential window acquisition of all theoretical mass spectrometry (SEC-SWATH-MS) data to predict protein complexes and characterize their changes across experimental conditions. We demonstrate improved performance of PCprophet over state-of-the-art approaches and introduce a Bayesian approach to analyze altered protein–protein interactions across conditions. We provide both command-line and graphical interfaces to support the application of PCprophet to any cofractionation MS dataset, independent of separation or quantitative liquid chromatography–MS workflow, for the detection and quantitative tracking of protein complexes and their physiological dynamics._

---

## <i class="fas fa-toolbox"></i> Tools

### [Jupyter-flex: Dashboards for Jupyter](https://jupyter-flex.danielfrg.com/)

Build dashboard using Jupyter Notebooks.

### [First Steps #5: Pluto.jl](https://www.juliafordatascience.com/first-steps-5-pluto/)

Pluto.jl is a newcomer (PlutoCon 2021 was just held to celebrate its one-year anniversary!) to the world of notebook environments.  It provides a reactive environment specific to Julia.  People are doing some very cool things with Pluto.  Check out MIT's Introduction to Computational Thinking course for some fantastic public lectures with Pluto.

### [davidbrochart/nbterm](https://github.com/davidbrochart/nbterm)

Jupyter Notebooks in the terminal.

### [achiku/planter](https://github.com/achiku/planter)

Generate PlantUML ER diagram textual description from PostgreSQL tables.

### [arzzen/git-quick-stats](https://github.com/arzzen/git-quick-stats)

▁▅▆▃▅ Git quick statistics is a simple and efficient way to access various statistics in git repository.

### [BurntSushi/rust-snappy](https://github.com/BurntSushi/rust-snappy)

Snappy compression implemented in Rust (including the Snappy frame format).

### [QingWei-Li/notea](https://github.com/QingWei-Li/notea)

📒 Self hosted note taking app stored on S3.

### [Bhupesh-V/ugit](https://github.com/Bhupesh-V/ugit)

🚨️ ugit helps you undo your last git command with grace. Your damage control git buddy.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [The hidden performance overhead of Python C extensions](https://pythonspeed.com/articles/python-extension-performance/)

Python is slow, and compiled languages like Rust, C, or C++ are fast. So when your application is too slow, rewriting some of your code in a compiled extension can seem like the natural approach to speeding things up.

### [Building Large Scale Systems and Products with Python](https://engineering.soroco.com/building-large-scale-systems-and-products-with-python/)

### [labml.ai Neural Networks](https://nn.labml.ai/)

This is a collection of simple PyTorch implementations of neural networks and related algorithms. These implementations are documented with explanations, and the website renders these as side-by-side formatted notes. We believe these would help you understand these algorithms better.

### [Toggle dark/light mode by clapping your hands](https://charliegerard.dev/blog/toggle-dark-mode-clapping-hands-chrome-extension/)

How to build it, and ended up with a Chrome extension using TensorFlow.js and a model trained with samples of me clapping my hands, that toggles dark mode on/off!

### [Encrypting and decrypting files with OpenSSL](https://opensource.com/article/21/4/encryption-decryption-openssl)

OpenSSL is a practical tool for ensuring your sensitive and secret messages can't be opened by outsiders.

---

## <i class="fas fa-rss"></i> Others

* [Sell yourself and your science in a compelling personal statement](https://www.nature.com/articles/d41586-021-01101-z)
* [30 Years Of Linux - An Interview With Linus Torvalds: Linux and Git](https://www.tag1consulting.com/blog/interview-linus-torvalds-linux-and-git)
* [A brief history of Rust at Facebook](https://engineering.fb.com/2021/04/29/developer-tools/rust/)
* [Why We Built Our Own DNS Infrastructure](https://blog.replit.com/dns)
* [Crypto miners are killing free CI](https://layerci.com/blog/crypto-miners-are-killing-free-ci/)
* [Lessons from the pandemic’s superstar data scientist, Youyang Gu](https://www.technologyreview.com/2021/04/27/1023657/lessons-from-the-pandemics-superstar-data-scientist-youyang-gu/)

---

## <i class="far fa-surprise"></i> Did you know?

<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Anatoli Bugorski is the only men, who was struck by a particle accelerator beam<a href="https://t.co/FM8lR3Z9ep">https://t.co/FM8lR3Z9ep</a> <a href="https://t.co/m5MnSLvMk4">pic.twitter.com/m5MnSLvMk4</a></p>&mdash; Lost in Wikipedia (@lostinwiki) <a href="https://twitter.com/lostinwiki/status/713357374631882752?ref_src=twsrc%5Etfw">March 25, 2016</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
