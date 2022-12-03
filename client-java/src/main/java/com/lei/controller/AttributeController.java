package com.lei.controller;

import com.lei.controller.request.AttributeRequest;
import com.lei.controller.request.BuyPrivateAttributeRequest;
import com.lei.model.Attribute;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Transaction;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.*;
import java.util.concurrent.TimeoutException;

/**
 * @author lizhi
 */
@RestController
@Slf4j
@RequestMapping("/api/attribute/v1")
public class AttributeController {

    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @PostMapping("/publish")
    public JsonData publish(@RequestBody AttributeRequest request) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("PublishPrivateAttribute")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        request.setId("attribute:" + request.getResourceId() + ":" + UUID.randomUUID().toString());
        byte[] result = transaction.submit(JsonUtil.obj2Json(request));

        Map<String, Object> map = new HashMap<>(2);
        map.put("txId", transaction.getTransactionId());
        map.put("data", new String(result));
        return JsonData.buildSuccess(map);
    }

    @GetMapping("/findByResourceId")
    public JsonData find(String resourceId) throws ContractException {
        byte[] attributes = contract.evaluateTransaction("FindAttributeByResourceId", resourceId);

        return JsonData.buildSuccess(JsonUtil.bytes2Obj(attributes, Attribute[].class));
    }

    @PostMapping("/buy")
    public JsonData buy(@RequestBody BuyPrivateAttributeRequest request) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("BuyPrivateAttribute")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        byte[] result = transaction.submit(JsonUtil.obj2Json(request));

        Map<String, Object> map = new HashMap<>(2);
        map.put("txId", transaction.getTransactionId());
        map.put("data", new String(result));
        return JsonData.buildSuccess(map);
    }
}
