#! /usr/bin/env bash
set -ex

echo "debian" > /etc/hostname
echo root:rootroot | chpasswd


echo deb http://archive.ubuntu.com/ubuntu noble universe >> /etc/apt/sources.list
apt update
apt install -y nodejs libgl1 gfortran ruby php-cli python3-pil libopenblas-dev

apt update
apt install -y openssh-server build-essential gpg wget libblas3 liblapack3 liblapack-dev libblas-dev gfortran libffi-dev python3.12 python3-pip
wget https://github.com/Kitware/CMake/releases/download/v3.22.2/cmake-3.22.2-linux-x86_64.tar.gz -O /opt/cmake-3.22.2-linux-x86_64.tar.gz
pushd /opt
tar xzvf cmake-3.22.2-linux-x86_64.tar.gz
export PATH=$PATH:/opt/cmake-3.22.2-linux-x86_64/bin/
popd
# apt install -y tcpdump build-essential pkg-config python3-setuptools python-dev python3-dev gcc libpq-dev python-pip python3-dev python3-pip python3-venv python3-wheel
MAKEFLAGS="-j54" pip3 install --break-system-packages wheel six scikit-learn flask pillow pyaes chameleon pandas tensorflow torch torchvision minio psutil keras-preprocessing keras-applications opencv-python
mkdir -p /etc/systemd/system/serial-getty@ttyS0.service.d/
cat <<EOF > /etc/systemd/system/serial-getty@ttyS0.service.d/autologin.conf
[Service]
ExecStart=
ExecStart=-/sbin/agetty --autologin root -o '-p -- \\u' --keep-baud 115200,38400,9600 %I $TERM
EOF


cat <<EOF > /etc/systemd/network/eth0.network
[Match]
Name=eth0

[Network]
Address=172.16.0.2/24
Gateway=172.16.0.1
DNS=8.8.8.8
EOF

cat <<EOF > /etc/systemd/system/init-entropy.service
[Unit]
Description=Init entropy
Wants=network-online.target
After=network-online.target
[Service]
Type=simple
User=root
ExecStart=python3 /app/entropy.py
[Install]
WantedBy=multi-user.target
EOF
chmod 644 /etc/systemd/system/init-entropy.service
systemctl enable init-entropy.service

cat <<EOF > /etc/systemd/system/function-daemon.service
[Unit]
Description=Serverless function daemon
Wants=init-entropy.service
After=init-entropy.service
StartLimitIntervalSec=0
[Service]
Type=simple
Restart=always
RestartSec=1
User=root
Environment="FLASK_APP=/app/daemon.py"
ExecStart=python3 -m flask run --host=172.16.0.2
[Install]
WantedBy=multi-user.target
EOF

cat <<EOF >> /etc/sysctl.conf
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
net.ipv6.conf.lo.disable_ipv6 = 1
EOF

# cat <<EOF >> /etc/rsyslog.conf
# *.*    -/dev/shm/syslog
# EOF

cat <<EOF >> /etc/ssh/sshd_config
PermitRootLogin yes
EOF

echo "tmpfs /tmp tmpfs defaults,nosuid,nodev 0 0" >> /etc/fstab

ln -s /dev/shm /usr/tmp

chmod 644 /etc/systemd/system/function-daemon.service
systemctl enable function-daemon.service
systemctl enable systemd-networkd

systemctl disable systemd-timesyncd.service
systemctl disable systemd-update-utmp.service
