#!/bin/bash -eu
set -e
echo "Working on council"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@council.ifantasy.net:7050
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name orderer1 --id.secret orderer1 --id.type orderer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name orderer2 --id.secret orderer2 --id.type orderer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name orderer3 --id.secret orderer3 --id.type orderer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name orderer4 --id.secret orderer4 --id.type orderer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name orderer5 --id.secret orderer5 --id.type orderer -u https://council.ifantasy.net:7050

fabric-ca-client register -d --id.name peer1soft --id.secret peer1soft --id.type peer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name peer1web --id.secret peer1web --id.type peer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name peer1hard --id.secret peer1hard --id.type peer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name peer1org4 --id.secret peer1org4 --id.type peer -u https://council.ifantasy.net:7050
fabric-ca-client register -d --id.name peer1org5 --id.secret peer1org5 --id.type peer -u https://council.ifantasy.net:7050


echo "Working on soft"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/soft.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/soft.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@soft.ifantasy.net:7250
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://soft.ifantasy.net:7250
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://soft.ifantasy.net:7250
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://soft.ifantasy.net:7250
echo "Working on web"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/web.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/web.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@web.ifantasy.net:7350
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://web.ifantasy.net:7350
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://web.ifantasy.net:7350
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://web.ifantasy.net:7350
echo "Working on hard"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/hard.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/hard.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@hard.ifantasy.net:7450
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://hard.ifantasy.net:7450
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://hard.ifantasy.net:7450
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://hard.ifantasy.net:7450

# 第4个组织
echo "Working on org4"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org4.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org4.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org4.ifantasy.net:7550
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org4.ifantasy.net:7550
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org4.ifantasy.net:7550
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org4.ifantasy.net:7550

# 第5个组织
echo "Working on org5"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org5.ifantasy.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org5.ifantasy.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org5.ifantasy.net:7650
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org5.ifantasy.net:7650
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org5.ifantasy.net:7650
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org5.ifantasy.net:7650

echo "All CA and registration done"