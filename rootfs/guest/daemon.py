import importlib
import os
import sys

from flask import Flask, request

app = Flask(__name__)

sys.path.append("/usr/local/lib/python3.7/site-packages")
handlers = {}

# use SCHED_FIFO to help mitigate scheduling delay on restore
param = os.sched_param(os.sched_get_priority_max(os.SCHED_FIFO))
os.sched_setscheduler(0, os.SCHED_FIFO, param)


@app.route("/invoke", methods=["POST"])
def invoke():
    funcname = request.args["function"]

    prog = funcname
    if funcname == "chameleon" or funcname == "pyaes":
        prog = funcname + "1"
    try:
        handler = importlib.import_module(f"{prog}.{prog}").function_handler
        return handler(request.json)
    except Exception as e:
        return str(e)
