package main

import (
	"client-go-gateway/contract"
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
	pool, err2 := ants.NewPool(6)
	if err2 != nil {
		log.Fatal("goroutine 池子创建失败")
	}
	util.Pool = *pool

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

	contract.Contract1 = network.GetContract(chaincodeName)
	log.Println("peer1.soft.ifantasy.net 连接成功")
	//log.Println("getAllSeets from clientInfo1")
	//getAllAssets(contract)

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
	contract.Contract2 = network2.GetContract(chaincodeName)
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
		PeerEndpoint: "peer1.hard.ifantasy.net:7351",
		GatewayPeer:  "peer1.hard.ifantasy.net",
	}
	gateway3 := newGateway(clientInfo3)
	defer gateway3.Close()
	network3 := gateway3.GetNetwork(channelName)
	contract.Contract3 = network3.GetContract(chaincodeName)

	contract.ContractList = append(contract.ContractList, contract.Contract1)
	contract.ContractList = append(contract.ContractList, contract.Contract2)
	contract.ContractList = append(contract.ContractList, contract.Contract3)

	log.Println("peer1.hard.ifantasy.net 连接成功")

	log.Println("启动web服务 :7788")
	http.HandleFunc("/decideNoRecord", controller.DecideNoRecord)
	http.HandleFunc("/decideWithRecord", controller.DecideWithRecord)
	http.HandleFunc("/users", controller.GetAllUsers)
	http.HandleFunc("/addUser", controller.AddUser)
	http.HandleFunc("/my", controller.GetSubmittingClientIdentity)

	err := http.ListenAndServe(":7788", nil)
	if err != nil {
		log.Fatalln(err)
	}
}