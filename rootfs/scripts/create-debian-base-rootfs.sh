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
truncate -s 32G $TMPOUT
mkfs.ext4 $TMPOUT

sudo mount -o loop,discard $TMPOUT $MNT
sudo debootstrap --arch=amd64 $DEBIAN_VERSION $MNT http://archive.ubuntu.com/ubuntu/

sudo umount $MNT
mv $TMPOUT $OUT
