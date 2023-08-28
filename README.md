# Miney

minecraft server management

## Getting Started
####  This is minecraft server management service.
  You will get its port and info via response;
  you can create/delete LXD port with testcode_client.py.
  management port is minecraft_port+2.
##### Usage
* make
* ./initial_setup.sh
* systemctl start --now miney
### Front-end Application
#### Usage
cd app
python3 main.py
#### Requirements 
* python3,requests,kivy

### Back-end Application
* Written in Go, binary is generated when you run Make Task
* Working as REST API Server.
#### Commands
* create:POST Method
* delete:POST Method
* request: GET Method

#### About Virtual Machines Management
* Virtual Machines are managed by LXD.
* All connections can be managed in a single domain
* Used Nginx Reverse-Proxy
