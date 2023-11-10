package main

import (
	"client-go-gateway/model"
	"client-go-gateway/routers"
	"client-go-gateway/setting"
	"log"
	"path"
	"strconv"
)

const (
	channelName   = "testchannel" // 连接的通道
	chaincodeName = "abac"        // 连接的链码
)

func main() {

	cryptoPath1 := "E:/code/orgs/org1.ifantasy.net"
	certPath1 := path.Join(cryptoPath1, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath1 := path.Join(cryptoPath1, "registers", "user1", "msp", "keystore")
	tlsCertPath1 := path.Join(cryptoPath1, "assets", "tls-ca-cert.pem")

	clientInfo1 := model.ClientInfo{
		MspID:        "org1MSP",
		CryptoPath:   cryptoPath1,
		CertPath:     certPath1,
		KeyPath:      keyPath1,
		TlsCertPath:  tlsCertPath1,
		PeerEndpoint: "peer1.org1.lei.net:7251",
		GatewayPeer:  "peer1.org1.lei.net",
	}

	gateway := newGateway(clientInfo1)
	defer gateway.Close()
	network := gateway.GetNetwork(channelName)
	clientInfo1.Contract = network.GetContract(chaincodeName)
	clientInfo1.Live = true
	log.Println("peer1.org1.lei.net 连接成功")

	/////////////////////////////clientInfo2

	//cryptoPath2 := "E:/code/orgs/web.lei.net"
	//certPath2 := path.Join(cryptoPath2, "registers", "user1", "msp", "signcerts", "cert.pem")
	//keyPath2 := path.Join(cryptoPath2, "registers", "user1", "msp", "keystore")
	//tlsCertPath2 := path.Join(cryptoPath2, "assets", "tls-ca-cert.pem")
	//clientInfo2 := model.ClientInfo{
	//	MspID:        "webMSP",
	//	CryptoPath:   cryptoPath2,
	//	CertPath:     certPath2,
	//	KeyPath:      keyPath2,
	//	TlsCertPath:  tlsCertPath2,
	//	PeerEndpoint: "peer1.web.lei.net:7351",
	//	GatewayPeer:  "peer1.web.lei.net",
	//}
	//gateway2 := newGateway(clientInfo2)
	//defer gateway2.Close()
	//network2 := gateway2.GetNetwork(channelName)
	//clientInfo2.Contract = network2.GetContract(chaincodeName)
	//clientInfo2.Live = true
	//log.Println("peer1.web.lei.net 连接成功")
	//

	//填到到map中去
	setting.ClientInfoMap[clientInfo1.MspID] = &clientInfo1
	//setting.ClientInfoMap[clientInfo2.MspID] = &clientInfo2

	//
	//setting.GlobalConsistent.Add(clientInfo1.MspID)
	//setting.GlobalConsistent.Add(clientInfo2.MspID)

	setting.Setup()

	router := routers.InitRouter(setting.WebSetting.ContextPath)

	err := router.Run(":" + strconv.Itoa(setting.WebSetting.Port))
	if err != nil {
		setting.MyLogger.Info("启动失败，err=", err)
		return
	}
}
