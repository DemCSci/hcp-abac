package com.lei.controller;


import com.alibaba.fastjson2.JSON;
import com.lei.controller.request.ResourceRequest;
import com.lei.model.Resource;
import com.lei.model.User;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import io.swagger.annotations.Api;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Transaction;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.nio.channels.Channel;
import java.util.ArrayList;
import java.util.EnumSet;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.TimeoutException;

/**
 * @author lizhi
 */
@Slf4j
@RestController
@RequestMapping("/api/v1/resource")
@Api(tags = "资源相关")
public class ResourceController {

    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @PostMapping("/create")
    public JsonData createResource(@RequestBody ResourceRequest request) throws ContractException, InterruptedException, TimeoutException {
        request.setId("resource:"+ UUID.randomUUID());
        request.setOwner("");
        Transaction transaction = contract.createTransaction("CreateResource")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
        byte[] invokeResult = transaction.submit(JSON.toJSONString(request));
        log.info("调用结果:" + new String(invokeResult));
        String transactionId = transaction.getTransactionId();
        return JsonData.buildSuccess(transactionId);
    }

    @GetMapping("/all")
    public JsonData createResource() throws ContractException, InterruptedException, TimeoutException {

        byte[] result = contract.evaluateTransaction("GetAllResource");

        log.info("调用结果:" + new String(result));

        return JsonData.buildSuccess(JsonUtil.bytes2Obj(result, Resource[].class));
    }
}
