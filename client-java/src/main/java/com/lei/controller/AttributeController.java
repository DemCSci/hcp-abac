package com.lei.controller;

import com.lei.controller.request.AttributeRequest;
import com.lei.controller.request.BuyPrivateAttributeRequest;
import com.lei.enums.AttributeTypeEnum;
import com.lei.model.Attribute;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
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
@Api(tags = "属性操作相关接口")
@RequestMapping("/api/attribute/v1")
public class AttributeController {

    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @GetMapping("/attributes")
    @ApiOperation("根据用户id查找该用户的所有属性")
    public JsonData attributes(@RequestParam("user_id") String userId) throws ContractException {
        byte[] attributes = contract.evaluateTransaction("FindAttributeByUserId", userId);
        if (attributes == null) {
            return JsonData.buildError("属性为null");
        }
        return JsonData.buildSuccess(JsonUtil.bytes2Obj(attributes, Attribute[].class));
    }

    @PostMapping("/addAttribute")
    public JsonData  addAttribute(@RequestBody AttributeRequest request) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("AddAttribute")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        Attribute attribute = Attribute.builder()
                .id("attribute:" + UUID.randomUUID())
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

    /**
     * 发布私有属性
     * @param request
     * @return
     * @throws ContractException
     * @throws InterruptedException
     * @throws TimeoutException
     */
    @PostMapping("/publish")
    @ApiOperation("发布私有属性")
    public JsonData publish(@RequestBody AttributeRequest request) throws ContractException, InterruptedException, TimeoutException {
        Transaction transaction = contract.createTransaction("PublishPrivateAttribute")
                .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

        request.setId("attribute:" + request.getResourceId() + ":" + UUID.randomUUID());
        byte[] result = transaction.submit(JsonUtil.obj2Json(request));

        Map<String, Object> map = new HashMap<>(2);
        map.put("txId", transaction.getTransactionId());
        // 里面应该是 属性id
        map.put("data", new String(result));
        return JsonData.buildSuccess(map);
    }

    /**
     * 根据资源id查找属性
     * @param resourceId
     * @return
     * @throws ContractException
     */
    @GetMapping("/findByResourceId")
    @ApiOperation("根据资源id查找属性")
    public JsonData find(String resourceId) throws ContractException {
        byte[] attributes = contract.evaluateTransaction("FindAttributeByResourceId", resourceId);

        return JsonData.buildSuccess(JsonUtil.bytes2Obj(attributes, Attribute[].class));
    }

    /**
     * 购买私有属性
     * @param request
     * @return
     * @throws ContractException
     * @throws InterruptedException
     * @throws TimeoutException
     */
    @PostMapping("/buy")
    @ApiOperation("购买私有属性")
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
