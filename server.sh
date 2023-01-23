#!/bin/bash
nohup /usr/local/bin/minecraft/server "$(ip route get 1 | awk '{print $7}')" &
