#!/bin/bash
TAG="$1"
source /root/.bashrc
lxc stop $TAG
lxc delete $TAG
rm -rf /usr/lobal/bin/minecraft/ontainer/minecraft-*
rm -rf /usr/local/bin/minecraft/properties/minecraft-*
cat drop_all.props | mongosh --port 27017
echo -n > /usr/local/bin/minecraft/container/latest_access
cp /usr/local/bin/minecraft/nginx.conf /etc/nginx/nginx.conf
sudo rm -rf nohup*.out

kill -9 `pgrep init_server`
