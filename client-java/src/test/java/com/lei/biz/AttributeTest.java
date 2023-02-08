package com.lei.biz;



import com.lei.controller.request.AttributeRequest;
import com.lei.controller.request.BuyPrivateAttributeRequest;
import com.lei.enums.AttributeTypeEnum;
import com.lei.model.Attribute;

import com.lei.util.JsonUtil;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Transaction;
import org.hyperledger.fabric.sdk.Peer;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;



import java.util.EnumSet;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;


/**
 * @author lei
 * @since 2023/2/7
 */

@SpringBootTest
@Slf4j
public class AttributeTest {

    @Autowired
    private Contract contract;

    @Autowired
    private Network network;

    @Test
    public void testAddPublicAttribute() throws ContractException {

        long startTime  = System.currentTimeMillis();
        for (int i = 0; i < 1000; i++) {
            Transaction transaction = contract.createTransaction("AddAttribute")
                    .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
            Attribute attribute = Attribute.builder()
                    .id("attribute:" + UUID.randomUUID())
                    .type(AttributeTypeEnum.PUBLIC.name())
                    .ownerId("")
                    .key("name")
                    .value("张三")
                    .notBefore("1669791474807")
                    .notAfter("1672383443000")
                    .build();

            byte[] invokeResult = transaction.evaluate(JsonUtil.obj2Json(attribute));
            //log.info("调用结果:" +  new String(invokeResult));
            String transactionId = transaction.getTransactionId();
            Map<String, String > res = new HashMap(2);
            res.put("txId", transactionId);
            res.put("data", JsonUtil.obj2Json(invokeResult));
            //log.info(JsonUtil.obj2Json(res));
        }
        long endTime = System.currentTimeMillis();
        log.error("执行时间：{}ms",endTime-startTime); //6158ms
    }

    @Test
    public void testAddPrivateAttribute() throws ContractException {
        long startTime  = System.currentTimeMillis();
        for (int i = 0; i < 1000; i++) {
            Transaction transaction = contract.createTransaction("BuyPrivateAttribute")
                    .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));
            BuyPrivateAttributeRequest request = new BuyPrivateAttributeRequest();
            request.setAttributeId("attribute:private:resource:571971ca-932f-4e39-bc8d-475778f44401:occupation0");
            request.setBuyer("user:x509::CN=Admin@org1.example.com,OU=admin,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US");
            request.setSeller("user:x509::CN=Admin@org1.example.com,OU=admin,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US");

            byte[] invokeResult = transaction.evaluate(JsonUtil.obj2Json(request));
            Map<String, Object> map = new HashMap<>(2);
            map.put("txId", transaction.getTransactionId());
            map.put("data", new String(invokeResult));
            //log.info(JsonUtil.obj2Json(map));
        }
        long endTime = System.currentTimeMillis();
        log.error("执行时间：{}ms",endTime-startTime); // 6504ms
    }

    @Test
    public void testPublishPrivateAttribute() throws ContractException {
        long startTime  = System.currentTimeMillis();
        for (int i = 0; i < 1000; i++) {
            Transaction transaction = contract.createTransaction("PublishPrivateAttribute")
                    .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)));

            AttributeRequest attributeRequest = new AttributeRequest();
            attributeRequest.setType("PRIVATE");
            attributeRequest.setResourceId("resource:571971ca-932f-4e39-bc8d-475778f44401");
            attributeRequest.setOwnerId("user:x509::CN=Admin@org1.example.com,OU=admin,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US");
            attributeRequest.setKey("occupation0");
            attributeRequest.setValue("doctor");
            attributeRequest.setMoney(0);
            attributeRequest.setNotBefore("1669791474807");
            attributeRequest.setNotAfter("1672383443000");
            byte[] result = transaction.evaluate(JsonUtil.obj2Json(attributeRequest));

            Map<String, Object> map = new HashMap<>(2);
            map.put("txId", transaction.getTransactionId());
            // 里面应该是 属性id
            map.put("data", new String(result));
            //log.info(JsonUtil.obj2Json(map));
        }
        long endTime = System.currentTimeMillis();
        log.error("执行时间：{}ms",endTime-startTime); // 5195ms
    }

    //@Autowired
    //private AttributeController attributeController;
    //
    //@Test
    //public void publishAttributeTest() throws ContractException, InterruptedException, TimeoutException, IOException, JSONException {
    //    FileWriter fstream  = new FileWriter("C:\\Users\\lizhi\\Desktop\\attribute.csv",false);
    //
    //    BufferedWriter out = new BufferedWriter(fstream);
    //
    //    //out.write(jsonobj.getJSONObject("data").getString("data")+"\n");
    //
    //    ExecutorService executorService = Executors.newFixedThreadPool(10);
    //
    //    for (int i = 0; i < 10000; i++) {
    //
    //        Future<?> submit = executorService.submit(() -> {
    //            try {
    //                exeThread(out);
    //            } catch (ContractException e) {
    //                throw new RuntimeException(e);
    //            } catch (InterruptedException e) {
    //                throw new RuntimeException(e);
    //            } catch (TimeoutException e) {
    //                throw new RuntimeException(e);
    //            } catch (IOException e) {
    //                throw new RuntimeException(e);
    //            }
    //        });
    //
    //    }
    //    out.flush();
    //    TimeUnit.SECONDS.sleep(60);
    //    executorService.shutdown();
    //    out.close();
    //    fstream.close();
    //}
    //
    //public void exeThread(BufferedWriter out) throws ContractException, InterruptedException, TimeoutException, IOException {
    //    AttributeRequest attributeRequest = new AttributeRequest();
    //    attributeRequest.setType("PRIVATE");
    //    attributeRequest.setResourceId("resource:8b395393-c556-4273-9ee3-3c158ecde223");
    //    attributeRequest.setOwnerId("user:x509::CN=Admin@org1.example.com,OU=admin,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US");
    //    attributeRequest.setKey("occupation-"+ UUID.randomUUID().toString());
    //    attributeRequest.setValue("doctor");
    //    attributeRequest.setMoney(0);
    //    attributeRequest.setNotBefore("1669791474807");
    //    attributeRequest.setNotAfter("1672383443000");
    //    JsonData jsonData = attributeController.publish(attributeRequest);
    //    System.out.println(jsonData);
    //    JSONObject data = jsonData.getData(new TypeReference<JSONObject>() {
    //    });
    //    String attributeId = data.getString("data");
    //    synchronized (AttributeTest.class) {
    //        out.write(attributeId+"\n");
    //        out.flush();
    //    }
    //}
}
