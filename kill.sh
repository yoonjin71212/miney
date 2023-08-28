#!/bin/bash
kill -9 $(pgrep server)
kill -9 $(pgrep server.sh)
source /root/.bashrc
lxc stop $(lxc list | awk '{print $2}' | grep --invert-match NAME)
lxc delete $(lxc list | awk '{print $2}' | grep --invert-match NAME)
rm -rf container/minecraft-*
rm -rf properties/minecraft-*
cat drop_all.props | mongosh --port 27017
echo -n > container/latest_access
cp /usr/local/bin/minecraft/nginx.conf /etc/nginx/nginx.conf
cp /usr/local/bin/minecraft/nginx.conf /etc/nginx.conf
sudo rm -rf nohup*.out
kill -9 `pgrep init_server`
systemctl restart --now nginx
