.PHONY: all clean debian-base-rootfs.ext4 debian-deps-rootfs.ext4 debian-provisioned-rootfs.ext4 debian-rootfs.ext4

OUTDIR ?= .
export OUTDIR

DIST=debian
DEBIAN_VERSION=noble
GUESTFILES := $(wildcard guest/*)

.INTERMEDIATE: $(OUTDIR)/debian-base-rootfs.ext4 $(OUTDIR)/debian-deps-rootfs.ext4 $(OUTDIR)/debian-provisioned-rootfs.ext4

all:

debian-base-rootfs.ext4: $(OUTDIR)/debian-base-rootfs.ext4
debian-deps-rootfs.ext4: $(OUTDIR)/debian-deps-rootfs.ext4
debian-provisioned-rootfs.ext4: $(OUTDIR)/debian-provisioned-rootfs.ext4
debian-rootfs.ext4: $(OUTDIR)/debian-rootfs.ext4


$(OUTDIR)/debian-base-rootfs.ext4: scripts/create-debian-base-rootfs.sh
	mkdir -p $(OUTDIR)
	scripts/create-debian-base-rootfs.sh $(DEBIAN_VERSION) # install base image

$(OUTDIR)/debian-deps-rootfs.ext4: $(OUTDIR)/debian-base-rootfs.ext4 scripts/install-deps.sh scripts/install-deps-host.sh
	scripts/install-deps-host.sh


$(OUTDIR)/debian-provisioned-rootfs.ext4: $(OUTDIR)/debian-deps-rootfs.ext4 $(OUTDIR)/debian-base-rootfs.ext4 scripts/provision-debian-rootfs.sh scripts/setup-debian-rootfs.sh guest/*
	scripts/provision-debian-rootfs.sh # install dependencies, sys config

$(OUTDIR)/debian-rootfs.ext4: $(OUTDIR)/debian-provisioned-rootfs.ext4 scripts/ready-debian-rootfs.sh $(GUESTFILES)
	scripts/ready-debian-rootfs.sh # install functions

alpinefs.ext4: alpine/alpine-rootfs.sh
	cd alpine/ && ./alpine-rootfs.sh
	mv alpine/alpinefs.ext4 .


clean:
	rm -rf $(OUTDIR)/*.ext4
