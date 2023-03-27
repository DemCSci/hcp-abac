package com.lei.biz;

import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;

import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;


import java.io.IOException;
import java.io.UnsupportedEncodingException;

/**
 * @author lei
 * @since 2023-03-25
 */

@SpringBootTest
public class InitTest {
    @Autowired
    private CloseableHttpClient closeableHttpClient;

    public static String goServerRootUrl = "http://127.0.0.1:7788";

    public static String javaServerRootUrl = "http://127.0.0.1:9001";
    /**
     * 注册所有用户
     */
    @Test
    public void registerAllUser() throws IOException {
        HttpGet httpGet = new HttpGet(goServerRootUrl + "/addAllUser");
        CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpGet);
        System.out.println(httpResponse.getEntity().toString());
        httpResponse.close();
        System.out.println("注册所有用户成功");
    }

    @Test
    public void addPublicAttribute() throws UnsupportedEncodingException {
        String json = "{\n" +
                "\n" +
                "    \"id\": \"xxx\",\n" +
                "    \"type\": \"PUBLIC\",\n" +
                "    \"resourceId\": \"\",\n" +
                "    \"ownerId\": \"\",\n" +
                "    \"key\": \"age\",\n" +
                "    \"value\": \"40\",\n" +
                "    \"money\": 50,\n" +
                "    \"notBefore\": \"1669791474807\",\n" +
                "    \"notAfter\": \"1672383443000\"\n" +
                "}";
        StringEntity stringEntity = new StringEntity(json);

        HttpPost httpPost = new HttpPost(javaServerRootUrl + "/api/attribute/v1/addAttribute");
        httpPost.setHeader("Content-Type", "application/json");
        httpPost.setEntity(stringEntity);

        try (CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost)) {
            System.out.println(httpResponse.getEntity().toString());
            System.out.println("增加公有属性成功");
       }catch (Exception e) {

       }
    }

    @Test
    public void addResource() throws UnsupportedEncodingException {
        String json = "{\n" +
                "\n" +
                "    \"id\": \"xxx\",\n" +
                "    \"owner\": \"xxx\",\n" +
                "    \"url\": \"http://www.baidu.com\",\n" +
                "    \"description\": \"访问百度1\"\n" +
                "}";
        StringEntity stringEntity = new StringEntity(json);

        HttpPost httpPost = new HttpPost(javaServerRootUrl + "/api/v1/resource/create");
        httpPost.setHeader("Content-Type", "application/json");
        httpPost.setEntity(stringEntity);

        try (CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost)) {
            System.out.println(httpResponse.getEntity().toString());
            System.out.println("增加资源成功");
        }catch (Exception e) {

        }
    }

    @Test
    public void publishPrivateAttribute() throws UnsupportedEncodingException {
        //需要更改resourceId 和 ownerId
        String json = "{\n" +
                "\n" +
                "    \"id\": \"xxx\",\n" +
                "    \"type\": \"PRIVATE\",\n" +
                "    \"resourceId\": \"resource:7b76c72a-6a1f-40c0-b6f5-7df42dcca349\",\n" +
                "    \"ownerId\": \"user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e\",\n" +
                "    \"key\": \"occupation\",\n" +
                "    \"value\": \"doctor\",\n" +
                "    \"money\": 50,\n" +
                "    \"notBefore\": \"1669791474807\",\n" +
                "    \"notAfter\": \"1672383443000\"\n" +
                "}";
        StringEntity stringEntity = new StringEntity(json);

        HttpPost httpPost = new HttpPost(javaServerRootUrl + "/api/attribute/v1/publish");
        httpPost.setHeader("Content-Type", "application/json");
        httpPost.setEntity(stringEntity);

        try (CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost)) {
            System.out.println(httpResponse.getEntity().toString());
            System.out.println("发布私有属性成功");
        }catch (Exception e) {

        }
    }
    @Test
    public void addPrivateAttribute() throws UnsupportedEncodingException {
        //需要更改resourceId 和 ownerId
        String json = "{\n" +
                "\n" +
                "    \"buyer\": \"user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e\",\n" +
                "    \"seller\": \"user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e\",\n" +
                "    \"attributeId\": \"attribute:private:resource:7b76c72a-6a1f-40c0-b6f5-7df42dcca349:occupation\"\n" +
                "}";
        StringEntity stringEntity = new StringEntity(json);

        HttpPost httpPost = new HttpPost(javaServerRootUrl + "/api/attribute/v1/buy");
        httpPost.setHeader("Content-Type", "application/json");
        httpPost.setEntity(stringEntity);

        try (CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost)) {
            System.out.println(httpResponse.getEntity().toString());
            System.out.println("增加私有属性成功");
        }catch (Exception e) {

        }
    }

    @Test
    public void addResourceControllers() throws UnsupportedEncodingException {
        //需要更改resourceId 和 ownerId
        String json = "{\n" +
                "\n" +
                "    \"buyer\": \"user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e\",\n" +
                "    \"seller\": \"user:654455774f546f365130343964584e6c636a457354315539593278705a5735304c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a3156557a6f36513034396332396d644335705a6d46756447467a655335755a58517354315539526d4669636d6c6a4c45383953486c775a584a735a57526e5a58497355315139546d3979644767675132467962327870626d4573517a315655773d3dd41d8cd98f00b204e9800998ecf8427e\",\n" +
                "    \"attributeId\": \"attribute:private:resource:7b76c72a-6a1f-40c0-b6f5-7df42dcca349:occupation\"\n" +
                "}";
        StringEntity stringEntity = new StringEntity(json);

        HttpPost httpPost = new HttpPost(javaServerRootUrl + "/api/attribute/v1/buy");
        httpPost.setHeader("Content-Type", "application/json");
        httpPost.setEntity(stringEntity);

        try (CloseableHttpResponse httpResponse = closeableHttpClient.execute(httpPost)) {
            System.out.println(httpResponse.getEntity().toString());
            System.out.println("增加私有属性成功");
        }catch (Exception e) {

        }
    }
}
