#! /usr/bin/env bash
set -ex

OUTDIR=${OUTDIR:-.}
MNT=$OUTDIR/mountpoint
IN=$OUTDIR/debian-deps-rootfs.ext4
OUT=$OUTDIR/debian-provisioned-rootfs.ext4
TMPOUT=$OUTDIR/.debian-provisioned-rootfs.ext4

sudo umount $MNT/dev || true
sudo umount $MNT/sys || true
sudo umount $MNT/proc || true
sudo umount $MNT || true
sudo rm -rf $MNT
mkdir -p $MNT
mv $IN $TMPOUT

sudo mount $TMPOUT $MNT

sudo mkdir -p $MNT/app
sudo cp -r guest/* $MNT/app/

sudo mount --bind /dev $MNT/dev
sudo mount --bind /proc $MNT/proc
sudo mount --bind /sys $MNT/sys

sudo cp scripts/setup-debian-rootfs.sh $MNT/
sudo chroot $MNT /bin/bash /setup-debian-rootfs.sh

sudo umount $MNT/dev
sudo umount $MNT/sys
sudo umount $MNT/proc
sudo umount $MNT
mv $TMPOUT $OUT
