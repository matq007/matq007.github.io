---
layout: weekly
title: "Weekly news #4"
date: 2021-01-24
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [RStudio 1.4](https://blog.rstudio.com/2021/01/19/announcing-rstudio-1-4/) is out 🎉
* [Raspberry Pi Pico](https://www.raspberrypi.org/blog/raspberry-pi-silicon-pico-now-on-sale/) now on sale at $4
* [IPFS Support in Brave browser](https://brave.com/ipfs-support/)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [Ten computer codes that transformed science](https://www.nature.com/articles/d41586-021-00075-2)

_Jeffrey M. Perkel_

_From Fortran to arXiv.org, these advances in programming and platforms sent biology, climate science and physics into warp speed._

### [Noninvasive Precision Screening of Prostate Cancer by Urinary Multimarker Sensor and Artificial Intelligence Analysis](https://pubs.acs.org/doi/10.1021/acsnano.0c06946)

_Hojun et al., ACS Nano, 2021_

<a href="#" class="badge badge-primary">#prostate-cancer</a>
<a href="#" class="badge badge-primary">#ai</a>

_Screening for prostate cancer relies on the serum prostate-specific antigen test, which provides a high rate of false positives (80%). This results in a large number of unnecessary biopsies and subsequent overtreatment. Considering the frequency of the test, there is a critical unmet need of precision screening for prostate cancer. Here, we introduced a urinary multimarker biosensor with a capacity to learn to achieve this goal. The correlation of clinical state with the sensing signals from urinary multimarkers was analyzed by two common machine learning algorithms. As the number of biomarkers was increased, both algorithms provided a monotonic increase in screening performance. Under the best combination of biomarkers, the machine learning algorithms screened prostate cancer patients with more than 99% accuracy using 76 urine specimens. Urinary multimarker biosensor leveraged by machine learning analysis can be an important strategy of precision screening for cancers using a drop of bodily fluid._

### [Assessing single-cell transcriptomic variability through density-preserving data visualization](https://www.nature.com/articles/s41587-020-00801-7#data-availability)

_Narayan et al., Nat Biotechnol, 2021_

<a href="#" class="badge badge-primary">#visualization</a>
<a href="#" class="badge badge-primary">#scrna-seq</a>

_Nonlinear data visualization methods, such as t-distributed stochastic neighbor embedding (t-SNE) and uniform manifold approximation and projection (UMAP), summarize the complex transcriptomic landscape of single cells in two dimensions or three dimensions, but they neglect the local density of data points in the original space, often resulting in misleading visualizations where densely populated subsets of cells are given more visual space than warranted by their transcriptional diversity in the dataset. Here we present den-SNE and densMAP, which are density-preserving visualization tools based on t-SNE and UMAP, respectively, and demonstrate their ability to accurately incorporate information about transcriptomic variability into the visual interpretation of single-cell RNA sequencing data. Applied to recently published datasets, our methods reveal significant changes in transcriptomic variability in a range of biological processes, including heterogeneity in transcriptomic variability of immune cells in blood and tumor, human immune cell specialization and the developmental trajectory of Caenorhabditis elegans. Our methods are readily applicable to visualizing high-dimensional data in other scientific domains._

### [Deep neural networks identify sequence context features predictive of transcription factor binding](https://www.nature.com/articles/s42256-020-00282-y)

_Zheng et al., Nat Mach Intell, 2021_

<a href="#" class="badge badge-primary">#ai</a>
<a href="#" class="badge badge-primary">#prediction</a>
<a href="#" class="badge badge-primary">#transcription-factor</a>
<a href="#" class="badge badge-primary">#chip-seq</a>

_Transcription factors bind DNA by recognizing specific sequence motifs, which are typically 6–12 bp long. A motif can occur many thousands of times in the human genome, but only a subset of those sites are actually bound. Here we present a machine-learning framework leveraging existing convolutional neural network architectures and model interpretation techniques to identify and interpret sequence context features most important for predicting whether a particular motif instance will be bound. We apply our framework to predict binding at motifs for 38 transcription factors in a lymphoblastoid cell line, score the importance of context sequences at base-pair resolution and characterize context features most predictive of binding. We find that the choice of training data heavily influences classification accuracy and the relative importance of features such as open chromatin. Overall, our framework enables novel insights into features predictive of transcription factor binding and is likely to inform future deep learning applications to interpret non-coding genetic variants._

### [AutoMap is a high performance homozygosity mapping tool using next-generation sequencing data](https://www.nature.com/articles/s41467-020-20584-4)

_Quinodoz et al., Nat Commun, 2021_

<a href="#" class="badge badge-primary">#homozygosity-mapping</a>

_Homozygosity mapping is a powerful method for identifying mutations in patients with recessive conditions, especially in consanguineous families or isolated populations. Historically, it has been used in conjunction with genotypes from highly polymorphic markers, such as DNA microsatellites or common SNPs. Traditional software performs rather poorly with data from Whole Exome Sequencing (WES) and Whole Genome Sequencing (WGS), which are now extensively used in medical genetics. We develop AutoMap, a tool that is both web-based or downloadable, to allow performing homozygosity mapping directly on VCF (Variant Call Format) calls from WES or WGS projects. Following a training step on WES data from 26 consanguineous families and a validation procedure on a matched cohort, our method shows higher overall performances when compared with eight existing tools. Most importantly, when tested on real cases with negative molecular diagnosis from an internal set, AutoMap detects three gene-disease and multiple variant-disease associations that were previously unrecognized, projecting clear benefits for both molecular diagnosis and research activities in medical genetics._

### [A spatially resolved brain region- and cell type-specific isoform atlas of the postnatal mouse brain](https://www.nature.com/articles/s41467-020-20343-5)

<a href="#" class="badge badge-primary">#isoform</a>
<a href="#" class="badge badge-primary">#long-read</a>
<a href="#" class="badge badge-primary">#spatial-transcriptomics</a>

_Splicing varies across brain regions, but the single-cell resolution of regional variation is unclear. We present a single-cell investigation of differential isoform expression (DIE) between brain regions using single-cell long-read sequencing in mouse hippocampus and prefrontal cortex in 45 cell types at postnatal day 7 ([www.isoformAtlas.com](www.isoformAtlas.com)). Isoform tests for DIE show better performance than exon tests. We detect hundreds of DIE events traceable to cell types, often corresponding to functionally distinct protein isoforms. Mostly, one cell type is responsible for brain-region specific DIE. However, for fewer genes, multiple cell types influence DIE. Thus, regional identity can, although rarely, override cell-type specificity. Cell types indigenous to one anatomic structure display distinctive DIE, e.g. the choroid plexus epithelium manifests distinct transcription-start-site usage. Spatial transcriptomics and long-read sequencing yield a spatially resolved splicing map. Our methods quantify isoform expression with cell-type and spatial resolution and it contributes to further our understanding of how the brain integrates molecular and cellular complexity._

---

## <i class="far fa-keyboard"></i> Programming

### [Improving Indian Language Transliterations in Google Maps](https://ai.googleblog.com/2021/01/improving-indian-language.html)

### [Google Details Patched Bugs in Signal, FB Messenger, JioChat Apps](https://thehackernews.com/2021/01/google-discloses-flaws-in-signal-fb.html)

### [Why you should never ever use NixOS](https://hands-on.cloud/why-you-should-never-ever-use-nixos/)

Some interesting points about NixOS nad its configuration.

---

## <i class="fas fa-toolbox"></i> Tools

### [mackup](https://github.com/lra/mackup)

Keep your application settings in sync (OS X/Linux)

### [covariants](https://github.com/hodcroftlab/covariants)

Real-time updates and information about key SARS-CoV-2 variants, plus the scripts that generate this information.

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Build 12 Data Science Apps with Python and Streamlit - Full Course](https://www.youtube.com/watch?v=JwSS70SZdyM)

Replace that ShinyApp with Streamlit.

### [From Novice to Expert: How to Write a Configuration file in Python](https://towardsdatascience.com/from-novice-to-expert-how-to-write-a-configuration-file-in-python-273e171a8eb3)

Multiple ways how to create configuration file for Python: `YAML/JSON`, `configparser`, `python-dotenv` or `dynaconf`.

### [Testing Python Applications with Pytest [Guide]](https://stribny.name/blog/pytest/)

Guide how to write tests in python with coverage reports.

---

## <i class="fas fa-rss"></i> Others

### [Imposters hijack journal’s peer review process to publish substandard papers](https://www.chemistryworld.com/news/imposters-hijack-journals-peer-review-process-to-publish-substandard-papers/4013050.article)

### [Award-winning mentors share their secrets](https://www.nature.com/articles/d41586-021-00081-4)

---

## <i class="far fa-surprise"></i> Did you know?

Since the beginning of the pandemic, Island has sequenced so far 6,000 positive
cases of COVID-19 to identify various strains of the virus. The main wheel behind
the sequencing is local biopharm called _deCODE Genetics_ situated in the capital
Reykjavik. The whole process from isolating DNA to sequencing takes roughly a
day and a half, says _Olafur Thor Magnusson_, head of the lab. So far, no obvious
mutation patterns have been identified, says the founder _Kári Stefánsson_. Read
the whole article [here](https://www.sciencealert.com/iceland-tracks-and-contains-covid-19-by-genetically-sequencing-every-positive-case).
