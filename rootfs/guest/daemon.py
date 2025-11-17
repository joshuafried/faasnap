import sys
import importlib
import subprocess
import os
import time
import psutil
import rdtsc
import socket
import threading

from flask import Flask, request
app = Flask(__name__)

sys.path.append("/usr/local/lib/python3.7/site-packages")
handlers = {}

sched_ret = None
# param = os.sched_param(os.sched_get_priority_max(os.SCHED_FIFO))
# sched_ret = os.sched_setscheduler(0, os.SCHED_FIFO, param)
msg = None
def listen():
    global msg
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    msg = "TEST"
    try:
        msg = "MADE IT TO TRY"
        sock.bind(('172.16.0.2', 5001))
        msg = "MADE IT PAST BIND"
    except Exception as e:
        msg = str(e)
    sock.listen()
    msg = "MADE IT PAST LISTEN"
    while True:
        conn, addr = sock.accept()
        msg = "MADE IT PAST ACCEPT"
        data = conn.recv(1024)
        msg = data.decode()

listener = threading.Thread(target=listen, daemon=True)
listener.start()

def run(cmd, quiet=False):
    if not quiet:
        print(cmd)
        sys.stdout.flush()

    return subprocess.check_output(cmd, shell=True)

# run("mount -t tracefs nodev /sys/kernel/debug/tracing")
# run("echo 1 > /sys/kernel/debug/tracing/events/sched/sched_switch/enable")

def get_process_tree_pids(root_pid):
    """
    Return a list of all PIDs in the process tree rooted at `root_pid`.
    """
    try:
        root = psutil.Process(root_pid)
        descendants = root.children(recursive=True)
        return [root_pid] + [p.pid for p in descendants]
    except psutil.NoSuchProcess:
        return []

@app.route('/getpid', methods=['POST'])
def getpid():
    run("dmesg -C")
    pid = os.getpid()
    return str(pid)

@app.route('/invoke', methods=['POST'])
def invoke():
    begin_tsc = rdtsc.get_cycles()
    funcname = request.args['function']
    now = time.time()

    prog = funcname
    if funcname == "chameleon" or funcname == "pyaes":
        prog = funcname + "1"
    try:
        handler = importlib.import_module(f"{prog}.{prog}").function_handler
        ret = handler(request.json)
        pid = os.getpid()
        all_pids = get_process_tree_pids(pid)
        schedstat = run(f"cat /proc/{pid}/schedstat").decode('utf-8')
        # dmesg = run("dmesg | grep BCWH").decode('utf-8')
        # run("echo 0 | sudo tee /sys/kernel/debug/tracing/events/sched/sched_switch/enable")

        # with open('/sys/kernel/debug/tracing/trace', 'r') as f:
        #     dmesg = f.readlines()


        # dmesg = run("sudo cat /sys/kernel/debug/tracing/trace")
        # dmesg = None
        tsc = rdtsc.get_cycles()
        return f"begin_tsc={begin_tsc}, latency={time.time() - now}, pids={all_pids}, schedstat={schedstat}, sched_ret={sched_ret}, tsc={tsc}, msg={msg}"
    except Exception as e:
        return str(e)
