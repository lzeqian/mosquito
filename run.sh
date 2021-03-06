mkdir -p /etc/mosquito
setenforce 0
docker run -d -p 8089:8089  -v /etc/mosquito:/application/rbac --env-file ./.envfile --name mosquito liaomin789/mosquito:1.0.0
