package com.lei.controller;

import com.alibaba.fastjson2.JSON;
import com.alibaba.fastjson2.JSONObject;
import com.google.protobuf.Descriptors;
import com.google.protobuf.InvalidProtocolBufferException;
import com.lei.config.GatewayConfig;
import com.lei.util.JsonData;
import com.lei.util.JsonUtil;
import com.lei.vo.HeightInfo;
import com.lei.vo.KVWrite;
import com.lei.vo.TransactionActionInfo;
import com.lei.vo.TransactionEnvelopeInfo;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.codec.DecoderException;
import org.apache.commons.codec.binary.Hex;
import org.apache.commons.io.HexDump;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.protos.ledger.rwset.kvrwset.KvRwset;
import org.hyperledger.fabric.protos.peer.Query;
import org.hyperledger.fabric.protos.peer.TransactionPackage;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.nio.charset.Charset;
import java.util.*;
import java.util.concurrent.atomic.AtomicReference;

@RestController
@Slf4j
@RequestMapping("/api/v1/channel")
@Api(tags = "与通道有关的操作")
public class ChannelController {
    @Autowired
    private Channel channel;

    @Autowired
    private GatewayConfig gatewayConfig;

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

        return JsonData.buildSuccess();
        //log.info("transactionInfo: {}", JSON.toJSONString(transactionInfo));
        ////transactionInfo.getValidationCode()
        //String jsonString = JSON.toJSONString(transactionInfo);
        //return JsonData.buildSuccess(JsonUtil.json2Obj(jsonString, Map.class));
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


    @GetMapping("/height")
    @ApiOperation("获取区块高度")
    public JsonData getHeight() throws InvalidArgumentException, ProposalException {
        BlockchainInfo blockchainInfo = channel.queryBlockchainInfo();
        long height = blockchainInfo.getHeight();

        String currentBlockHash = Hex.encodeHexString(blockchainInfo.getCurrentBlockHash());
        String previousBlockHash = Hex.encodeHexString(blockchainInfo.getPreviousBlockHash());
        HeightInfo heightInfo = HeightInfo.builder().height(height).currentBlockHash(currentBlockHash).previousBlockHash(previousBlockHash).build();
        return JsonData.buildSuccess(heightInfo);
    }

    @GetMapping("/queryBlockByHash")
    @ApiOperation("根据hash查询区块")
    public JsonData queryBlockByHash(String hash) throws DecoderException, InvalidArgumentException, ProposalException, InvalidProtocolBufferException {
        BlockInfo blockInfo = channel.queryBlockByHash(Hex.decodeHex(hash));

        String channelId = blockInfo.getChannelId();
        long blockNumber = blockInfo.getBlockNumber();
        String dataHash = Hex.encodeHexString(blockInfo.getDataHash());
        String previousHash = Hex.encodeHexString(blockInfo.getPreviousHash());
        int transactionCount = blockInfo.getTransactionCount();

        //log.info("channelId={},blockNumber={},dataHash={},previousHash={},transactionCount={}",channelId,
        //        blockNumber,dataHash,previousHash, transactionCount);
        int envelopeCount = blockInfo.getEnvelopeCount();
        //log.info("envelopeCount={}",envelopeCount);
        Iterable<BlockInfo.EnvelopeInfo> envelopeInfoIterable = blockInfo.getEnvelopeInfos();

        List<TransactionEnvelopeInfo> transactions = new ArrayList<>();

        for (BlockInfo.EnvelopeInfo envelopeInfo : envelopeInfoIterable) {
            String envelopeInfoChannelId = envelopeInfo.getChannelId();
            String type = envelopeInfo.getType().name();
            BlockInfo.EnvelopeInfo.IdentitiesInfo creator = envelopeInfo.getCreator();
            String creatorId = creator.getId();
            String creatorMspid = creator.getMspid();
            String nonce = Hex.encodeHexString(envelopeInfo.getNonce());
            Date timestamp = envelopeInfo.getTimestamp();
            String transactionID = envelopeInfo.getTransactionID();
            byte validationCode = envelopeInfo.getValidationCode();
            String validation = TransactionPackage.TxValidationCode.forNumber(validationCode).name();
            //log.info("envelopeInfoChannelId={},type={},creatorId={},creatorMspid={},nonce={},timestamp={},transactionID={},validation={}",
            //        envelopeInfoChannelId, type,creatorId,creatorMspid, nonce, timestamp, transactionID, validation);


            // 强转类型
            BlockInfo.TransactionEnvelopeInfo transactionEnvelopeInfo = (BlockInfo.TransactionEnvelopeInfo) envelopeInfo;
            List<TransactionActionInfo> transactionActionInfos = new ArrayList<>();
            // 获取操作事务的信息
            for (BlockInfo.TransactionEnvelopeInfo.TransactionActionInfo transactionActionInfo : transactionEnvelopeInfo.getTransactionActionInfos()) {

                // 链码名称
                String chaincodeIDName = transactionActionInfo.getChaincodeIDName();
                // 链码版本
                String chaincodeIDVersion = transactionActionInfo.getChaincodeIDVersion();
                //log.info("proposal chaincodeIDName:{}, chaincodeIDVersion: {}", chaincodeIDName, chaincodeIDVersion);

                // 操作事务的读写集
                TxReadWriteSetInfo rwsetInfo = transactionActionInfo.getTxReadWriteSet();
                List<KVWrite> kvWriteList = new ArrayList<>();
                if (null != rwsetInfo) {

                    //我只要了写集合的数据
                    for (TxReadWriteSetInfo.NsRwsetInfo nsRwsetInfo : rwsetInfo.getNsRwsetInfos()) {
                        // 含有默认链码 _lifecycle  和 自定义链码
                        String namespace = nsRwsetInfo.getNamespace();
                        // 只要符合要求的链码的
                        if (!namespace.equals(gatewayConfig.getContractName())) {
                            //log.error("链码名称不正确 跳过 ,namespace ={}", namespace);
                            continue;
                        }
                        KvRwset.KVRWSet rws = nsRwsetInfo.getRwset();
                        for (KvRwset.KVWrite writeList : rws.getWritesList()) {
                            //String valAsString = printableString(new String(writeList.getValue().toByteArray(), Charset.forName("UTF-8")));
                            String key = writeList.getKey();
                            //log.info("Namespace {}  key {} has value '{}' ", namespace, writeList.getKey(), valAsString);
                            String value = new String(writeList.getValue().toByteArray(), Charset.forName("UTF-8"));
                            //if (StringUtils.isNotBlank(valAsString)) {
                            //   value = JSON.parseObject(valAsString, Map.class);
                            //}
                            KVWrite kvWrite = new KVWrite();
                            kvWrite.setKey(key);
                            kvWrite.setValue(value);
                            kvWriteList.add(kvWrite);
                        }
                    }
                }
                TransactionActionInfo actionInfo = TransactionActionInfo.builder().chaincodeIDName(chaincodeIDName).chaincodeIDVersion(chaincodeIDVersion)
                        .writeList(kvWriteList).build();
                transactionActionInfos.add(actionInfo);
            }
            TransactionEnvelopeInfo envelopeInfo1 = TransactionEnvelopeInfo.builder().channelId(channelId).creatorId(creatorId).creatorMspid(creatorMspid)
                    .nonce(nonce).timestamp(timestamp).transactionID(transactionID).validation(validation)
                    .transactionActionInfos(transactionActionInfos).build();
            transactions.add(envelopeInfo1);

        }

        com.lei.vo.BlockInfo info = com.lei.vo.BlockInfo.builder().channelId(channelId)
                .blockNumber(blockNumber)
                .dataHash(dataHash)
                .previousHash(previousHash)
                .transactionCount(transactionCount)
                .envelopeCount(envelopeCount)
                .transactions(transactions).build();

        return JsonData.buildSuccess(info);
    }
    static String printableString(final String string) {
        int maxLogStringLength = 64;
        if (string == null || string.length() == 0) {
            return string;
        }

        String ret = string.replaceAll("[^\\p{Print}]", "?");

        ret = ret.substring(0, Math.min(ret.length(), maxLogStringLength)) + (ret.length() > maxLogStringLength ? "..." : "");

        return ret;

    }
}
