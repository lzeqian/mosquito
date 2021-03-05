#!/bin/bash
go build -x
docker build -t liaomin789/mosquito:1.0.0 .
mkdir -p /etc/mosquito
cp -rf conf/app.conf /etc/mosquito/app.conf
cp -rf conf/rbac.yml /etc/mosquito/rbac.yml
setenforce 0
docker run -d -p 8089:8089  -v /etc/mosquito:/application/conf --name mosquito liaomin789/mosquito:1.0.0
