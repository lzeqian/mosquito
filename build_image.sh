#!/bin/bash
go build -x
source ./version.sh
docker build -t liaomin789/mosquito:${MOS_VERSION} .
