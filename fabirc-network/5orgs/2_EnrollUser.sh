#!/bin/bash -eu
set -e
echo "Preparation============================="
mkdir -p $LOCAL_CA_PATH/council.lei.net/assets
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org1.lei.net/assets
cp $LOCAL_CA_PATH/org1.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org1.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org3.lei.net/assets 
cp $LOCAL_CA_PATH/org3.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org3.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org3.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org2.lei.net/assets
cp $LOCAL_CA_PATH/org2.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org2.lei.net/assets/tls-ca-cert.pem

mkdir -p $LOCAL_CA_PATH/org4.lei.net/assets
cp $LOCAL_CA_PATH/org4.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org4.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org4.lei.net/assets/tls-ca-cert.pem


mkdir -p $LOCAL_CA_PATH/org5.lei.net/assets
cp $LOCAL_CA_PATH/org5.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org5.lei.net/assets/ca-cert.pem
cp $LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem $LOCAL_CA_PATH/org5.lei.net/assets/tls-ca-cert.pem


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

echo "Enroll Orderer4"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/orderer4
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer4:orderer4@council.lei.net:7050
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/orderer4/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/orderer4/msp/admincerts/cert.pem
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://orderer4:orderer4@council.lei.net:7050 --enrollment.profile tls --csr.hosts orderer4.council.lei.net
cp $LOCAL_CA_PATH/council.lei.net/registers/orderer4/tls-msp/keystore/*_sk $LOCAL_CA_PATH/council.lei.net/registers/orderer4/tls-msp/keystore/key.pem

echo "Enroll Orderer5"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/registers/orderer5
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://orderer5:orderer5@council.lei.net:7050
mkdir -p $LOCAL_CA_PATH/council.lei.net/registers/orderer5/msp/admincerts
cp $LOCAL_CA_PATH/council.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/council.lei.net/registers/orderer5/msp/admincerts/cert.pem
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://orderer5:orderer5@council.lei.net:7050 --enrollment.profile tls --csr.hosts orderer5.council.lei.net
cp $LOCAL_CA_PATH/council.lei.net/registers/orderer5/tls-msp/keystore/*_sk $LOCAL_CA_PATH/council.lei.net/registers/orderer5/tls-msp/keystore/key.pem



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

echo "Start org3============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org3.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org3.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@org3.lei.net:7350
mkdir -p $LOCAL_CA_PATH/org3.lei.net/registers/user1/msp/admincerts

echo "Enroll Admin1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org3.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org3.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@org3.lei.net:7350
mkdir -p $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
# for identity
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org3.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org3.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@org3.lei.net:7350
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org3.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1org3:peer1org3@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.org3.lei.net
cp $LOCAL_CA_PATH/org3.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/org3.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/org3.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org3.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/org3.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/org3.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/org3.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/org3.lei.net/msp/users
cp $LOCAL_CA_PATH/org3.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/org3.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/org3.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/org3.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org3.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org3.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org3.lei.net/registers/user1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org3.lei.net/registers/admin1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org3.lei.net/registers/peer1/msp
echo "End org3============================="

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

echo "Start org4============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org4.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@org4.lei.net:7550
mkdir -p $LOCAL_CA_PATH/org4.lei.net/registers/user1/msp/admincerts

echo "Enroll Admin"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org4.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@org4.lei.net:7550
mkdir -p $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org4.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@org4.lei.net:7550
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1org4:peer1org4@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.org4.lei.net
cp $LOCAL_CA_PATH/org4.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/org4.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/org4.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org4.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/org4.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/org4.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/org4.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/org4.lei.net/msp/users
cp $LOCAL_CA_PATH/org4.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/org4.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/org4.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/org4.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org4.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org4.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org4.lei.net/registers/user1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org4.lei.net/registers/admin1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org4.lei.net/registers/peer1/msp
echo "End org4============================="

echo "Start org5============================="
echo "Enroll User1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org5.lei.net/registers/user1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://user1:user1@org5.lei.net:7650
mkdir -p $LOCAL_CA_PATH/org5.lei.net/registers/user1/msp/admincerts

echo "Enroll Admin"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org5.lei.net/registers/admin1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://admin1:admin1@org5.lei.net:7650
mkdir -p $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp/admincerts
cp $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp/admincerts/cert.pem

echo "Enroll Peer1"
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org5.lei.net/registers/peer1
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.lei.net/assets/ca-cert.pem
export FABRIC_CA_CLIENT_MSPDIR=msp
fabric-ca-client enroll -d -u https://peer1:peer1@org5.lei.net:7650
# for TLS
export FABRIC_CA_CLIENT_MSPDIR=tls-msp
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.lei.net/assets/tls-ca-cert.pem
fabric-ca-client enroll -d -u https://peer1org5:peer1org5@council.lei.net:7050 --enrollment.profile tls --csr.hosts peer1.org5.lei.net
cp $LOCAL_CA_PATH/org5.lei.net/registers/peer1/tls-msp/keystore/*_sk $LOCAL_CA_PATH/org5.lei.net/registers/peer1/tls-msp/keystore/key.pem
mkdir -p $LOCAL_CA_PATH/org5.lei.net/registers/peer1/msp/admincerts
cp $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org5.lei.net/registers/peer1/msp/admincerts/cert.pem

mkdir -p $LOCAL_CA_PATH/org5.lei.net/msp/admincerts
mkdir -p $LOCAL_CA_PATH/org5.lei.net/msp/cacerts
mkdir -p $LOCAL_CA_PATH/org5.lei.net/msp/tlscacerts
mkdir -p $LOCAL_CA_PATH/org5.lei.net/msp/users
cp $LOCAL_CA_PATH/org5.lei.net/assets/ca-cert.pem $LOCAL_CA_PATH/org5.lei.net/msp/cacerts/
cp $LOCAL_CA_PATH/org5.lei.net/assets/tls-ca-cert.pem $LOCAL_CA_PATH/org5.lei.net/msp/tlscacerts/
cp $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp/signcerts/cert.pem $LOCAL_CA_PATH/org5.lei.net/msp/admincerts/cert.pem

cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org5.lei.net/msp/config.yaml
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org5.lei.net/registers/user1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org5.lei.net/registers/admin1/msp
cp $LOCAL_ROOT_PATH/config/config-msp.yaml $LOCAL_CA_PATH/org5.lei.net/registers/peer1/msp
echo "End org5============================="


find orgs/ -regex ".+cacerts.+.pem" -not -regex ".+tlscacerts.+" | rename 's/cacerts\/.+\.pem/cacerts\/ca-cert\.pem/'