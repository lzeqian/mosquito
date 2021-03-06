#!/bin/bash
if docker ps -a | grep liaomin789/mosquito ;then
  docker stop mosquito
  docker rm mosquito
fi
docker rmi liaomin789/mosquito:1.0.0
