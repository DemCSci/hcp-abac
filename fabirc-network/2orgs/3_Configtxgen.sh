#!/bin/bash -eu
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d peer1.org1.lei.net  peer1.org2.lei.net 
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d orderer1.council.lei.net orderer2.council.lei.net orderer3.council.lei.net
sleep 5

configtxgen -profile OrgsChannel -outputCreateChannelTx $LOCAL_ROOT_PATH/data/testchannel.tx -channelID testchannel
configtxgen -profile OrgsChannel -outputBlock $LOCAL_ROOT_PATH/data/testchannel.block -channelID testchannel

cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/org1.lei.net/assets/
cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/org2.lei.net/assets/

source envpeer1org1
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.lei.net/registers/orderer1/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.lei.net/registers/orderer1/tls-msp/keystore/key.pem
osnadmin channel join -o orderer1.council.lei.net:7052 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer1.council.lei.net:7052 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.lei.net/registers/orderer2/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.lei.net/registers/orderer2/tls-msp/keystore/key.pem
osnadmin channel join -o orderer2.council.lei.net:7055 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer2.council.lei.net:7055 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.lei.net/registers/orderer3/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.lei.net/registers/orderer3/tls-msp/keystore/key.pem
osnadmin channel join -o orderer3.council.lei.net:7058 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer3.council.lei.net:7058 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

source envpeer1org1
peer channel join -b $LOCAL_CA_PATH/org1.lei.net/assets/testchannel.block
peer channel list

source envpeer1org2
peer channel join -b $LOCAL_CA_PATH/org2.lei.net/assets/testchannel.block
peer channel list