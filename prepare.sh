#!/bin/bash
apt-get update -y
apt-get install -y sshpass
cd /
mv miney minecraft
cd /minecraft
unzip *.zip
cd properties
cp $(hostname).properties ../server.properties
echo  >> ../server.properties
echo "emit-server-telemetry=true" >> ../server.properties
cd ..
nohup ./bedrock_server &

