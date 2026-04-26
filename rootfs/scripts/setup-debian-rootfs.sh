#! /usr/bin/env bash
set -ex

echo "debian" > /etc/hostname
echo root:rootroot | chpasswd

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

cat <<EOF >> /etc/sysctl.conf
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
net.ipv6.conf.lo.disable_ipv6 = 1
EOF

pushd /app/simple_server
./gradlew build
JAVA_DEPS=$(cat deps.out)
popd

pushd /app/helloworld
./gradlew build
JAVA_DEPS=${JAVA_DEPS}:$(cat deps.out)
popd

pushd /app/image_rotate_s3
./gradlew build
JAVA_DEPS=${JAVA_DEPS}:$(cat deps.out)
popd

pushd /ap/matmul
./gradlew build
JAVA_DEPS=${JAVA_DEPS}:$(cat deps.out)
popd

pkill java
sleep 1

pushd /app/node
npm install minio sharp fs util path
popd


cat <<EOF > /etc/systemd/system/python-daemon.service
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
ExecStart=/app/simple_server.py
[Install]
WantedBy=multi-user.target
EOF

# cat <<EOF > /etc/systemd/system/node-daemon.service
# [Unit]
# Description=Serverless function node daemon
# Wants=init-entropy.service
# After=init-entropy.service
# StartLimitIntervalSec=0
# [Service]
# Type=simple
# Restart=always
# RestartSec=1
# User=root
# WorkingDirectory=/app/node
# Environment="NODE_PATH=/app/node/node_modules:/app/node/node_modules_addon"
# ExecStart=chrt -f 99 /usr/bin/node /app/node/server.js
# [Install]
# WantedBy=multi-user.target
# EOF

# cat <<EOF > /etc/systemd/system/java-daemon.service
# [Unit]
# Description=Java serverless function daemon
# Wants=init-entropy.service
# After=init-entropy.service
# StartLimitIntervalSec=0
# [Service]
# Type=simple
# Restart=always
# RestartSec=1
# User=root
# ExecStart=/usr/bin/java -cp "$JAVA_DEPS" SimpleServer
# [Install]
# WantedBy=multi-user.target
# EOF

# cat <<EOF >> /etc/rsyslog.conf
# *.*    -/dev/shm/syslog
# EOF

cat <<EOF >> /etc/ssh/sshd_config
PermitRootLogin yes
EOF

echo "tmpfs /tmp tmpfs defaults,nosuid,nodev 0 0" >> /etc/fstab

ln -s /dev/shm /usr/tmp

chmod 644 /etc/systemd/system/python-daemon.service
# chmod 644 /etc/systemd/system/node-daemon.service
# chmod 644 /etc/systemd/system/java-daemon.service


systemctl enable python-daemon.service
# systemctl enable node-daemon.service
# systemctl enable java-daemon.service

systemctl enable systemd-networkd
systemctl disable systemd-timesyncd.service
systemctl disable systemd-update-utmp.service
