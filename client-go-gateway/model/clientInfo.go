package model

import "github.com/hyperledger/fabric-gateway/pkg/client"

type ClientInfo struct {
	MspID        string           `json:"msp_id"`        // 所属组织的MSPID
	CryptoPath   string           `json:"crypto_path"`   // 该组织加密材料路径根路径
	CertPath     string           `json:"cert_path"`     // client 用户的签名证书
	KeyPath      string           `json:"key_path"`      // client 用户的私钥路径
	TlsCertPath  string           `json:"tls_cert_path"` // client 用户的 tls 通信证书
	PeerEndpoint string           `json:"peer_endpoint"` // 所连 peer 节点的地址 域名:端口
	GatewayPeer  string           `json:"gateway_peer"`  // 只包含域名,不含端口
	Contract     *client.Contract // 合约对象，通过此对象进行调用合约
	Live         bool             `json:"live"` //标识连接的peer是否存活
}
