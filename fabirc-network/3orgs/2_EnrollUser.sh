#!/bin/bash -eu
echo "Preparation============================="
mkdir -p $LOCAL_CA_PATH/council.lei.net/assets
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org1.lei.net/assets
cp $LOCAL_CA_PATH/org1.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/web.lei.net/assets 
cp $LOCAL_CA_PATH/web.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/web.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/web.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org2.lei.net/assets
cp $LOCAL_CA_PATH/org2.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/assets/tls-ca-cert.pem
echo "Preparation end=========================="

echo "Start Council============================="
echo "Enroll Admin"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@council.lei.net:7050
# 加入通道时会用到admin/msp，其下必须要有admincers
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Orderer1"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/orderer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer1:orderer1@council.lei.net:7050
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/orderer1/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/orderer1/msp/admincerts/cert.pem
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://orderer1:orderer1@council.lei.net:7050 --enrollment.profile tls --csr.hosts orderer1.council.lei.net
cp $LOCAL_CA_PATH/council.lei.net/registers/orderer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/council.lei.net/registers/orderer1/tls-msp/keystore/key.pem

echo "Enroll Orderer2"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/orderer2
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer2:orderer2@council.lei.net:7050
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/orderer2/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/orderer2/msp/admincerts/cert.pem
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://orderer2:orderer2@council.lei.net:7050 --enrollment.profile tls --csr.hosts orderer2.council.lei.net
cp $LOCAL_CA_PATH/council.lei.net/registers/orderer2/tls-msp/keystore/*_sk $LOCAL_CA_PATH/council.lei.net/registers/orderer2/tls-msp/keystore/key.pem

echo "Enroll Orderer3"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/orderer3
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer3:orderer3@council.lei.net:7050
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/orderer3/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/orderer3/msp/admincerts/cert.pem
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://orderer3:orderer3@council.lei.net:7050 --enrollment.profile tls --csr.hosts orderer3.council.lei.net
cp $LOCAL_CA_PATH/council.lei.net/registers/orderer3/tls-msp/keystore/*_sk $LOCAL_CA_PATH/council.lei.net/registers/orderer3/tls-msp/keystore/key.pem

mkdir -p $LOCAL_CA_PATH/council.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/council.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/council.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/council.lei.net/msp/users
cp $LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/council.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/council.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/msp/admincerts/cert.pem
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/council.lei.net/msp/config.yaml
echo "End council============================="


echo "Start org1============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org1.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@org1.lei.net:7250

echo "Enroll Admin1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org1.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@org1.lei.net:7250
mkdir -p $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org1.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@org1.lei.net:7250
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1org1:peer1org1@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.org1.lei.net
cp $LOCAL_CA_PATH/org1.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/org1.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/org1.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org1.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/org1.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/org1.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/org1.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/org1.lei.net/msp/users
cp $LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/org1.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org1.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org1.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org1.lei.net/registers/user1/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org1.lei.net/registers/admin1/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org1.lei.net/registers/peer1/msp/config.yaml
echo "End org1============================="

echo "Start Web============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/web.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@web.lei.net:7350
mkdir -p $LOCAL_CA_PATH/web.lei.net/registers/user1/msp/admincerts

echo "Enroll Admin1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/web.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@web.lei.net:7350
mkdir -p $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/web.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@web.lei.net:7350
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1web:peer1web@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.web.lei.net
cp $LOCAL_CA_PATH/web.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/web.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/web.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/web.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/web.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/web.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/web.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/web.lei.net/msp/users
cp $LOCAL_CA_PATH/web.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/web.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/web.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/web.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/web.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/web.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/web.lei.net/registers/user1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/web.lei.net/registers/admin1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/web.lei.net/registers/peer1/msp
echo "End Web============================="

echo "Start org2============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org2.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@org2.lei.net:7450
mkdir -p $LOCAL_CA_PATH/org2.lei.net/registers/user1/msp/admincerts

echo "Enroll Admin"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org2.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@org2.lei.net:7450
mkdir -p $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org2.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@org2.lei.net:7450
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1org2:peer1org2@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.org2.lei.net
cp $LOCAL_CA_PATH/org2.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/org2.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/org2.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org2.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/org2.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/org2.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/org2.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/org2.lei.net/msp/users
cp $LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/org2.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org2.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org2.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org2.lei.net/registers/user1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org2.lei.net/registers/admin1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org2.lei.net/registers/peer1/msp
echo "End org2============================="

find orgs/ -regex ".+cacerts.+.pem" -not -regex ".+tlscacerts.+" | rename 's/cacerts\/.+\.pem/cacerts\/ca-cert\.pem/'