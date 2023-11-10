#!/bin/bash -eu
echo "Working on council"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/council.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/council.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@council.lei.net:7050
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer1 --id.secret orderer1 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer2 --id.secret orderer2 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name orderer3 --id.secret orderer3 --id.type orderer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1org1 --id.secret peer1org1 --id.type peer -u https://council.lei.net:7050
fabric-ca-client register -d --id.name peer1org2 --id.secret peer1org2 --id.type peer -u https://council.lei.net:7050
echo "Working on org1"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org1.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org1.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org1.lei.net:7250
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org1.lei.net:7250
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org1.lei.net:7250
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org1.lei.net:7250

echo "Working on org2"
export FABRIC_CA_CLIENT_TLS_CERTFILES=$LOCAL_CA_PATH/org2.lei.net/ca/crypto/ca-cert.pem
export FABRIC_CA_CLIENT_HOME=$LOCAL_CA_PATH/org2.lei.net/ca/admin
fabric-ca-client enroll -d -u https://ca-admin:ca-adminpw@org2.lei.net:7450
fabric-ca-client register -d --id.name user1 --id.secret user1 --id.type client -u https://org2.lei.net:7450
fabric-ca-client register -d --id.name peer1 --id.secret peer1 --id.type peer -u https://org2.lei.net:7450
fabric-ca-client register -d --id.name admin1 --id.secret admin1 --id.type admin -u https://org2.lei.net:7450
echo "All CA and registration done"