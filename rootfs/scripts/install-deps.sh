#!/usr/bin/env bash

set -ex

echo deb http://archive.ubuntu.com/ubuntu noble universe >> /etc/apt/sources.list
apt update
apt install -y python3.12 python3-pip python3.12-venv nodejs openjdk-21-jre-headless npm binutils

python3 -m venv /app/python/

export MAKEFLAGS="-j54"
/app/python/bin/pip3 install --no-cache-dir scikit-learn  tensorflow-cpu minio rdtsc wheel six scikit-learn flask pillow pyaes chameleon pandas tensorflow-cpu minio psutil keras-preprocessing keras-applications opencv-python
/app/python/bin/pip3 install --no-cache-dir torch torchvision --index-url https://download.pytorch.org/whl/cpu

# Find and remove all pycache and tests
find /app/python/ -type d -name "__pycache__" -exec rm -rf {} +
find /app/python/ -type d -name "tests" -exec rm -rf {} +

# Strip debug symbols from binary extensions
find /app/python/ -name "*.so" -exec strip --strip-unneeded {} +

apt-get purge -y binutils
apt-get autoremove -y
apt-get clean
rm -rf /var/lib/apt/lists/*
rm -rf /root/.cache/pip
rm -rf /tmp/*
