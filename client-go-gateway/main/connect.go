package main

import (
	"client-go-gateway/model"
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"io/ioutil"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 创建指向联盟链网络的 gRPC 连接.
func newGrpcConnection(info model.ClientInfo) *grpc.ClientConn {

	// client 用户的 tls 通信证书
	certificate, err := loadCertificate(info.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	//"peer1.soft.ifantasy.net"     // 网关 peer 节点名称
	transportCredentials := credentials.NewClientTLSFromCert(certPool, info.GatewayPeer)

	// 所连 peer 节点的地址 "peer1.soft.ifantasy.net:7251"
	connection, err := grpc.Dial(info.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// 根据用户指定的X.509证书为这个网关连接创建一个客户端标识。
func newIdentity(info model.ClientInfo) *identity.X509Identity {
	// client 用户的签名证书
	certificate, err := loadCertificate(info.CertPath)
	if err != nil {
		panic(err)
	}
	//// 所属组织的MSPID
	id, err := identity.NewX509Identity(info.MspID, certificate)
	if err != nil {
		panic(err)
	}
	return id
}

// 加载证书文件
func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// 使用私钥从消息摘要生成数字签名
func newSign(info model.ClientInfo) identity.Sign {
	// client 用户的私钥路径
	keyPath := info.KeyPath
	files, err := ioutil.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func newGateway(clintInfo model.ClientInfo) *client.Gateway {
	clientConnection1 := newGrpcConnection(clintInfo)
	//defer clientConnection1.Close()

	id := newIdentity(clintInfo)
	sign := newSign(clintInfo)

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection1),
		client.WithEvaluateTimeout(15*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(15*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	return gateway
}
