#! /usr/bin/env bash
set -ex

IN=debian-deps-rootfs.ext4
OUT=debian-provisioned-rootfs.ext4
TMPOUT=.$OUT

sudo umount ./mountpoint/dev || true
sudo umount ./mountpoint/sys || true
sudo umount ./mountpoint/proc || true
sudo umount ./mountpoint || true
sudo rm -rf ./mountpoint
mkdir -p ./mountpoint
cp $IN $TMPOUT

sudo mount $TMPOUT mountpoint

sudo mkdir -p mountpoint/app
sudo cp -r guest/* mountpoint/app/

sudo mount --bind /dev ./mountpoint/dev
sudo mount --bind /proc ./mountpoint/proc
sudo mount --bind /sys ./mountpoint/sys

sudo cp scripts/setup-debian-rootfs.sh mountpoint/
sudo chroot mountpoint /bin/bash /setup-debian-rootfs.sh

sudo umount ./mountpoint/dev
sudo umount ./mountpoint/sys
sudo umount ./mountpoint/proc
sudo umount mountpoint
mv $TMPOUT $OUT
