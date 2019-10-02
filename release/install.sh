#!/bin/bash
BASE=$(cd `dirname $0`; pwd)

INSTALL_DIR=/usr/local/bin/


## start install

mkdir -p ${INSTALL_DIR} &&  cd ${BASE} && cp port_forward_linux_64 port_forward ${INSTALL_DIR} && chmod +x /usr/local/bin/port_forward*

echo " Install Complete "


echo -e "\033[32m usage: port_forward
                \n  port_forward 0.0.0.0:53 47.93.244.11:53 udp
                \n  port_forward 0.0.0.0:2022 47.93.244.11:22 tcp \033[0m"