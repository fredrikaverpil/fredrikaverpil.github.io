---
date: 2017-08-28
authors:
  - fredrikaverpil
comments: true
tags:
- python
- pyside
- pyqt
- qt.py
---

# PySide2 easy install!

Yesterday, [@jschueller](https://github.com/jschueller) added [`pyside2-feedstock`](https://github.com/conda-forge/pyside2-feedstock) to [conda-forge](https://conda-forge.org). This means we can now *finally* install PySide2 **easily** in Python 2.7, 3.5 and 3.6 on Windows, Linux and macOS using conda.

<!-- more -->

```bash
# Enable conda-forge
conda config --add channels conda-forge

# Install PySide2
conda install pyside2
```

And. It. Frickin'. Just. Works.  

**Update 2018-03-09**: The Qt Company now offers official and standalone wheels, read more [here](2018-03-09-official-pyside2-wheels.md).  
**Update 2018-07-17**: PySide2 can now be installed from pypi.org: `pip install PySide2`!

## Install conda

I prefer to use the [miniconda](https://conda.io/miniconda.html) distribution of Conda, which means you just get conda without any extras.

Below, some examples on how to install miniconda silently. As usual, more info in [their docs](https://conda.io/docs/user-guide/install/macos.html#installing-in-silent-mode) on this.

### macOS

```bash
# Download
wget https://repo.continuum.io/miniconda/Miniconda3-latest-MacOSX-x86_64.sh -O ~/miniconda.sh

# Install
bash ~/miniconda.sh -b -p $HOME/miniconda3

# Add "conda" to $PATH
echo 'export PATH="$HOME/miniconda3/bin:$PATH"' >> ~/.bashrc
```

### Linux

```bash
# Download
wget https://repo.continuum.io/miniconda/Miniconda3-latest-Linux-x86_64.sh -O ~/miniconda.sh

# Install
bash ~/miniconda.sh -b -p $HOME/miniconda3

# Add "conda" to $PATH
echo 'export PATH="$HOME/miniconda3/bin:$PATH"' >> ~/.bashrc
```

### Windows (powershell)


```powershell
# Download
(New-Object System.Net.WebClient).DownloadFile("https://repo.continuum.io/miniconda/Miniconda3-latest-Windows-x86_64.exe", "${HOME}\Downloads\Miniconda3-latest-Windows-x86_64.exe")

# Install, add "conda" and "python" etc to $PATH
& "${HOME}\Downloads\Miniconda3-latest-Windows-x86_64.exe" /S /AddToPath=1 /D=${HOME}\miniconda3
```

## Install conda environment

This is the same for all platforms, given that `conda` exists on `$PATH` and that you're using Powershell if on Windows.

```bash
# Enable conda-forge
conda config --add channels conda-forge

# Create environment with Python 3.6, PySide2
conda create --mkdir --prefix ~/condaenvs/myenv python=3.6 pyside2

# Run Python
~/condaenvs/myenv/bin/python --version

# Run pip
~/condaenvs/myenv/bin/pip --version
```