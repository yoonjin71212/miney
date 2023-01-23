if [ $(whoami) = "root" ]
then
    cp -r miney.service /usr/lib/systemd/system/miney.service
    cd ..
    rm -r /usr/local/bin/minecraft
    rm -r /usr/local/bin/apply_nginx.sh
    rm -r /usr/local/bin/clean.sh
    rm -r /usr/local/bin/conSSH.sh
    rm -r /usr/local/bin/container_creation.sh
    rm -r /usr/local/bin/delete_container.sh
    rm -r /usr/local/bin/easy_access.sh
    rm -r /usr/local/bin/add_port.sh
    rm -r /usr/local/bin/init_server.sh
    rm -r /usr/local/bin/remove-service.sh
    rm -r /usr/local/bin/initial_setup.sh
    rm -r /usr/local/bin/install_svc.sh
    rm -r /usr/local/bin/kill.sh
    rm -r /usr/local/bin/prepare.sh
    rm -r /usr/local/bin/server.sh
    rm -r /usr/local/bin/server_reload.sh
    rm -r /usr/local/bin/server
    echo  "Copying files..."
    mkdir /usr/local/bin/minecraft
    cp -Rf miney/* /usr/local/bin/minecraft
    ln -s /usr/local/bin/minecraft/*.sh /usr/local/bin
    ln -s /usr/local/bin/minecraft/server /usr/local/bin
    systemctl daemon-reload
    systemctl enable --now miney
    systemctl start  --now miney
    echo "Done"
else
    sudo -s
fi
