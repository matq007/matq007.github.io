---
title: "Weekly news #5"
date: 2021-01-31
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Announcing the 2020 Golden Kitty Award Winners](https://www.producthunt.com/stories/announcing-the-2020-golden-kitty-award-winners)
* [Postgres broken on Mac OS Big Sur](https://github.com/PostgresApp/PostgresApp/issues/610)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Seven myths of how transcription factors read the cis-regulatory code](https://www.sciencedirect.com/science/article/pii/S2452310020300305)

_Zeitlinger et al., Current Opinion 2021_

<a href="#" class="badge badge-primary">#transcription-factors</a>
<a href="#" class="badge badge-primary">#enhancers</a>

_Genomics data are now being generated at large quantities, of exquisite high resolution and from single cells. They offer a unique opportunity to develop powerful machine learning algorithms, including neural networks, to uncover the rules of the cis-regulatory code. However, current modeling assumptions are often not based on state-of-the-art knowledge of the cis-regulatory code from transcription, developmental genetics, imaging, and structural studies. Here I aim to fill this gap by giving a brief historical overview of the field, describing common misconceptions and providing knowledge that might help to guide computational approaches. I will describe the principles and mechanisms involved in the combinatorial requirement of transcription factor binding motifs for enhancer activity, including the role of chromatin accessibility, repressors, and low-affinity motifs in the cis-regulatory code. Deciphering the cis-regulatory code would unlock an enormous amount of regulatory information in the genome and would allow us to locate cis-regulatory genetic variants involved in development and disease._

### [Probabilistic harmonization and annotation of single‐cell transcriptomics data with deep generative models](https://www.embopress.org/doi/full/10.15252/msb.20209620)

_Xu et al., Mol Syst Biol 2021_

<a href="#" class="badge badge-primary">#autoencoders</a>
<a href="#" class="badge badge-primary">#scvi</a>
<a href="#" class="badge badge-primary">#scanvi</a>

_As the number of single‐cell transcriptomics datasets grows, the natural next step is to integrate the accumulating data to achieve a common ontology of cell types and states. However, it is not straightforward to compare gene expression levels across datasets and to automatically assign cell type labels in a new dataset based on existing annotations. In this manuscript, we demonstrate that our previously developed method, scVI, provides an effective and fully probabilistic approach for joint representation and analysis of scRNA‐seq data, while accounting for uncertainty caused by biological and measurement noise. We also introduce single‐cell ANnotation using Variational Inference (scANVI), a semi‐supervised variant of scVI designed to leverage existing cell state annotations. We demonstrate that scVI and scANVI compare favorably to state‐of‐the‐art methods for data integration and cell state annotation in terms of accuracy, scalability, and adaptability to challenging settings. In contrast to existing methods, scVI and scANVI integrate multiple datasets with a single generative model that can be directly used for downstream tasks, such as differential expression. Both methods are easily accessible through scvi‐tools._

### [tidybulk: an R tidy framework for modular transcriptomic data analysis](https://genomebiology.biomedcentral.com/articles/10.1186/s13059-020-02233-7)

_Mangiola et al., Genome Biol 2021_

<a href="#" class="badge badge-primary">#r</a>
<a href="#" class="badge badge-primary">#bulk</a>
<a href="#" class="badge badge-primary">#framework</a>

_Recently, efforts have been made toward the harmonization of transcriptomic data structures and workflows using the concept of data tidiness, to facilitate modularisation. We present tidybulk, a modular framework for bulk transcriptional analyses that introduces a tidy transcriptomic data structure paradigm and analysis grammar. Tidybulk covers a wide variety of analysis procedures and integrates a large ecosystem of publicly available analysis algorithms under a common framework. Tidybulk decreases coding burden, facilitates reproducibility, increases efficiency for expert users, lowers the learning curve for inexperienced users, and bridges transcriptional data analysis with the tidyverse. Tidybulk is available at R/Bioconductor bioconductor.org/packages/tidybulk._

### [Benchmarking of lightweight-mapping based single-cell RNA-seq pre-processing](https://www.biorxiv.org/content/10.1101/2021.01.25.428188v1)

_Booeshaghi et al., bioRvix 2021_

<a href="#" class="badge badge-primary">#alvin</a>
<a href="#" class="badge badge-primary">#kallisto</a>
<a href="#" class="badge badge-primary">#benchmarking</a>
<a href="#" class="badge badge-primary">#scrna-seq</a>

_We compare and benchmark the two lightweight-mapping tools that have been developed for pre-processing single-cell RNA-seq data, namely the kallisto-bustools and Salmon-Alevin-fry programs. We find that they output similar results, and to the extent that there are differences, they are irrelevant for downstream analysis. However, the Salmon-Alevin-fry program is significantly slower and requires much more memory to run, making it much more expensive to process large datasets limiting its use to larger servers._

### [Nanobody-mediated control of gene expression and epigenetic memory](https://www.nature.com/articles/s41467-020-20757-1)

_Van et al., Nat Commun 2021_

<a href="#" class="badge badge-primary">#nanobody</a>
<a href="#" class="badge badge-primary">#gene-expression</a>

_Targeting chromatin regulators to specific genomic locations for gene control is emerging as a powerful method in basic research and synthetic biology. However, many chromatin regulators are large, making them difficult to deliver and combine in mammalian cells. Here, we develop a strategy for gene control using small nanobodies that bind and recruit endogenous chromatin regulators to a gene. We show that an antiGFP nanobody can be used to simultaneously visualize GFP-tagged chromatin regulators and control gene expression, and that nanobodies against HP1 and DNMT1 can silence a reporter gene. Moreover, combining nanobodies together or with other regulators, such as DNMT3A or KRAB, can enhance silencing speed and epigenetic memory. Finally, we use the slow silencing speed and high memory of antiDNMT1 to build a signal duration timer and recorder. These results set the basis for using nanobodies against chromatin regulators for controlling gene expression and epigenetic memory._

---

## <i class="far fa-keyboard"></i> Programming

### [Discover the best developer blogs on any tech stack](https://bloggingfordevs.com/trends/)

Is meant to be a fun way to discover excellent tech blogs, made by and for developers! Here's how it works.

### [Improving how we deploy GitHub](https://github.blog/2021-01-25-improving-how-we-deploy-github/)

Insight into how GitHub changed their deployment strategy to keep track of everything
with increasing number of developers.

### [How to Read Rust Functions, Part 1](https://www.possiblerust.com/guide/how-to-read-rust-functions-part-1)

Don't fear rust and its syntax. Master how to write a function instead.

### [R language for programmers](https://www.johndcook.com/blog/r_language_for_programmers/)

If you are coming from different programming language, this article will speed you
up with the R syntax and types.

---

## <i class="fas fa-toolbox"></i> Tools

### [rust-starter](https://github.com/rust-starter/rust-starter)

Starter boilerplate to create a Rust CLI application. It comes with a battery of components like argument parsing and configuration. It also has different tooling to create your binary, or automate your build process.

### [micro](https://github.com/micro/micro)

Platform for cloud native development.

### [starship](https://starship.rs/)

The minimal, blazing-fast, and infinitely customizable prompt for any shell written in Rust.

### [fusen](https://www.r-bloggers.com/2021/01/fusen-create-a-package-from-a-simple-rmarkdown-file/)

Convert RMarkdown file into a full blown R package in just few commands.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Build Your Own Container Using Less than 100 Lines of Go](https://www.infoq.com/articles/build-a-container-golang/)

Super simple implementation of a container written in go.

### [nn](https://github.com/lab-ml/nn)

Minimal implementations/tutorials of deep learning papers with side-by-side notes.

---

## <i class="fas fa-rss"></i> Others

### [Scaling Kubernetes to 7,500 Nodes](https://openai.com/blog/scaling-kubernetes-to-7500-nodes)

### [SpaceX Just Launched 'Most Spacecraft Ever Deployed on a Single Mission'](https://www.sciencealert.com/spacex-just-launched-most-spacecraft-ever-deployed-on-a-single-mission)

---

## <i class="far fa-surprise"></i> Did you know?

Back in 2020 researchers from Madrid, Spain observed 666 patients with
mild-to-moderate COVID-19. They observed swelling and small spots on tongue.
The patients also reported burning sensation, swelling of hands and feet. The
authors mention these studies were limited only to adults and 2-week period,
meaning they could have missed early and late progression of COVID-19. Finally,
the observed dermatology conditions could be associated to hospitalization
conditions as well as ventilation masks.

<details><summary>Figure from Nuno-Gonzalez et al., British Journal of Dermatology, 2020</summary>
<p>

![](https://onlinelibrary.wiley.com/cms/asset/c511fbcb-ffe6-4d7e-b97a-5fbfb5328c33/bjd19564-fig-0001-m.jpg).

</p>
</details>
