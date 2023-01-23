#!/bin/bash
source /root/.bashrc
for i in $(lxc list | grep minecraft | awk '{print $2}');do lxc exec $i /bin/bash;done

