#!/bin/bash -u
docker stop $(docker ps -aq)
docker rm $(docker ps -a | grep fabric | awk '{print $1}')
docker rmi $(docker images dev-* -q)
sudo rm -rf orgs data
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d council.ifantasy.net soft.ifantasy.net web.ifantasy.net hard.ifantasy.net
sudo chmod 0777 -R orgs/