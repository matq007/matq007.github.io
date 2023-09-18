---
title: "Weekly news #8"
date: 2021-02-21
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Mars Perseverance Photo Booth](https://mars.nasa.gov/mars2020/participate/photo-booth/)
* [Huston, the Rover landed on Mars!](https://www.nasa.gov/content/live-landing-of-the-mars-2020-perseverance-rover)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Blame the Data](https://michelebusby.tumblr.com/post/643211974587629568/so-you-want-to-start-a-biotech-a-bioinformatics)

### [Ten simple rules for getting started with command-line bioinformatics](https://journals.plos.org/ploscompbiol/article?id=10.1371%2Fjournal.pcbi.1008645)

_Brandies et al., PLoS Comput Biol 2021_

<a href="#" class="badge badge-primary">#command-line</a>

### [Robust decomposition of cell type mixtures in spatial transcriptomics](https://www.nature.com/articles/s41587-021-00830-w)

_Cable et al., Nat Biotechnol 2021_

<a href="#" class="badge badge-primary">#spatial-transcriptomics</a>

_A limitation of spatial transcriptomics technologies is that individual measurements may contain contributions from multiple cells, hindering the discovery of cell-type-specific spatial patterns of localization and expression. Here, we develop robust cell type decomposition (RCTD), a computational method that leverages cell type profiles learned from single-cell RNA-seq to decompose cell type mixtures while correcting for differences across sequencing technologies. We demonstrate the ability of RCTD to detect mixtures and identify cell types on simulated datasets. Furthermore, RCTD accurately reproduces known cell type and subtype localization patterns in Slide-seq and Visium datasets of the mouse brain. Finally, we show how RCTD’s recovery of cell type localization enables the discovery of genes within a cell type whose expression depends on spatial environment. Spatial mapping of cell types with RCTD enables the spatial components of cellular identity to be defined, uncovering new principles of cellular organization in biological tissue. RCTD is publicly available as an open-source R package at https://github.com/dmcable/RCTD._

### [Genozip - A Universal Extensible Genomic Data Compressor](https://academic.oup.com/bioinformatics/advance-article/doi/10.1093/bioinformatics/btab102/6135077)

_Divon et al., Bioinformatics 2021_

<a href="#" class="badge badge-primary">#compression</a>

_We present Genozip, a universal and fully featured compression software for genomic data. Genozip is designed to be a general-purpose software and a development framework for genomic compression by providing five core capabilities – universality (support for all common genomic file formats), high compression ratios, speed, feature-richness, and extensibility. Genozip delivers high-performance compression for widely-used genomic data formats in genomics research, namely FASTQ, SAM/BAM/CRAM, VCF, GVF, FASTA, PHYLIP, and 23andMe formats. Our test results show that Genozip is fast and achieves greatly improved compression ratios, even when the files are already compressed. Further, Genozip is architected with a separation of the Genozip Framework from file-format-specific Segmenters and data-type-specific Codecs. With this, we intend for Genozip to be a general-purpose compression platform where researchers can implement compression for additional file formats, as well as new codecs for data types or fields within files, in the future. We anticipate that this will ultimately increase the visibility and adoption of these algorithms by the user community, thereby accelerating further innovation in this space._

### [Dimensionality reduction using singular vectors](https://www.nature.com/articles/s41598-021-83150-y)

_Afshar et al., Sci Rep 2021_

<a href="#" class="badge badge-primary">#feature-selection</a>
<a href="#" class="badge badge-primary">#singular-vectors</a>

_A common problem in machine learning and pattern recognition is the process of identifying the most relevant features, specifically in dealing with high-dimensional datasets in bioinformatics. In this paper, we propose a new feature selection method, called Singular-Vectors Feature Selection (SVFS). Let 𝐷=[𝐴∣𝐛] be a labeled dataset, where 𝐛 is the class label and features (attributes) are columns of matrix A. We show that the signature matrix 𝑆𝐴=𝐼−𝐴†𝐴 can be used to partition the columns of A into clusters so that columns in a cluster correlate only with the columns in the same cluster. In the first step, SVFS uses the signature matrix 𝑆𝐷 of D to find the cluster that contains 𝐛. We reduce the size of A by discarding features in the other clusters as irrelevant features. In the next step, SVFS uses the signature matrix 𝑆𝐴 of reduced A to partition the remaining features into clusters and choose the most important features from each cluster. Even though SVFS works perfectly on synthetic datasets, comprehensive experiments on real world benchmark and genomic datasets shows that SVFS exhibits overall superior performance compared to the state-of-the-art feature selection methods in terms of accuracy, running time, and memory usage. A Python implementation of SVFS along with the datasets used in this paper are available at https://github.com/Majid1292/SVFS._

---

## <i class="far fa-keyboard"></i> Programming

### [Building Rich terminal dashboards](https://www.willmcgugan.com/blog/tech/post/building-rich-terminal-dashboards/)

If you haven't heard of package `Rich` for python, I strongly encourage to check it out.
You can even build a full dashboard with it.

### [Write the docs](https://www.writethedocs.org/)

Every application/code needs a proper documentation.

### [Free for developers](https://free-for.dev/)

Resource list of all free tier services for development.

---

## <i class="fas fa-toolbox"></i> Tools

### [apankrat/nullboard](https://github.com/apankrat/nullboard)

Nullboard is a minimalist kanban board, focused on compactness and readability.

### [Sublime Merge](https://www.sublimemerge.com/)

Git Client, done Sublime - Line-by-line Staging. Commit Editing. Unmatched Performance.

### [Ray](https://ray.so/)

Pretty code as image.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Reasons why SELECT * is bad for SQL performance](https://tanelpoder.com/posts/reasons-why-select-star-is-bad-for-sql-performance/)

In depth analysis of `SELECT * FROM`.

### [Advanced Git Features You Didn’t Know You Needed](https://martinheinz.dev/blog/43)

### [Why you really need to upgrade pip](https://pythonspeed.com/articles/upgrade-pip/)

Upgrade, upgrade and upgrade.

### [Introducing Model Search: An Open Source Platform for Finding Optimal ML Models](https://ai.googleblog.com/2021/02/introducing-model-search-open-source.html)

Model search (MS) is a framework that implements AutoML algorithms for model architecture search at scale. It aims to help researchers speed up their exploration process for finding the right model architecture for their classification problems (i.e., DNNs with different types of layers).

### [Efficiently Cleaning Text with Pandas](https://pbpython.com/text-cleaning.html)

Did you know `str.contains` has extra parameters in pandas?

---

## <i class="fas fa-rss"></i> Others

### [Are The New M1 Macbooks Any Good for Deep Learning? Let’s Find Out](https://www.betterdatascience.com/m1-deep-learning/)

### [Mitigating Memory Safety Issues in Open Source Software](https://security.googleblog.com/2021/02/mitigating-memory-safety-issues-in-open.html)

Google approves Rust.

### [What on Earth is this Encryption Scheme?](https://capnfabs.net/posts/wtf-encryption-scheme-synology-diskstation-nas/)

Synology NAS and its encryption.

---

## <i class="far fa-surprise"></i> Did you know?

Have you ever wondered why some people are more resilient towards cold? Is it
only the climate environment, or is there something more to it? Well, researchers
from Karolinska Institutet discovered that people who lack $\alpha-activin-3$ are
able to maintain warm better due to changes in skeletal muscle thermogensis. The
lack of protein results in higher ratio of slow-twitch compared to fast-twitch
fibers. The researchers speculate this mutation has a evolutionary origin as people
started to migrated from Africa to central and northern Europe. Read the full
study [here](https://www.cell.com/ajhg/fulltext/S0002-9297(21)00013-6).

_Wyckelsma et al., AJHG 2021_