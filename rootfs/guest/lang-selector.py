import os
import socket
import subprocess
import time

MARKER = "/lang-select"

SERVICES = {
    "python": ("python-daemon.service", 5000),
    "node": ("node-daemon.service", 5003),
    "java": ("java-daemon.service", 5001),
}


def wait_for_port(port, timeout=30):
    deadline = time.time() + timeout
    while time.time() < deadline:
        try:
            s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            s.connect(("127.0.0.1", port))
            s.close()
            return True
        except OSError:
            time.sleep(0.05)
    return False


def run():
    if not os.path.exists(MARKER):
        return

    with open(MARKER) as f:
        lang = f.read().strip()

    os.remove(MARKER)

    service, port = SERVICES[lang]
    subprocess.run(["systemctl", "start", service], check=True)
    wait_for_port(port)


if __name__ == "__main__":
    run()
