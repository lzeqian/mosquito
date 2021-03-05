#!/bin/bash
docker stop mosquito
docker rm mosquito
docker rmi liaomin789/mosquito:1.0.0
