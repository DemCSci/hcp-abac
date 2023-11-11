#!/bin/bash
docker stop $(docker ps -aq)
docker rm $(docker ps -a | grep fabric | awk '{print $1}')
docker rmi $(docker images dev-* -q)
docker rm $(docker ps -a | grep 'dev-*' | awk '{print $1}')
sudo rm -rf orgs data
docker volume prune
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d council.ifantasy.net org1.ifantasy.net web.ifantasy.net org2.ifantasy.net
sudo chmod 0777 -R orgs/