#!/bin/bash
if docker ps -a | grep liaomin789/mosquito ;then
  docker stop mosquito
  docker rm mosquito
fi
source ./version.sh
docker rmi liaomin789/mosquito:${MOS_VERSION}
