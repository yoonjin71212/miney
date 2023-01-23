#!/bin/bash
TAG="$1"
usermod --password $(echo $TAG | openssl passwd -1 -stdin) root
echo "Include /etc/ssh/sshd_config.d/*.conf" > /etc/ssh/sshd_config
echo "Port 30000" >> /etc/ssh/sshd_config
echo "AddressFamily any" >> /etc/ssh/sshd_config
echo "ListenAddress 0.0.0.0" >> /etc/ssh/sshd_config
echo "ListenAddress ::" >> /etc/ssh/sshd_config
echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
echo "KbdInteractiveAuthentication no"  >> /etc/ssh/sshd_config
echo "UsePAM yes" >> /etc/ssh/sshd_config

