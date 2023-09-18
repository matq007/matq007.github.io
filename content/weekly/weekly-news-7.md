---
title: "Weekly news #7"
date: 2021-02-14
tags: [weekly]
---

## <i class="fas fa-bullhorn"></i> News

* [Stanford Makes iPhone Development Course Free](https://cs193p.sites.stanford.edu/)
* [Rust gets its own foundation](https://foundation.rust-lang.org/posts/2021-02-08-hello-world/)

---

## <i class="fas fa-dna"></i> Science/Bioinformatics

### [The next 20 years of human genomics must be more equitable and more open](https://www.nature.com/articles/d41586-021-00328-0)

_DOI: 10.1038/d41586-021-00328-0_

### [Brain2Pix: Fully convolutional naturalistic video reconstruction from brain activity](https://www.biorxiv.org/content/10.1101/2021.02.02.429430v1)

_Le et al., biorvix 2021_

<a href="#" class="badge badge-primary">#neural-network</a>
<a href="#" class="badge badge-primary">#fmri</a>
<a href="#" class="badge badge-primary">#image-reconstruction</a>

_Reconstructing complex and dynamic visual perception from brain activity remains a major challenge in machine learning applications to neuroscience. Here we present a new method for reconstructing naturalistic images and videos from very large single-participant functional magnetic resonance data that leverages the recent success of image-to-image transformation networks. This is achieved by exploiting spatial information obtained from retinotopic mappings across the visual system. More specifically, we first determine what position each voxel in a particular region of interest would represent in the visual field based on its corresponding receptive field location. Then, the 2D image representation of the brain activity on the visual field is passed to a fully convolutional image-to-image network trained to recover the original stimuli using VGG feature loss with an adversarial regularizer. In our experiments, we show that our method offers a significant improvement over existing video reconstruction techniques._

### [Compression of quantification uncertainty for scRNA-seq counts](https://academic.oup.com/bioinformatics/advance-article/doi/10.1093/bioinformatics/btab001/6104828)

_Buren et al., Bioinformatics 2021_

<a href="#" class="badge badge-primary">#scrna-seq</a>
<a href="#" class="badge badge-primary">#compression</a>

_Quantification estimates of gene expression from single-cell RNA-seq (scRNA-seq) data have inherent uncertainty due to reads that map to multiple genes. Many existing scRNA-seq quantification pipelines ignore multi-mapping reads and therefore underestimate expected read counts for many genes. alevin accounts for multi-mapping reads and allows for the generation of ‘inferential replicates’, which reflect quantification uncertainty. Previous methods have shown improved performance when incorporating these replicates into statistical analyses, but storage and use of these replicates increases computation time and memory requirements. We demonstrate that storing only the mean and variance from a set of inferential replicates (‘compression’) is sufficient to capture gene-level quantification uncertainty, while reducing disk storage to as low as 9% of original storage, and memory usage when loading data to as low as 6%. Using these values, we generate ‘pseudo-inferential’ replicates from a negative binomial distribution and propose a general procedure for incorporating these replicates into a proposed statistical testing framework. When applying this procedure to trajectory-based differential expression analyses, we show false positives are reduced by more than a third for genes with high levels of quantification uncertainty. We additionally extend the Swish method to incorporate pseudo-inferential replicates and demonstrate improvements in computation time and memory usage without any loss in performance. Lastly, we show that discarding multi-mapping reads can result in significant underestimation of counts for functionally important genes in a real dataset._

### [Ten simple rules for engaging with artificial intelligence in biomedicine](https://journals.plos.org/ploscompbiol/article?id=10.1371%2Fjournal.pcbi.1008531)

_Malik et at., Plos Comp Bio 2021_

<a href="#" class="badge badge-primary">#ai</a>
<a href="#" class="badge badge-primary">#tips-and-tricks</a>
<a href="#" class="badge badge-primary">#image-processing</a>

### [Origins of modern human ancestry](https://www.nature.com/articles/s41586-021-03244-5)

_Bergström et al., Nature 2021_

<a href="#" class="badge badge-primary">#history</a>

_New finds in the palaeoanthropological and genomic records have changed our view of the origins of modern human ancestry. Here we review our current understanding of how the ancestry of modern humans around the globe can be traced into the deep past, and which ancestors it passes through during our journey back in time. We identify three key phases that are surrounded by major questions, and which will be at the frontiers of future research. The most recent phase comprises the worldwide expansion of modern humans between 40 and 60 thousand years ago (ka) and their last known contacts with archaic groups such as Neanderthals and Denisovans. The second phase is associated with a broadly construed African origin of modern human diversity between 60 and 300 ka. The oldest phase comprises the complex separation of modern human ancestors from archaic human groups from 0.3 to 1 million years ago. We argue that no specific point in time can currently be identified at which modern human ancestry was confined to a limited birthplace, and that patterns of the first appearance of anatomical or behavioural traits that are used to define Homo sapiens are consistent with a range of evolutionary histories._

---

## <i class="far fa-keyboard"></i> Programming

### [All Pythons are slow, but some are faster than others](https://pythonspeed.com/articles/faster-python/)

Looks like compiling python makes a quite of a difference.

### [satwikkansal/wtfpython](https://github.com/satwikkansal/wtfpython)

What the f*ck Python?

---

## <i class="fas fa-toolbox"></i> Tools

### [conwnet/github1s](https://github.com/conwnet/github1s)

Open any GitHub repository in Visual Studio Code right in your browser.

### [kloudi](https://kloudi.tech/)

Enter commands and queries to search, view and perform actions on data from all your engineering tools.

### [jarred-sumner/git-peek](https://github.com/jarred-sumner/git-peek)

Fastest way to open a remote git repository in your local text editor.

### [Sourcetrail](https://www.sourcetrail.com/)

Free and open-source cross-platform source explorer

### [Privacy-First AI Security Camera App](https://www.ai-cam.app/)

---

## <i class="fas fa-graduation-cap"></i> Guides and Tutorials

### [Containerized Python Development – Part 1](https://www.docker.com/blog/containerized-python-development-part-1/)

Guides from Docker on how to containerized your next Python application.

### [Shrink your Conda Docker images with conda-pack](https://pythonspeed.com/articles/conda-docker-image-size/)

Must read if you are using conda in docker container.

---

## <i class="fas fa-rss"></i> Others

### [Dependency Confusion: How I Hacked Into Apple, Microsoft and Dozens of Other Companies](https://medium.com/@alex.birsan/dependency-confusion-4a5d60fec610)

### [Phishing with Morse code](https://hackaday.com/2021/02/11/phishing-with-morse-code/)

### [Information Processing](https://infoproc.blogspot.com/2021/02/gradient-descent-models-are-kernel.html)

### [DECIDE-AI: new reporting guidelines to bridge the development-to-implementation gap in clinical artificial intelligence](https://www.nature.com/articles/s41591-021-01229-5)

### [Microsoft Excel Becomes a Programming Language](https://thenewstack.io/microsoft-excel-becomes-a-programming-language)

### [New research tackles a central challenge of powerful quantum computing](https://phys.org/news/2021-02-tackles-central-powerful-quantum.html)

### [You Shall Not Pass](https://youshallnotpass.glitch.me/?)

---

## <i class="far fa-surprise"></i> Did you know?

Researchers are working on a novel approach to treat skin cancer - melanoma by
altering microbiome using fecal transplant. In the study, fecal samples were
collected from patients who responded well to normal immunotherapy. The fecal
samples were processed and transplanted to patients who didn't respond to immunotherapy
at all with additional administration of pembrolizumab. Follow-up tumor biopsy showed
that 6 ot ouf 15 patients had their tumors shrank or remained its original size.
Interestingly, further analysis revealed an increase of bacteria associated with
immunotherapy response. Read the full story at [Science](https://science.sciencemag.org/content/371/6529/595).

_Davar et al., Science 2021_