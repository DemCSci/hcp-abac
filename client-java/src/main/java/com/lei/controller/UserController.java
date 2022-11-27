package com.lei.controller;

import com.lei.util.JsonData;
import lombok.extern.slf4j.Slf4j;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.Network;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;



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
    public JsonData users() {
        return JsonData.buildSuccess();
    }
}
