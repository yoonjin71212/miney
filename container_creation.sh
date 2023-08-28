#!/bin/bash
NET_INTERFACE="$(ip route get 1 | awk '{print $5}')"
TAG="$1"
PORT="$2"
VERSION="22.04"
SERVER_IP="$(ip route get 1 | awk '{print $7}')"
echo -n "TAG:"
echo $TAG
if [ $(arch)="x86_64" ]
then
		ARCH="amd64"
elif [ $(arch)="amd64" ]
then
		ARCH=$(arch)
else
		echo "Sorry, This architecture is not supported;" 1>&2
		echo "Supported architecture for Minecraft: amd64" 1>&2
		echo "Your architecture is $(arch)" 1>&2
		return
fi
lxc launch ubuntu/$VERSION/$ARCH $TAG 


while true ; do 
	CONTAINER_IP=`lxc list | grep minecraft | grep $TAG | awk '{print $6}' | grep --invert-match "|" | tr -s " " `
	LENGTH_IP=`echo $CONTAINER_IP | awk '{print length}'`
	if [  $LENGTH_IP = 0 ]; then
		sleep 0.5
	else 
		break
 fi
done
tail -n 1 /etc/nginx/nginx.conf | wc -c | xargs -I {} truncate /etc/nginx/nginx.conf -s -{}
echo "
	server {
		listen 0.0.0.0:$((PORT+1));
		proxy_pass $CONTAINER_IP:19133;
	}
	server {
		listen 0.0.0.0:$PORT;
		proxy_pass $CONTAINER_IP:19132;
	}
	server {
		listen 0.0.0.0:$((PORT+1)) udp;
		proxy_pass $CONTAINER_IP:19133;
	}
	server {
		listen 0.0.0.0:$PORT udp;
		proxy_pass $CONTAINER_IP:19132;
	}
	server {
		listen 0.0.0.0:$((PORT+2));
		proxy_pass $CONTAINER_IP:30000;
	}

}" >> /etc/nginx/nginx.conf

nginx -s reload
echo -n "CURRENT IP:"
echo $CONTAINER_IP
lxc file push -r /usr/local/bin/minecraft/miney.zip $TAG/
lxc exec $TAG -- /bin/apt install -y unzip
lxc exec $TAG -- /bin/unzip /miney.zip 
lxc exec $TAG -- /usr/bin/mv    miney /minecraft
lxc exec $TAG -- /bin/mkdir /usr/local/bin/minecraft/properties
touch /usr/local/bin/minecraft/container/$TAG
echo $CONTAINER_IP > /usr/local/bin/minecraft/container/$TAG
echo $TAG > /usr/local/bin/minecraft/container/latest_access
lxc exec $TAG -- /bin/apt install -y openssh-server 
lxc exec $TAG -- /bin/wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1-1ubuntu2.1\~18.04.20_amd64.deb
lxc exec $TAG -- /bin/dpkg -i libssl1.1_1.1.1-1ubuntu2.1\~18.04.20_amd64.deb
 
lxc file push -r /usr/local/bin/minecraft/minecraft-$TAG.properties $TAG/minecraft/server.properties
lxc exec $TAG -- /bin/rm -rf /etc/ssh/sshd_config
lxc exec $TAG -- /bin/systemctl restart --now ssh
lxc exec $TAG -- /bin/bash /minecraft/conSSH.sh $TAG
echo "LXC DEVICE STATUS:"
lxc list
