#!/bin/bash
docker stop $(docker ps -aq)
docker rm $(docker ps -a | grep fabric | awk '{print $1}')
docker rmi $(docker images dev-* -q)
docker rm $(docker ps -a | grep 'dev-*' | awk '{print $1}')
sudo rm -rf orgs data
docker volume prune
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d council.lei.net org1.lei.net org3.lei.net org2.lei.net org4.lei.net org5.lei.net
sudo chmod 0777 -R orgs/
