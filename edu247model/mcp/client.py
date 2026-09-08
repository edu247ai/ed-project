import json
import sys

def send(cmd):
    sys.stdout.write(json.dumps(cmd) + "\n")
    sys.stdout.flush()

def read():
    return json.loads(sys.stdin.readline())

def list_tools():
    send({"cmd": "list_tools"})
    return read()

def call_tool(name, args):
    send({"cmd": "call_tool", "name": name, "args": args})
    return read()
