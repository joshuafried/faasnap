#! /usr/bin/env bash
set -ex

DEBIAN_VERSION=${1}

sudo umount ./mountpoint || true
sudo rm -rf ./mountpoint
mkdir -p ./mountpoint
dd if=/dev/zero of=.debian-base-rootfs.ext4 bs=2M count=16384
mkfs.ext4 .debian-base-rootfs.ext4

sudo mount .debian-base-rootfs.ext4 mountpoint
sudo debootstrap --arch=amd64 $DEBIAN_VERSION mountpoint http://archive.ubuntu.com/ubuntu/

sudo umount mountpoint
mv .debian-base-rootfs.ext4 debian-base-rootfs.ext4
