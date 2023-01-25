#!/bin/bash
if [ "$(whoami)" == "root" ]
then
    rm -r /usr/local/bin/minecraft
    rm -r /usr/local/bin/apply_nginx.sh
    rm -r /usr/local/bin/clean.sh
    rm -r /usr/local/bin/conSSH.sh
    rm -r /usr/local/bin/container_creation.sh
    rm -r /usr/local/bin/delete_container.sh
    rm -r /usr/local/bin/easy_access.sh
    rm -r /usr/local/bin/remove-service.sh
    rm -r /usr/local/bin/add_port.sh
    rm -r /usr/local/bin/init_server.sh
    rm -r /usr/local/bin/initial_setup.sh
    rm -r /usr/local/bin/install_svc.sh
    rm -r /usr/local/bin/kill.sh
    rm -r /usr/local/bin/prepare.sh
    rm -r /usr/local/bin/server.sh
    rm -r /usr/local/bin/server_reload.sh
    rm -r /usr/local/bin/server
    systemctl disable --now miney
    rm -r /usr/lib/systemd/system/miney.service
fi
