import sys
import importlib
from flask import Flask, request
app = Flask(__name__)

sys.path.append("/usr/local/lib/python3.7/site-packages")
handlers = {}

@app.route('/invoke', methods=['POST'])
def invoke():
    funcname = request.args['function']

    prog = funcname
    if funcname == "chameleon" or funcname == "pyaes":
        prog = funcname + "1"
    try:
        handler = importlib.import_module(f"{prog}.{prog}").function_handler
        return handler(request.json)
    except Exception as e:
        return str(e)
