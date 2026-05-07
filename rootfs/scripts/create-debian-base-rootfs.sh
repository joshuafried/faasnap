#! /usr/bin/env bash
set -ex

DEBIAN_VERSION=${1}
OUTDIR=${OUTDIR:-.}
MNT=$OUTDIR/mountpoint
TMPOUT=$OUTDIR/.debian-base-rootfs.ext4
OUT=$OUTDIR/debian-base-rootfs.ext4

sudo umount $MNT || true
sudo rm -rf $MNT
mkdir -p $MNT
dd if=/dev/zero of=$TMPOUT bs=2M count=16384
mkfs.ext4 $TMPOUT

sudo mount $TMPOUT $MNT
sudo debootstrap --arch=amd64 $DEBIAN_VERSION $MNT http://archive.ubuntu.com/ubuntu/

sudo umount $MNT
mv $TMPOUT $OUT
