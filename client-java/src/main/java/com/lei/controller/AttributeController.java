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
import org.hyperledger.fabric.gateway.*;
import org.hyperledger.fabric.gateway.impl.ContractImpl;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
import org.hyperledger.fabric.gateway.spi.CommitHandler;
import org.hyperledger.fabric.gateway.spi.CommitHandlerFactory;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.json.Json;
import java.util.*;
import java.util.concurrent.TimeUnit;
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

    @Autowired
    private Channel channel;

    @Autowired
    private Gateway gateway;

    // 经过客户端验证的 peer 的背书响应
    Collection<ProposalResponse> validProposalResponses;


    @GetMapping("/attributes")
    @ApiOperation("根据用户id查找该用户的所有属性")
    public JsonData attributes(@RequestParam("user_id") String userId) throws ContractException {
        byte[] attributes = contract.evaluateTransaction("FindAttributeByUserId", userId);
        if (attributes == null || attributes.length == 0) {
            return JsonData.buildError("属性为null 或者长度为0");
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



    // 尝试只进行背书，不提交到orderer节点
    @PostMapping("/addAttributeOnlyPeer")
    @ApiOperation("增加属性(只进行背书)")
    public JsonData  addAttributeOnlyPeer(@RequestBody AttributeRequest request) throws ContractException, InterruptedException, TimeoutException, InvalidArgumentException, ProposalException {

        TransactionProposalRequest transactionProposalRequest = network.getGateway().getClient().newTransactionProposalRequest();
        transactionProposalRequest.setChaincodeName("abac");
        transactionProposalRequest.setChaincodeLanguage(TransactionRequest.Type.GO_LANG);
        transactionProposalRequest.setFcn("AddAttribute");
        Attribute attribute = Attribute.builder()
                .id("attribute:" + UUID.randomUUID())
                .type(AttributeTypeEnum.PUBLIC.name())
                .ownerId(request.getOwnerId())
                .key(request.getKey())
                .value(request.getValue())
                .notBefore(request.getNotBefore())
                .notAfter(request.getNotAfter())
                .build();

        transactionProposalRequest.setArgs(JsonUtil.obj2Json(attribute));
        Collection<ProposalResponse> proposalResponses = channel.sendTransactionProposal(transactionProposalRequest,
                network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
        // 自己验证一下
        validProposalResponses = validatePeerResponses(proposalResponses);

        Map<String, String > res = new HashMap(2);
        res.put("txId", proposalResponses.toString());
        return JsonData.buildSuccess(res);
    }

    // 验证 peer 节点背书的是否正确
    private Collection<ProposalResponse> validatePeerResponses(final Collection<ProposalResponse> proposalResponses)
            throws ContractException {
        final Collection<ProposalResponse> validResponses = new ArrayList<>();
        final Collection<String> invalidResponseMsgs = new ArrayList<>();
        proposalResponses.forEach(response -> {
            String peerUrl = response.getPeer() != null ? response.getPeer().getUrl() : "<unknown>";
            if (response.getStatus().equals(ChaincodeResponse.Status.SUCCESS)) {
                log.debug(String.format("validatePeerResponses: valid response from peer %s", peerUrl));
                validResponses.add(response);
            } else {
                log.warn(String.format("validatePeerResponses: invalid response from peer %s, message %s", peerUrl, response.getMessage()));
                invalidResponseMsgs.add(response.getMessage());
            }
        });

        if (validResponses.size() < 1) {
            String msg = String.format("No valid proposal responses received. %d peer error responses: %s",
                    invalidResponseMsgs.size(), String.join("; ", invalidResponseMsgs));
            log.error(msg);
            throw new ContractException(msg, proposalResponses);
        }

        return validResponses;
    }


    // 只提交到orderer节点
    @PostMapping("/addAttributeOnlyOrder")
    @ApiOperation("增加属性(只发送到orderer节点)")
    public JsonData  addAttributeOnlyOrder() throws ContractException, InterruptedException, TimeoutException, InvalidArgumentException, ProposalException {

        for (int i = 0; i < 100; i++) {
            TimeUnit.MILLISECONDS.sleep(300);
            new Thread(new Runnable() {
                @Override
                public void run() {
                    Channel.TransactionOptions transactionOptions = Channel.TransactionOptions.createTransactionOptions()
                            .nOfEvents(Channel.NOfEvents.createNoEvents()); // Disable default commit wait behaviour
                    channel.sendTransaction(validProposalResponses, transactionOptions);
                }
            }).start();
        }


//        ProposalResponse proposalResponse = validProposalResponses.iterator().next();
//        GatewayImpl gatewayImpl = (GatewayImpl)gateway;
//        CommitHandlerFactory commitHandlerFactory = gatewayImpl.getCommitHandlerFactory();
//        CommitHandler commitHandler = commitHandlerFactory.create(proposalResponse.getTransactionID(), network);
//        commitHandler.startListening();
//
//        try {
//            Channel.TransactionOptions transactionOptions = Channel.TransactionOptions.createTransactionOptions()
//                    .nOfEvents(Channel.NOfEvents.createNoEvents()); // Disable default commit wait behaviour
//            channel.sendTransaction(validProposalResponses, transactionOptions)
//                    .get(5, TimeUnit.MINUTES);
//        } catch (TimeoutException e) {
//            commitHandler.cancelListening();
//            throw e;
//        } catch (Exception e) {
//            commitHandler.cancelListening();
//            throw new ContractException("Failed to send transaction to the orderer", e);
//        }
//
//        commitHandler.waitForEvents(5, TimeUnit.MINUTES);

        Map<String, String > res = new HashMap(2);
        //res.put("txId", proposalResponse.getTransactionID());
        return JsonData.buildSuccess(res);
    }

    @DeleteMapping("/clear")
    @ApiOperation("清空该用户公有属性")
    public JsonData clearPublicAttribute() {

        return null;
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
