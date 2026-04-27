#!/usr/bin/env python3
import atexit
import json
import logging
import os
import shutil
import socket
import subprocess
import sys
import tempfile
import time
from collections import defaultdict
from datetime import datetime
from enum import unique
from multiprocessing.pool import Pool
from pathlib import Path
from types import SimpleNamespace

import requests
import urllib3

logging.basicConfig(level=logging.DEBUG)
logging.getLogger("urllib3").setLevel(logging.DEBUG)


sys.path.extend(["./python-client"])
import swagger_client as faasnap
from swagger_client.api.default_api import DefaultApi
from swagger_client.configuration import Configuration

bpf_map = {
    "brq": 'tracepoint:block:block_rq_issue /strncmp("fc_vcpu", comm, 7)==0 || comm =="main"/ {@blockrq[comm] = count(); @bsize[comm] = sum(args->bytes);}',
    "bsize": 'tracepoint:block:block_rq_issue /strncmp("fc_vcpu", comm, 7)==0 || comm =="main"/ {@blockrqsize[comm] = sum(args->bytes)}',
    "_bsize": "tracepoint:block:block_rq_issue {@blockrqsize[comm] = sum(args->bytes)}",
    "pf": 'kprobe:handle_mm_fault /strncmp("fc_vcpu", comm, 7)==0 || comm =="main" || comm=="firecracker"/ {@pf[comm] = count()}',
    "_pf": "kprobe:handle_mm_fault {@pf[comm] = count()}",
    "mpf": 'kretprobe:handle_mm_fault / (retval & 4) == 4 && (strncmp("fc_vcpu", comm, 7)==0 || comm =="main")/ {@majorpf[comm] = count()}',
    "pftime": "kprobe:kvm_mmu_page_fault { @start[tid] = nsecs; } kretprobe:kvm_mmu_page_fault /@start[tid]/ {@n[comm] = count(); $delta = nsecs - @start[tid];  @dist[comm] = hist($delta); @avrg[comm] = avg($delta); delete(@start[tid]); }",
    "vcpublock": "kprobe:kvm_vcpu_block { @start[tid] = nsecs; } kprobe:kvm_vcpu_block /@start[tid]/ {@n[comm] = count(); $delta = nsecs - @start[tid];  @dist[comm] = hist($delta); @avrg[comm] = avg($delta); delete(@start[tid]); }",
    "cache": 'hardware:cache-misses:1000  /strncmp("fc_vcpu", comm, 7)==0/ {@misses[comm] = count()}',
    "mpf-tl": 'BEGIN { @start = nsecs; } kretprobe:handle_mm_fault / @start != 0 && (retval & 4) == 4 && (strncmp("fc_vcpu", comm, 7)==0 ) / { printf("%d\\n", (nsecs - @start) / 1000000); }',
}

PAUSE = None
TESTID = None
RESULT_DIR = None
BPF = None


def cleanup():
    subprocess.Popen(["pkill", "main"])
    subprocess.Popen(["pkill", "faasnap-fc"])


atexit.register(cleanup)


def addNetwork(client: DefaultApi, idx: int):
    ns = "fc%d" % idx
    guest_mac = "AA:FC:00:00:00:01"  # fixed MAC
    guest_addr = "172.16.0.2"  # fixed guest IP
    unique_addr = "192.168.0.%d" % (idx + 2)
    client.net_ifaces_namespace_put(
        namespace=ns,
        interface={
            "host_dev_name": "vmtap0",
            "iface_id": "eth0",
            "guest_mac": guest_mac,
            "guest_addr": guest_addr,
            "unique_addr": unique_addr,
        },
    )


clients = {}


def prepareVanilla(params, client: DefaultApi, setting, func, func_param, par_snap):
    all_snaps = []
    _set_lang_marker(func.lang, params.faasnap.images.debian)
    vm = client.vms_post(vm={"func_name": func.name, "namespace": "fc%d" % 1})
    time.sleep(5)
    invoc = faasnap.Invocation(
        func_name=func.name,
        vm_id=vm.vm_id,
        params=func_param,
        mincore=-1,
        enable_reap=False,
    )
    ret = client.invocations_post(invocation=invoc)
    print("prepare invoc ret:", ret)
    base = faasnap.Snapshot(
        vm_id=vm.vm_id,
        snapshot_type="Full",
        snapshot_path=params.test_dir + "/Full.snapshot",
        mem_file_path=params.test_dir + "/Full.memfile",
        version="0.23.0",
        **vars(setting.record_regions),
    )
    base_snap = client.snapshots_post(snapshot=base)
    all_snaps.append(base_snap)
    client.vms_vm_id_delete(vm_id=vm.vm_id)
    time.sleep(2)
    for i in range(par_snap - 1):
        all_snaps.append(
            client.snapshots_put(
                base_snap.ss_id, "%s/Full.memfile.%d" % (params.test_dir, i)
            )
        )
    for snap in all_snaps:
        client.snapshots_ss_id_patch(
            ss_id=snap.ss_id, state=vars(setting.patch_state)
        )  # drop cache
    time.sleep(1)
    return [snap.ss_id for snap in all_snaps]


def _set_lang_marker(lang, rootfs_path):
    with tempfile.TemporaryDirectory() as mnt:
        subprocess.run(["mount", rootfs_path, mnt], check=True)
        with open(os.path.join(mnt, "lang-select"), "w") as f:
            f.write(lang)
        subprocess.run(["umount", mnt], check=True)


def prepareMincore(
    params,
    client: DefaultApi,
    setting,
    func,
    func_param,
    par_snap,
    snapshot_only=False,
    restore_only=False,
):
    all_snaps = []
    _set_lang_marker(func.lang, params.faasnap.images.debian)
    time.sleep(1)
    vm = client.vms_post(vm={"func_name": func.name, "namespace": "fc%d" % 1})
    time.sleep(10)

    snapshot_path = params.test_dir + "/Full.snapshot"
    mem_file_path = params.test_dir + "/Full.memfile"

    if snapshot_only or restore_only:
        os.makedirs(params.test_dir + "/faasnap", exist_ok=True)
        snapshot_path = params.test_dir + "/faasnap/Full.snapshot"
        mem_file_path = params.test_dir + "/faasnap/Full.memfile"

    if not restore_only:
        base_snap = client.snapshots_post(
            snapshot=faasnap.Snapshot(
                vm_id=vm.vm_id,
                snapshot_type="Full",
                snapshot_path=snapshot_path,
                mem_file_path=mem_file_path,
                version="0.23.0",
            )
        )
        client.vms_vm_id_delete(vm_id=vm.vm_id)
        client.snapshots_ss_id_patch(
            ss_id=base_snap.ss_id, state=vars(setting.patch_base_state)
        )  # drop cache

    if snapshot_only:
        print("Snapshot saved to", snapshot_path)
        cleanup()
        return []

    if setting.mincore_size > 0:
        mincore = -1
    else:
        mincore = 100
    time.sleep(1)
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        ss_id=base_snap.ss_id,
        params=func_param,
        mincore=mincore,
        mincore_size=setting.mincore_size,
        enable_reap=False,
        namespace="fc%d" % 1,
        use_mem_file=True,
    )
    ret = client.invocations_post(invocation=invoc).to_dict()
    newVmID = ret["vm_id"]
    print("prepare invoc ret:", ret)

    # disable page zeroing
    func_param = json.loads(func_param)
    func_param["disable_sanpage"] = True
    func_param = json.dumps(func_param)

    ret = client.invocations_post(
        invocation=faasnap.Invocation(
            func_lang=func.lang,
            func_name="run",
            vm_id=newVmID,
            params=func_param,
            mincore=-1,
            enable_reap=False,
        )
    )  # disable sanitizing
    print("disable sanitizing invoc ret: ", ret)
    warm_snap = client.snapshots_post(
        snapshot=faasnap.Snapshot(
            vm_id=newVmID,
            snapshot_type="Full",
            snapshot_path=params.test_dir + "/Warm.snapshot",
            mem_file_path=params.test_dir + "/Warm.memfile",
            version="0.23.0",
            **vars(setting.record_regions),
        )
    )
    all_snaps.append(warm_snap)
    client.vms_vm_id_delete(vm_id=newVmID)
    time.sleep(2)
    client.snapshots_ss_id_mincore_put(
        ss_id=warm_snap.ss_id, source=base_snap.ss_id
    )  # carry over mincore to new snapshot
    client.snapshots_ss_id_mincore_patch(
        ss_id=warm_snap.ss_id, state=vars(setting.patch_mincore)
    )
    for i in range(par_snap - 1):
        all_snaps.append(
            client.snapshots_put(
                warm_snap.ss_id, "%s/Full.memfile.%d" % (params.test_dir, i)
            )
        )
    client.snapshots_ss_id_patch(
        ss_id=base_snap.ss_id, state=vars(setting.patch_base_state)
    )  # drop cache
    for snap in all_snaps:
        client.snapshots_ss_id_patch(
            ss_id=snap.ss_id, state=vars(setting.patch_state)
        )  # drop cache
        client.snapshots_ss_id_mincore_patch(
            ss_id=warm_snap.ss_id, state={"drop_ws_cache": True}
        )
    # input("Press Enter to start finish invocation...")
    time.sleep(1)

    return [snap.ss_id for snap in all_snaps]


def prepareReap(
    params,
    client: DefaultApi,
    setting,
    func,
    func_param,
    idx,
    delay_record=False,
    snapshot_only=False,
    restore_only=False,
):
    _set_lang_marker(func.lang, params.faasnap.images.debian)
    vm = client.vms_post(vm={"func_name": func.name, "namespace": "fc%d" % idx})
    time.sleep(10)
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        vm_id=vm.vm_id,
        params=func_param,
        mincore=-1,
        enable_reap=False,
    )
    ret = client.invocations_post(invocation=invoc)
    print("1st prepare invoc ret:", ret)

    snapshot_path = params.test_dir + "/Full.snapshot" + str(idx)
    mem_file_path = params.test_dir + "/Full.memfile" + str(idx)

    if snapshot_only or restore_only:
        os.makedirs(params.test_dir + "/reap_" + func.name, exist_ok=True)
        snapshot_file = params.test_dir + f"/reap_{func.name}/Full.snapshot" + str(idx)
        mem_file_path = params.test_dir + f"/reap_{func.name}/Full.memfile" + str(idx)

    if restore_only:
        snapshot = faasnap.Snapshot(vm.vm_id)
        snapshot.mem_file_path = mem_file_path
        snapshot.snapshot_path = snapshot_path
        snapshot.snapshot_type = "Full"
        snapshot.version = "0.23.0"
        snapshot.ss_id = f"reap_{func.name}"

        base_snap = client.snapshots_post(snapshot=snapshot)
        print(base_snap)
    else:
        base = faasnap.Snapshot(
            vm_id=vm.vm_id,
            snapshot_type="Full",
            snapshot_path=snapshot_path,
            mem_file_path=mem_file_path,
            version="0.23.0",
        )
        base_snap = client.snapshots_post(snapshot=base)

    client.vms_vm_id_delete(vm_id=vm.vm_id)
    time.sleep(1)
    client.snapshots_ss_id_patch(
        ss_id=base_snap.ss_id, state=vars(setting.patch_state)
    )  # drop cache

    if snapshot_only:
        print("saved snapshot")
        return []

    time.sleep(1)
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        ss_id=base_snap.ss_id,
        params=func_param,
        mincore=-1,
        enable_reap=True,
        ws_file_direct_io=True,
        namespace="fc%d" % 1,
    )
    ret = client.invocations_post(invocation=invoc).to_dict()
    print("2nd prepare invoc ret:", ret)
    time.sleep(1)
    client.vms_vm_id_delete(vm_id=ret["vm_id"])
    time.sleep(2)
    client.snapshots_ss_id_patch(
        ss_id=base_snap.ss_id, state=vars(setting.patch_state)
    )  # drop cache
    client.snapshots_ss_id_reap_patch(
        ss_id=base_snap.ss_id, cache=False
    )  # drop reap cache
    time.sleep(1)
    return [base_snap.ss_id]


def prepareEmuMincore(params, client: DefaultApi, setting, func, func_param):
    _set_lang_marker(func.lang, params.faasnap.images.debian)
    vm = client.vms_post(vm={"func_name": func.name, "namespace": "fc%d" % 1})
    time.sleep(5)
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        vm_id=vm.vm_id,
        params=func_param,
        mincore=-1,
        enable_reap=False,
    )
    ret = client.invocations_post(invocation=invoc)
    print("1st prepare invoc ret:", ret)
    snapshot = client.snapshots_post(
        snapshot=faasnap.Snapshot(
            vm_id=vm.vm_id,
            snapshot_type="Full",
            snapshot_path=params.test_dir + "/Full.snapshot",
            mem_file_path=params.test_dir + "/Full.memfile",
            version="0.23.0",
            **vars(setting.record_regions),
        )
    )
    client.vms_vm_id_delete(vm_id=vm.vm_id)
    time.sleep(1)
    client.snapshots_ss_id_patch(
        ss_id=snapshot.ss_id, state=vars(setting.patch_state)
    )  # drop cache
    time.sleep(1)
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        ss_id=snapshot.ss_id,
        params=func_param,
        mincore=-1,
        enable_reap=True,
        ws_file_direct_io=True,
        namespace="fc%d" % 1,
    )  # get emulated mincore
    ret = client.invocations_post(invocation=invoc)
    print("2nd prepare invoc ret:", ret)
    time.sleep(1)
    client.vms_vm_id_delete(vm_id=ret["vm_id"])
    time.sleep(2)
    client.snapshots_ss_id_reap_patch(
        ss_id=snapshot.ss_id, cache=False
    )  # drop reap cache
    client.snapshots_ss_id_mincore_patch(
        ss_id=snapshot.ss_id, state=vars(setting.patch_mincore)
    )
    client.snapshots_ss_id_patch(
        ss_id=snapshot.ss_id, state=vars(setting.patch_state)
    )  # drop cache
    time.sleep(1)
    return [snapshot.ss_id]


def invoke(args):
    (
        params,
        setting,
        func,
        func_param,
        idx,
        ss_id,
        par,
        par_snap,
        record_input,
        test_input,
    ) = args

    setting_name = setting.name
    if "faasnap" in setting_name and not setting.invocation.prefetch:
        setting_name += "_no_prefetch"

    if par > 1 or par_snap > 1:
        runId = "%s_%s_%d_%d" % (setting_name, func.id, par, par_snap)
    else:
        runId = "%s_%s_%d%d" % (setting_name, func.id, record_input, test_input)
    bpfpipe = None
    time.sleep(1)
    mcstate = None
    if setting.invoke_steps == "vanilla":
        invoc = faasnap.Invocation(
            func_lang=func.lang,
            func_name=func.name,
            ss_id=ss_id,
            params=func_param,
            mincore=-1,
            enable_reap=False,
            namespace="fc%d" % idx,
            **vars(setting.invocation),
        )
    elif setting.invoke_steps == "mincore":
        mcstate = clients[idx].snapshots_ss_id_mincore_get(ss_id=ss_id).to_dict()
        invoc = faasnap.Invocation(
            func_lang=func.lang,
            func_name=func.name,
            ss_id=ss_id,
            params=func_param,
            mincore=-1,
            load_mincore=[n + 1 for n in range(mcstate["nlayers"])],
            enable_reap=False,
            namespace="fc%d" % idx,
            **vars(setting.invocation),
        )
    elif setting.invoke_steps == "reap":
        invoc = faasnap.Invocation(
            func_lang=func.lang,
            func_name=func.name,
            ss_id=ss_id,
            params=func_param,
            mincore=-1,
            enable_reap=True,
            ws_single_read=True,
            namespace="fc%d" % idx,
        )
    else:
        print("invoke steps undefined")
        return
    if BPF:
        program = bpf_map[BPF]
        bpffile = (
            open("%s/%s/bpftrace" % (RESULT_DIR, TESTID), "a+")
            if RESULT_DIR
            else open("/tmp/bpftrace", "a+")
        )
        print("==== %s ====" % runId, file=bpffile, flush=True)
        bpfpipe = subprocess.Popen(
            ["bpftrace", "-e", program],
            cwd="/tmp/",
            stdout=bpffile,
            stderr=subprocess.STDOUT,
        )
        time.sleep(3)

    ret = clients[idx].invocations_post(invocation=invoc).to_dict()
    if bpfpipe:
        bpfpipe.terminate()
        bpfpipe.wait()
    clients[idx].vms_vm_id_delete(vm_id=ret["vm_id"])
    trace_id = ret["trace_id"]
    print("invoke", runId, "ret:", ret)
    time.sleep(2)
    if RESULT_DIR:
        directory = "%s/%s/%s" % (RESULT_DIR, TESTID, runId)
        os.makedirs(directory, exist_ok=True)
        with open("%s/%s.json" % (directory, trace_id), "w+") as f:
            resp = requests.get("%s/%s" % (params.trace_api, trace_id))
            json.dump(resp.json(), f)
        with open(f"{directory}/ret.json", "w+") as f:
            json.dump(ret, f)
        with open("%s/%s-mcstate.json" % (directory, trace_id), "w+") as f:
            json.dump([mcstate], f)
        vm_id = ret["vm_id"]
        shutil.copy(f"{params.test_dir}/{vm_id}/log", directory)


def run_snap(
    params,
    setting,
    par,
    par_snap,
    func,
    record_input,
    test_input,
    delay_record=False,
    snapshot_only=False,
    restore_only=False,
):
    if par_snap > 1:
        assert par == par_snap
    client: DefaultApi
    global clients
    # start faasnap
    snappipe = subprocess.Popen(
        ["./main", "--port=8080", "--host=0.0.0.0"],
        cwd=params.home_dir,
        stdout=open("%s/%s/stdout" % (RESULT_DIR, TESTID), "a+")
        if RESULT_DIR
        else open("/tmp/faasnap-stdout", "a+"),
        stderr=subprocess.STDOUT,
    )
    time.sleep(10)
    # set up
    keepalive_pool = urllib3.PoolManager(
        socket_options=[
            (socket.SOL_SOCKET, socket.SO_KEEPALIVE, 1),
            (socket.IPPROTO_TCP, socket.TCP_KEEPIDLE, 1),
            (socket.IPPROTO_TCP, socket.TCP_KEEPINTVL, 3),
            (socket.IPPROTO_TCP, socket.TCP_KEEPCNT, 5),
        ]
    )

    api_client = faasnap.ApiClient(conf)
    api_client.rest_client.pool_manager = keepalive_pool

    for idx in range(1, 1 + par):
        clients[idx] = faasnap.DefaultApi(api_client)
        addNetwork(clients[idx], idx)
    client = clients[1]
    client.functions_post(
        function=faasnap.Function(
            func_lang=func.lang,
            func_name=func.name,
            image=func.image,
            kernel=setting.kernel,
            vcpu=params.vcpu,
            mem_size=params.memSize,
        )
    )

    params0 = func.params[record_input]
    params1 = func.params[test_input]
    if setting.prepare_steps == "vanilla":
        ssIds = prepareVanilla(
            params, client, setting, func, params0, par_snap=par_snap
        )
    elif setting.prepare_steps == "mincore":
        ssIds = prepareMincore(
            params,
            client,
            setting,
            func,
            params0,
            par_snap=par_snap,
            snapshot_only=snapshot_only,
            restore_only=restore_only,
        )
    elif setting.prepare_steps == "reap":
        ssIds = []
        for idx in range(par_snap):
            ssIds += prepareReap(
                params,
                client,
                setting,
                func,
                params0,
                idx=idx + 1,
                delay_record=delay_record,
                snapshot_only=snapshot_only,
                restore_only=restore_only,
            )
    elif setting.prepare_steps == "emumincore":
        ssIds = prepareEmuMincore(params, client, setting, func, params0)

    time.sleep(1)
    if PAUSE:
        input("Press Enter to start...")
    with Pool(par) as p:
        if len(ssIds) > 1:
            vector = [
                (
                    params,
                    setting,
                    func,
                    params1,
                    idx,
                    ssIds[idx - 1],
                    par,
                    par_snap,
                    record_input,
                    test_input,
                )
                for idx in range(1, 1 + par)
            ]
        else:
            vector = [
                (
                    params,
                    setting,
                    func,
                    params1,
                    idx,
                    ssIds[0],
                    par,
                    par_snap,
                    record_input,
                    test_input,
                )
                for idx in range(1, 1 + par)
            ]
        p.map(invoke, vector)

    # input("Press Enter to finish...")
    snappipe.terminate()
    snappipe.wait()
    time.sleep(1)


def invoke_warm(args):
    client: DefaultApi
    params, setting, func, func_param, idx, vm_id = args
    client = clients[idx]
    runId = "%s_%s" % (setting.name, func.id)
    time.sleep(1)
    mcstate = None
    invoc = faasnap.Invocation(
        func_lang=func.lang,
        func_name=func.name,
        vm_id=vm_id,
        params=func_param,
        mincore=-1,
        enable_reap=False,
    )
    if BPF:
        program = bpf_map[BPF]
        bpffile = (
            open("%s/%s/bpftrace" % (RESULT_DIR, TESTID), "a+")
            if RESULT_DIR
            else open("/tmp/bpftrace", "a+")
        )
        print("==== %s ====" % runId, file=bpffile, flush=True)
        bpfpipe = subprocess.Popen(
            ["bpftrace", "-e", program],
            cwd="/tmp/",
            stdout=bpffile,
            stderr=subprocess.STDOUT,
        )
        time.sleep(3)
    ret = client.invocations_post(invocation=invoc).to_dict()
    if BPF:
        bpfpipe.terminate()
        bpfpipe.wait()
    print("2nd invoc ret:", ret)
    trace_id = ret["trace_id"]
    client.vms_vm_id_delete(vm_id=vm_id)
    time.sleep(2)
    if RESULT_DIR:
        directory = "%s/%s/%s" % (RESULT_DIR, TESTID, runId)
        os.makedirs(directory, exist_ok=True)
        with open("%s/%s.json" % (directory, trace_id), "w+") as f:
            resp = requests.get("%s/%s" % (params.trace_api, trace_id))
            json.dump(resp.json(), f)


def run_warm(params, setting, par, par_snap, func, record_input, test_input):
    client: DefaultApi
    snappipe = subprocess.Popen(
        ["./main", "--port=8080", "--host=0.0.0.0"],
        cwd=params.home_dir,
        stdout=open("%s/%s/stdout" % (RESULT_DIR, TESTID), "a+")
        if RESULT_DIR
        else open("/tmp/faasnap-stdout", "a+"),
        stderr=subprocess.STDOUT,
    )
    time.sleep(2)
    # set up
    for idx in range(1, 1 + par):
        clients[idx] = faasnap.DefaultApi(faasnap.ApiClient(conf))
        addNetwork(clients[idx], idx)
    client = clients[1]
    client.functions_post(
        function=faasnap.Function(
            func_lang=func.lang,
            func_name=func.name,
            image=func.image,
            kernel=setting.kernel,
            vcpu=params.vcpu,
        )
    )

    params0 = func.params[record_input]
    params1 = func.params[test_input]

    vms = {}
    _set_lang_marker(func.lang, params.faasnap.images.debian)
    for idx in range(1, 1 + par):
        vms[idx] = clients[idx].vms_post(
            vm={"func_name": func.name, "namespace": "fc%d" % idx}
        )
    time.sleep(5)

    for idx in range(1, 1 + par):
        invoc = faasnap.Invocation(
            func_lang=func.lang,
            func_name=func.name,
            vm_id=vms[idx].vm_id,
            params=params0,
            mincore=-1,
            enable_reap=False,
        )
        ret = clients[idx].invocations_post(invocation=invoc)
        print("1st invoc ret:", ret)
    time.sleep(1)

    if PAUSE:
        input("Press Enter to start...")
    with Pool(par) as p:
        vector = [
            (params, setting, func, params1, idx, vms[idx].vm_id)
            for idx in range(1, 1 + par)
        ]
        p.map(invoke_warm, vector)

    snappipe.terminate()
    snappipe.wait()
    time.sleep(5)


def run(
    params,
    setting,
    func,
    par,
    par_snap,
    repeat,
    record_input,
    test_input,
    delay_record=False,
    snapshot_only=False,
    restore_only=False,
):
    setting_name = setting.name
    if "faasnap" in setting_name and not setting.invocation.prefetch:
        setting_name += "_no_prefetch"
    for r in range(repeat):
        print("\n=========%s %s: %d=========\n" % (setting_name, func.id, r))
        if setting.name == "warm":
            run_warm(params, setting, par, par_snap, func, record_input, test_input)
        else:
            run_snap(
                params,
                setting,
                par,
                par_snap,
                func,
                record_input,
                test_input,
                delay_record=delay_record,
                snapshot_only=snapshot_only,
                restore_only=restore_only,
            )


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: %s <test.json>" % sys.argv[0])
        exit(1)
    PAUSE = os.environ.get("PAUSE", None)
    TESTID = os.environ.get("TESTID", datetime.now().strftime("%Y-%m-%dT%H-%M-%S"))
    print("TESTID:", TESTID)
    RESULT_DIR = os.environ.get("RESULT_DIR", None)
    if not RESULT_DIR:
        print("no RESULT_DIR set, will not save results")
    else:
        full_path = f"{RESULT_DIR}/{TESTID}"
        os.makedirs(full_path, mode=0o777, exist_ok=True)
        os.system(f"ln -sfn {full_path} {RESULT_DIR}/faasnap.recent")
    BPF = os.environ.get("BPF", None)
    with open(sys.argv[1], "r") as f:
        params = json.load(f, object_hook=lambda d: SimpleNamespace(**d))
    conf = Configuration()
    conf.timeout = 10000
    conf.host = params.host

    params.settings.faasnap.patch_mincore.to_ws_file = params.test_dir + "/wsfile"

    if RESULT_DIR:
        n = 1
        while True:
            p = Path("%s/%s/tests-%d.json" % (RESULT_DIR, TESTID, n))
            if not p.exists():
                break
            n += 1
        with p.open("w") as f:
            json.dump(
                params, f, default=lambda o: o.__dict__, sort_keys=False, indent=4
            )
    with open("/etc/faasnap.json", "w") as f:
        json.dump(
            params.faasnap, f, default=lambda o: o.__dict__, sort_keys=False, indent=4
        )

    print("test_dir:", params.test_dir)
    print("repeat:", params.repeat)
    print("parallelism:", params.parallelism)
    print("par_snapshots:", params.par_snapshots)
    print("kernels:", params.faasnap.kernels)
    print("vcpu:", params.vcpu)
    print("memSize:", params.memSize)
    print("record input:", params.record_input)
    print("test input:", params.test_input)
    print("delay record:", params.delay_record)

    print(params.function)

    for func in params.function:
        for setting in params.setting:
            for par, par_snap in zip(params.parallelism, params.par_snapshots):
                for record_input in params.record_input:
                    for test_input in params.test_input:
                        run(
                            params,
                            setting=vars(params.settings)[setting],
                            func=vars(params.functions)[func],
                            par=par,
                            par_snap=par_snap,
                            repeat=params.repeat,
                            record_input=record_input,
                            test_input=test_input,
                            delay_record=params.delay_record,
                            snapshot_only=params.snapshot_only,
                            restore_only=params.restore_only,
                        )
                        time.sleep(5)
