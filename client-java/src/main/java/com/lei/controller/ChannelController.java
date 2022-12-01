package com.lei.controller;

import com.alibaba.fastjson.JSON;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.protos.peer.Query;
import org.hyperledger.fabric.sdk.BlockInfo;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.TransactionInfo;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicReference;

@RestController
@Slf4j
@RequestMapping("/api/v1/channel")
public class ChannelController {
    @Autowired
    private Channel channel;

    /**
     * 根据交易id查询区块
     * @param txId
     * @return
     * @throws InvalidArgumentException
     * @throws ProposalException
     */
    @GetMapping("/queryBlockByTransactionID")
    public JsonData queryBlockByTransactionID(String txId) throws InvalidArgumentException, ProposalException {
        BlockInfo blockInfo = channel.queryBlockByTransactionID(txId);
        return JsonData.buildSuccess(JSON.toJSONString(blockInfo));
    }

    @GetMapping("/peers")
    public JsonData getPeers() {
        Collection<Peer> peers = channel.getPeers();
        log.info("peers: {}", peers);
        String peersJson = JsonUtil.obj2Json(peers);
        return JsonData.buildSuccess(JsonUtil.json2List(peersJson, Map.class));
    }

    /**
     * 获取通道的名字
     * @return
     */
    @GetMapping("/name")
    public JsonData getName() {
        String name = channel.getName();
        return JsonData.buildSuccess(name);

    }

    @GetMapping("/queryTransactionByID")
    public JsonData queryTransactionByID(String txId) throws InvalidArgumentException, ProposalException {
        TransactionInfo transactionInfo = channel.queryTransactionByID(txId);
        log.info("transactionInfo: {}", JSON.toJSONString(transactionInfo));
        String jsonString = JSON.toJSONString(transactionInfo);
        return JsonData.buildSuccess(JsonUtil.json2Obj(jsonString, Map.class));
    }

    /**
     * 获取所有链码的名字
     * @return
     */
    @GetMapping("/getChainCodeNames")
    public JsonData getChainCodeNames() {
        Collection<String> names = channel.getDiscoveredChaincodeNames();
        return JsonData.buildSuccess(names);
    }

    @GetMapping("/instantiatedChaincodes")
    public JsonData test() throws InvalidArgumentException, ProposalException {
        Collection<Peer> peers = channel.getPeers();
        AtomicReference<Peer> peer = new AtomicReference<>();
        peers.stream().forEach(obj -> {
            if (obj.getName().equalsIgnoreCase("peer0.org1.example.com")) {
                peer.set(obj);
            }
        });
        log.info("peers:{}",peers);
        List<Query.ChaincodeInfo> chaincodeInfos = channel.queryInstantiatedChaincodes(peer.get());
        log.info("已实例化的链码: {}", chaincodeInfos);

        return JsonData.buildSuccess(JsonUtil.obj2Json(chaincodeInfos));
    }
}
