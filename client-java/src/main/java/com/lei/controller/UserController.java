package com.lei.controller;

import com.alibaba.fastjson.JSON;
import com.lei.AttributeTypeEnum;
import com.lei.controller.request.AttributeRequest;
import com.lei.model.Attribute;
import com.lei.model.User;
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

import javax.json.Json;
import java.nio.charset.StandardCharsets;
import java.util.*;
import java.util.concurrent.TimeoutException;


/**
 * @author lizhi
 */
@RestController
@Slf4j
@RequestMapping("/api/v1/user")
public class UserController {


    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @GetMapping("/all")
    public JsonData all() throws ContractException {
        byte[] allUsers = contract.evaluateTransaction("GetAllUsers");

        return JsonData.buildSuccess(JsonUtil.bytes2Obj(allUsers, User[].class));
    }

    @PostMapping("/add")
    public JsonData add() throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("CreateUser")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
        User user = User.builder()
                .money(100)
                .attributes(new ArrayList<>())
                .build();
        byte[] invokeResult = transaction.submit(JSON.toJSONString(user));
        log.info("调用结果:" + new String(invokeResult));
        String transactionId = transaction.getTransactionId();
        return JsonData.buildSuccess(transactionId);

    }
    @DeleteMapping("/del")
    public JsonData del() throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("DeleteUser")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        byte[] invokeResult = transaction.submit();
        log.info("调用结果:" +  new String(invokeResult));
        String transactionId = transaction.getTransactionId();
        return JsonData.buildSuccess(transactionId);
    }
    @GetMapping("/history")
    public JsonData history() throws ContractException {
        byte[] history = contract.evaluateTransaction("GetUserHistory");

        return JsonData.buildSuccess(new String(history));
    }

    @PostMapping("/addAttribute")
    public JsonData  addAttribute(@RequestBody AttributeRequest request) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("AddAttribute")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        Attribute attribute = Attribute.builder()
                .id("attribute:" + UUID.randomUUID().toString())
                .type(AttributeTypeEnum.PUBLIC.name())
                .ownerId(request.getOwnerId())
                .key(request.getKey())
                .value(request.getValue())
                .notBefore(request.getNotBefore())
                .notAfter(request.getNotAfter())
                .build();

        byte[] invokeResult = transaction.submit(JsonUtil.obj2Json(attribute));
        log.info("调用结果:" +  new String(invokeResult));
        String transactionId = transaction.getTransactionId();
        Map<String, String > res = new HashMap(2);
        res.put("txId", transactionId);
        res.put("data", JsonUtil.obj2Json(invokeResult));
        return JsonData.buildSuccess(res);
    }
}
