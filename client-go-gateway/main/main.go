package main

import (
	"client-go-gateway/controller"
	"client-go-gateway/model"
	"client-go-gateway/util"
	"github.com/panjf2000/ants/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"path"
)

const (
	channelName   = "testchannel" // 连接的通道
	chaincodeName = "abac"        // 连接的链码
)

func main() {
	pool, err2 := ants.NewPool(3, ants.WithNonblocking(false))

	if err2 != nil {
		log.Fatal("goroutine 池子创建失败")
	}
	util.Pool = pool

	util.Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	cryptoPath1 := "E:/code/orgs/soft.ifantasy.net"
	certPath1 := path.Join(cryptoPath1, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath1 := path.Join(cryptoPath1, "registers", "user1", "msp", "keystore")
	tlsCertPath1 := path.Join(cryptoPath1, "assets", "tls-ca-cert.pem")

	clientInfo1 := model.ClientInfo{
		MspID:        "softMSP",
		CryptoPath:   cryptoPath1,
		CertPath:     certPath1,
		KeyPath:      keyPath1,
		TlsCertPath:  tlsCertPath1,
		PeerEndpoint: "peer1.soft.ifantasy.net:7251",
		GatewayPeer:  "peer1.soft.ifantasy.net",
	}

	gateway := newGateway(clientInfo1)
	defer gateway.Close()
	network := gateway.GetNetwork(channelName)
	clientInfo1.Contract = network.GetContract(chaincodeName)
	clientInfo1.Live = true
	log.Println("peer1.soft.ifantasy.net 连接成功")

	/////////////////////////////clientInfo2

	cryptoPath2 := "E:/code/orgs/web.ifantasy.net"
	certPath2 := path.Join(cryptoPath2, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath2 := path.Join(cryptoPath2, "registers", "user1", "msp", "keystore")
	tlsCertPath2 := path.Join(cryptoPath2, "assets", "tls-ca-cert.pem")
	clientInfo2 := model.ClientInfo{
		MspID:        "webMSP",
		CryptoPath:   cryptoPath2,
		CertPath:     certPath2,
		KeyPath:      keyPath2,
		TlsCertPath:  tlsCertPath2,
		PeerEndpoint: "peer1.web.ifantasy.net:7351",
		GatewayPeer:  "peer1.web.ifantasy.net",
	}
	gateway2 := newGateway(clientInfo2)
	defer gateway2.Close()
	network2 := gateway2.GetNetwork(channelName)
	clientInfo2.Contract = network2.GetContract(chaincodeName)
	clientInfo2.Live = true
	log.Println("peer1.web.ifantasy.net 连接成功")

	/////////////////////////////clientInfo3

	cryptoPath3 := "E:/code/orgs/hard.ifantasy.net"
	certPath3 := path.Join(cryptoPath3, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath3 := path.Join(cryptoPath3, "registers", "user1", "msp", "keystore")
	tlsCertPath3 := path.Join(cryptoPath3, "assets", "tls-ca-cert.pem")
	clientInfo3 := model.ClientInfo{
		MspID:        "hardMSP",
		CryptoPath:   cryptoPath3,
		CertPath:     certPath3,
		KeyPath:      keyPath3,
		TlsCertPath:  tlsCertPath3,
		PeerEndpoint: "peer1.hard.ifantasy.net:7451",
		GatewayPeer:  "peer1.hard.ifantasy.net",
	}
	gateway3 := newGateway(clientInfo3)
	defer gateway3.Close()
	network3 := gateway3.GetNetwork(channelName)
	clientInfo3.Contract = network3.GetContract(chaincodeName)
	clientInfo3.Live = true
	log.Println("peer1.hard.ifantasy.net 连接成功")

	/////////////////////////////clientInfo4
	cryptoPath4 := "E:/code/orgs/org4.ifantasy.net"
	certPath4 := path.Join(cryptoPath4, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath4 := path.Join(cryptoPath4, "registers", "user1", "msp", "keystore")
	tlsCertPath4 := path.Join(cryptoPath4, "assets", "tls-ca-cert.pem")
	clientInfo4 := model.ClientInfo{
		MspID:        "org4MSP",
		CryptoPath:   cryptoPath4,
		CertPath:     certPath4,
		KeyPath:      keyPath4,
		TlsCertPath:  tlsCertPath4,
		PeerEndpoint: "peer1.org4.ifantasy.net:7551",
		GatewayPeer:  "peer1.org4.ifantasy.net",
	}
	gateway4 := newGateway(clientInfo4)
	defer gateway4.Close()
	network4 := gateway4.GetNetwork(channelName)
	clientInfo4.Contract = network4.GetContract(chaincodeName)
	clientInfo4.Live = true
	log.Println("peer1.org4.ifantasy.net 连接成功")

	/////////////////////////////clientInfo5
	cryptoPath5 := "E:/code/orgs/org5.ifantasy.net"
	certPath5 := path.Join(cryptoPath5, "registers", "user1", "msp", "signcerts", "cert.pem")
	keyPath5 := path.Join(cryptoPath5, "registers", "user1", "msp", "keystore")
	tlsCertPath5 := path.Join(cryptoPath5, "assets", "tls-ca-cert.pem")
	clientInfo5 := model.ClientInfo{
		MspID:        "org5MSP",
		CryptoPath:   cryptoPath5,
		CertPath:     certPath5,
		KeyPath:      keyPath5,
		TlsCertPath:  tlsCertPath5,
		PeerEndpoint: "peer1.org5.ifantasy.net:7651",
		GatewayPeer:  "peer1.org5.ifantasy.net",
	}
	gateway5 := newGateway(clientInfo5)
	defer gateway5.Close()
	network5 := gateway5.GetNetwork(channelName)
	clientInfo5.Contract = network5.GetContract(chaincodeName)
	clientInfo5.Live = true
	log.Println("peer1.org5.ifantasy.net 连接成功")

	//填到到map中去
	util.ClientInfoMap[clientInfo1.MspID] = clientInfo1.Contract
	util.ClientInfoMap[clientInfo2.MspID] = clientInfo2.Contract
	util.ClientInfoMap[clientInfo3.MspID] = clientInfo3.Contract
	util.ClientInfoMap[clientInfo4.MspID] = clientInfo4.Contract
	util.ClientInfoMap[clientInfo5.MspID] = clientInfo5.Contract

	util.GlobalConsistent.Add(clientInfo1.MspID)
	util.GlobalConsistent.Add(clientInfo2.MspID)
	util.GlobalConsistent.Add(clientInfo3.MspID)
	util.GlobalConsistent.Add(clientInfo4.MspID)
	util.GlobalConsistent.Add(clientInfo5.MspID)

	log.Println("启动web服务 :7788")
	http.HandleFunc("/decideNoRecord", controller.DecideNoRecord)
	http.HandleFunc("/decideNoRecordPool", controller.DecideNoRecordPool)
	http.HandleFunc("/DecideHashNoRecordPool", controller.DecideHashNoRecordPool)
	http.HandleFunc("/decideNoRecordRedis", controller.DecideNoRecordRedis)
	http.HandleFunc("/decideWithRecord", controller.DecideWithRecord)

	http.HandleFunc("/users", controller.GetAllUsers)
	http.HandleFunc("/addUser", controller.AddUser)
	http.HandleFunc("/addAllUser", controller.AddAllUser)
	http.HandleFunc("/my", controller.GetSubmittingClientIdentity)

	err := http.ListenAndServe(":7788", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
