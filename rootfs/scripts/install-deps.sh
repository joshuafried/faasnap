#!/usr/bin/env bash

set -ex

echo deb http://archive.ubuntu.com/ubuntu noble universe >> /etc/apt/sources.list
apt update
apt install -y nodejs libgl1 gfortran ruby php-cli python3-pil libopenblas-dev npm python3.12-venv

apt update
apt install -f -y openssh-server build-essential gpg wget libblas3 liblapack3 liblapack-dev libblas-dev gfortran libffi-dev python3.12 python3-pip openjdk-21-jre-headless libcjson-dev

wget https://github.com/Kitware/CMake/releases/download/v3.22.2/cmake-3.22.2-linux-x86_64.tar.gz -O /opt/cmake-3.22.2-linux-x86_64.tar.gz
pushd /opt
tar xzvf cmake-3.22.2-linux-x86_64.tar.gz
export PATH=$PATH:/opt/cmake-3.22.2-linux-x86_64/bin/
popd

python3 -m venv /app/python/

MAKEFLAGS="-j54" /app/python/bin/pip3 install --break-system-packages wheel six scikit-learn flask pillow pyaes chameleon pandas tensorflow torch torchvision minio psutil keras-preprocessing keras-applications opencv-python rdtsc
