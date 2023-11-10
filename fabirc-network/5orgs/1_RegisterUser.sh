#!/bin/bash -eu
set -e
echo "Working on council"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@council.lei.net:7050
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer1 --id.secret orderer1 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer2 --id.secret orderer2 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer3 --id.secret orderer3 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer4 --id.secret orderer4 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer5 --id.secret orderer5 --id.type orderer -u https://council.lei.net:7050

fabric-ca-client register -d --id.name peer1org1 --id.secret peer1org1 --id.type peer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1web --id.secret peer1web --id.type peer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1org2 --id.secret peer1org2 --id.type peer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1org4 --id.secret peer1org4 --id.type peer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1org5 --id.secret peer1org5 --id.type peer -u https://council.lei.net:7050


echo "Working on org1"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org1.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org1.lei.net:7250
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org1.lei.net:7250
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org1.lei.net:7250
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org1.lei.net:7250
echo "Working on web"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/web.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@web.lei.net:7350
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://web.lei.net:7350
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://web.lei.net:7350
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://web.lei.net:7350
echo "Working on org2"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org2.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org2.lei.net:7450
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org2.lei.net:7450
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org2.lei.net:7450
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org2.lei.net:7450

# 第4个组织
echo "Working on org4"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org4.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org4.lei.net:7550
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org4.lei.net:7550
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org4.lei.net:7550
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org4.lei.net:7550

# 第5个组织
echo "Working on org5"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org5.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org5.lei.net:7650
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org5.lei.net:7650
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org5.lei.net:7650
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org5.lei.net:7650

echo "All CA and registration done"