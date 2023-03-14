#!/bin/bash -eu
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d peer1.soft.ifantasy.net peer1.web.ifantasy.net peer1.hard.ifantasy.net peer1.org4.ifantasy.net peer1.org5.ifantasy.net
docker-compose -f $LOCAL_ROOT_PATH/compose/docker-compose.yaml up -d orderer1.council.ifantasy.net orderer2.council.ifantasy.net orderer3.council.ifantasy.net orderer4.council.ifantasy.net orderer5.council.ifantasy.net
sleep 5

configtxgen -profile OrgsChannel -outputCreateChannelTx $LOCAL_ROOT_PATH/data/testchannel.tx -channelID testchannel
configtxgen -profile OrgsChannel -outputBlock $LOCAL_ROOT_PATH/data/testchannel.block -channelID testchannel

cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/soft.ifantasy.net/assets/
cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/web.ifantasy.net/assets/
cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/hard.ifantasy.net/assets/
cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/org4.ifantasy.net/assets/
cp $LOCAL_ROOT_PATH/data/testchannel.block $LOCAL_CA_PATH/org5.ifantasy.net/assets/

source envpeer1soft
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer1/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer1/tls-msp/keystore/key.pem
osnadmin channel join -o orderer1.council.ifantasy.net:7052 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer1.council.ifantasy.net:7052 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer2/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer2/tls-msp/keystore/key.pem
osnadmin channel join -o orderer2.council.ifantasy.net:7055 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer2.council.ifantasy.net:7055 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY
export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer3/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer3/tls-msp/keystore/key.pem
osnadmin channel join -o orderer3.council.ifantasy.net:7058 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer3.council.ifantasy.net:7058 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer4/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer4/tls-msp/keystore/key.pem
osnadmin channel join -o orderer4.council.ifantasy.net:7061 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer4.council.ifantasy.net:7061 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

export ORDERER_ADMIN_TLS_SIGN_CERT=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer5/tls-msp/signcerts/cert.pem
export ORDERER_ADMIN_TLS_PRIVATE_KEY=$LOCAL_CA_PATH/council.ifantasy.net/registers/orderer5/tls-msp/keystore/key.pem
osnadmin channel join -o orderer5.council.ifantasy.net:7064 --channelID testchannel --config-block $LOCAL_ROOT_PATH/data/testchannel.block --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY"
osnadmin channel list -o orderer5.council.ifantasy.net:7064 --ca-file $ORDERER_CA --client-cert $ORDERER_ADMIN_TLS_SIGN_CERT --client-key $ORDERER_ADMIN_TLS_PRIVATE_KEY

source envpeer1soft
peer channel join -b $LOCAL_CA_PATH/soft.ifantasy.net/assets/testchannel.block
peer channel list
source envpeer1web
peer channel join -b $LOCAL_CA_PATH/web.ifantasy.net/assets/testchannel.block
peer channel list
source envpeer1hard
peer channel join -b $LOCAL_CA_PATH/hard.ifantasy.net/assets/testchannel.block
peer channel list
source envpeer1org4
peer channel join -b $LOCAL_CA_PATH/org4.ifantasy.net/assets/testchannel.block
peer channel list
source envpeer1org5
peer channel join -b $LOCAL_CA_PATH/org5.ifantasy.net/assets/testchannel.block
peer channel list