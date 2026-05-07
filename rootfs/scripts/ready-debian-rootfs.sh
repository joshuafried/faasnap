#! /usr/bin/env bash

set -ex

OUTDIR=${OUTDIR:-.}
MNT=$OUTDIR/mountpoint
IN=$OUTDIR/debian-provisioned-rootfs.ext4
OUT=$OUTDIR/debian-rootfs.ext4
TMPOUT=$OUTDIR/.debian-rootfs.ext4

sudo umount $MNT || true
sudo rm -rf $MNT
mkdir -p $MNT
mv $IN $TMPOUT

sudo mount $TMPOUT $MNT

sudo umount $MNT
mv $TMPOUT $OUT
