package controller

import (
	"client-go-gateway/contract"
	"client-go-gateway/model"
	"client-go-gateway/request"
	"client-go-gateway/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func DecideNoRecord(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	var request request.DecideRequest
	err = json.Unmarshal(bodyByte, &request)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	contractResponse1 := contract.DecideNoRecord(util.ClientInfoMap["softMSP"], request)
	contractResponse2 := contract.DecideNoRecord(util.ClientInfoMap["webMSP"], request)
	contractResponse3 := contract.DecideNoRecord(util.ClientInfoMap["hardMSP"], request)
	contractResponse4 := contract.DecideNoRecord(util.ClientInfoMap["org4MSP"], request)
	contractResponse5 := contract.DecideNoRecord(util.ClientInfoMap["org5MSP"], request)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "结果不一致，拒绝该请求")
		return
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + request.ResourceId + ":" + request.RequesterId + ":" + util.GetUUID(),
		RequesterId: request.RequesterId,
		ResourceId:  request.RequesterId,
		Response:    contractResponse1,
	}

	contract.CreateRecord(util.ClientInfoMap["softMSP"], *record)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, contractResponse1)
}

/**
使用携程池的方式异步上链
*/
func DecideNoRecordPool(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	var request request.DecideRequest
	err = json.Unmarshal(bodyByte, &request)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	contractResponse1 := contract.DecideNoRecord(util.ClientInfoMap["softMSP"], request)
	contractResponse2 := contract.DecideNoRecord(util.ClientInfoMap["webMSP"], request)
	contractResponse3 := contract.DecideNoRecord(util.ClientInfoMap["hardMSP"], request)
	contractResponse4 := contract.DecideNoRecord(util.ClientInfoMap["org4MSP"], request)
	contractResponse5 := contract.DecideNoRecord(util.ClientInfoMap["org5MSP"], request)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "结果不一致，拒绝该请求")
		return
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + request.ResourceId + ":" + request.RequesterId + ":" + util.GetUUID(),
		RequesterId: request.RequesterId,
		ResourceId:  request.RequesterId,
		Response:    contractResponse1,
	}
	util.Pool.Submit(func() {
		contract.CreateRecord(util.ClientInfoMap["softMSP"], *record)
		time.Sleep(time.Second * 1)
	})
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, contractResponse1)
}

/**
使用redis消息队列异步上链
*/
func DecideNoRecordRedis(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	var request request.DecideRequest
	err = json.Unmarshal(bodyByte, &request)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	contractResponse1 := contract.DecideNoRecord(util.ClientInfoMap["softMSP"], request)
	contractResponse2 := contract.DecideNoRecord(util.ClientInfoMap["webMSP"], request)
	contractResponse3 := contract.DecideNoRecord(util.ClientInfoMap["hardMSP"], request)
	contractResponse4 := contract.DecideNoRecord(util.ClientInfoMap["org4MSP"], request)
	contractResponse5 := contract.DecideNoRecord(util.ClientInfoMap["org5MSP"], request)

	if contractResponse1 != contractResponse2 && contractResponse1 != contractResponse3 &&
		contractResponse1 != contractResponse4 && contractResponse1 != contractResponse5 {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "结果不一致，拒绝该请求")
		return
	}
	//异步发送record
	record := &model.Record{
		Id:          "record:" + request.ResourceId + ":" + request.RequesterId + ":" + util.GetUUID(),
		RequesterId: request.RequesterId,
		ResourceId:  request.RequesterId,
		Response:    contractResponse1,
	}
	//util.Pool.Submit(func() {
	//	contract.CreateRecord(contract.Contract2, *record)
	//	time.Sleep(time.Second * 1)
	//})

	//contract.CreateRecord(contract.Contract1, *record)
	recordJsonByte, err := json.Marshal(record)
	util.Rdb.Publish(util.Ctx, "test-channel", string(recordJsonByte))
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, contractResponse1)
}

func DecideWithRecord(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("读取body内容失败")
	}
	var request request.DecideRequest
	err = json.Unmarshal(bodyByte, &request)
	if err != nil {
		log.Fatal("反序列化失败")
	}
	record := contract.DecideWithRecord(util.ClientInfoMap["softMSP"], request)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, record)
}
