package com.lei.controller;

import com.lei.model.User;
import com.lei.request.DecideRequest;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import io.swagger.annotations.Api;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Transaction;
import org.hyperledger.fabric.sdk.Peer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.EnumSet;
import java.util.concurrent.TimeoutException;

@RestController
@RequestMapping("/api/decide/v1")
@Api("决策相关")
public class DecideController {
    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @PostMapping("/decideNoRecord")
    public JsonData decideNoRecord(@RequestBody DecideRequest decideRequest) throws ContractException {
        Transaction transaction = contract.createTransaction("Decide");
        DecideRequest request = DecideRequest.builder()
                .id(transaction.getTransactionId())
                .requesterId(decideRequest.getRequesterId())
                .resourceId(decideRequest.getResourceId())
                .build();
        byte[] result = transaction.evaluate(JsonUtil.obj2Json(request));

        return JsonData.buildSuccess(new String(result));
    }

    @PostMapping("/decideWithRecord")
    public JsonData decideWithRecord(@RequestBody DecideRequest decideRequest) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("DecideWithRecord")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
        DecideRequest request = DecideRequest.builder()
                .id(transaction.getTransactionId())
                .requesterId(decideRequest.getRequesterId())
                .resourceId(decideRequest.getResourceId())
                .build();
        byte[] result = transaction.submit(JsonUtil.obj2Json(request));

        return JsonData.buildSuccess(new String(result));
    }
}
