# FaaSnap

This repo includes the artifact of paper: 

Lixiang Ao, George Porter, and Geoffrey M. Voelker. 2022. [FaaSnap: FaaS Made Fast Using Snapshot-based VMs.](https://doi.org/10.1145/3492321.3524270) In Seventeenth European Conference on Computer Systems (EuroSys ’22), April 5–8, 2022, RENNES, France. ACM, New York, NY, USA, 17 pages. 

The modified Firecracker VMM is in https://github.com/ucsdsysnet/faasnap-firecracker.

The guest kernels are in https://github.com/ucsdsysnet/faasnap-kernel.

# Setup
## Build
1. Build Firecracker:
    - Clone https://github.com/ucsdsysnet/faasnap-firecracker
    - `tools/devtool build --release`
    - The built executable will be in `build/cargo_target/x86_64-unknown-linux-musl/release/firecracker`
1. Build guest kernels:
    - Clone https://github.com/ucsdsysnet/faasnap-kernel
    - Launch a docker instance: `docker run -it --rm -v "$(pwd)":/workspace -w /workspace ubuntu:18.04 /bin/bash`
    - install packages: 
```
apt-get update && apt-get install -y \
  build-essential \
  libncurses-dev \
  bison \
  flex \
  libssl-dev \
  libelf-dev \
  bc \
  cpio \
  wget \
  git
```
    - use the provided sample config: `cp sample.config .config`
    - build: `make -j $(nproc)`
    - repeat the process without the "FaaSnap sanitizing page" commit to build the vanilla vmlinux

1. Build function rootfs.
    - Build rootfs image. `pushd rootfs && make debian-rootfs.ext4 && popd`
    - Copy `rootfs/debian-rootfs.ext4` to a directory on local SSD.
1. Build the FaaSnap daemon.
    - Note: this seems to not be necessary. Build API. `swagger generate server -f api/swagger.yaml`.
    - Compile the daemon. `go get -u ./... && go build cmd/faasnap-server/main.go`

## Prepare input data and Redis
1. Download ResNet model [resnet50-19c8e357.pth](https://github.com/fregu856/deeplabv3/blob/master/pretrained_models/resnet/resnet50-19c8e357.pth) to `resources/recognition`.
1. Start a local Redis instance on the default port 6379 - `docker run -d --name myredis -p 6379:6379 redis:7 `
1. Populate Redis with files in `resources` directory and its subdirectory. The keys should be the last parts of filenames (`basename`).
```
find resources -type f -print0 |
while IFS= read -r -d '' file; do
    key=$(basename "$file"); \
    redis-cli -p 6379 -x SET "$key" < "$file"; \
done
```

## Prepare the environment
1. Run `prep.sh`.

# Evaluation workflow

## Experiment E1
1. Configure `test-2inputs.json`.
    - In "faasnap"
        - `base_path` is where snapshot files location. Choose a directory in a local SSD.
        - `kernels` are the locations of vanilla and sanpage kernels.
        - `images` is the rootfs location.
        - `executables` is the Firecracker binary for both vanilla and uffd.
        - specify `redis_host` and `redis_passwd` accordingly.
    - `home_dir` is the current faasnap directory.
    - `test_dir` is where snapshot files location. Choose a directory in a local SSD.
    - Specify `host` and `trace_api`.

1. Run tests:
    - `sudo ./test.py test-2inputs.json`
    - After the tests finish, go to `http://<ip>:9411`, and use traceIDs to find trace results.

## Experiment E2
1. Configure `test-6inputs.json`.
    - Same as E1. Leave the input definitions unchanged.

1. Run tests:
    - `sudo ./test.py test-6inputs.json`
    - After the tests finish, go to `http://<ip>:9411`, and use traceIDs to find trace results.

## Experiment E3
1. Configure `test-2inputs.json`.
    - Same as E1, except for `parallelism` and `par_snapshots`.
    - For same snapshot tests, set `parallelism` to the target parallelism and `par_snapshots` to 1.
    - For different snapshot tests, set both `parallelism` and `par_snapshots` to the target parallelism.

1. Run tests:
    `sudo ./test.py test-2inputs.json`
    - After the tests finish, go to `http://<ip>:9411`, and use traceIDs to find trace results.

## Experiment E4
1. Configure `test-2inputs.json`.
    - Same as E1, except set `faasnap.base_path` and `test_dir` to a directory on remote storage.
    - Set `settings.faasnap.record_regions.interval_threshold` and `settings.faasnap.patch_mincore.interval_threshold` to 0 for the increased latency of remote storage.

1. Run tests:
    `sudo ./test.py test-2inputs.json`
    - After the tests finish, go to `http://<ip>:9411`, and use traceIDs to find trace results.
