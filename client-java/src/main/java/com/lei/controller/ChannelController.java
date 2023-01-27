package com.lei.controller;

import com.alibaba.fastjson2.JSON;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
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
import java.util.Properties;
import java.util.concurrent.atomic.AtomicReference;

@RestController
@Slf4j
@RequestMapping("/api/v1/channel")
@Api(tags = "与通道有关的操作")
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
    @ApiOperation("通过交易id查询区块")
    public JsonData queryBlockByTransactionID(String txId) throws InvalidArgumentException, ProposalException {
        BlockInfo blockInfo = channel.queryBlockByTransactionID(txId);

        return JsonData.buildSuccess(JSON.toJSONString(blockInfo));
    }

    @ApiOperation("获取peer节点")
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
    @ApiOperation("获取通道的name")
    @GetMapping("/name")
    public JsonData getName() {
        String name = channel.getName();
        return JsonData.buildSuccess(name);
    }

    @ApiOperation("根据id查询交易")
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
    @ApiOperation("获取所有链码的名字")
    @GetMapping("/getChainCodeNames")
    public JsonData getChainCodeNames() {
        Collection<String> names = channel.getDiscoveredChaincodeNames();
        return JsonData.buildSuccess(names);
    }

    @GetMapping("/instantiatedChaincodes")
    @ApiOperation("查询已经实例化的链码")
    public JsonData instantiatedChaincodes() throws InvalidArgumentException, ProposalException {
        Collection<Peer> org1MSP = channel.getPeersForOrganization("Org1MSP");
        List<Query.ChaincodeInfo> chaincodeInfos = channel.queryInstantiatedChaincodes(org1MSP.iterator().next());
        log.info("已实例化的链码: {}", chaincodeInfos);
        return JsonData.buildSuccess(JsonUtil.obj2Json(chaincodeInfos));
    }

    @GetMapping("/getServiceDiscoveryProperties")
    @ApiOperation("获取服务发现的properties")
    public JsonData getServiceDiscoveryProperties() {
        Properties serviceDiscoveryProperties = channel.getServiceDiscoveryProperties();
        return JsonData.buildSuccess(serviceDiscoveryProperties);
    }
}
