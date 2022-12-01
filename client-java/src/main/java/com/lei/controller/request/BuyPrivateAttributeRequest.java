package com.lei.controller.request;

import lombok.Data;

/**
 * @author lizhi
 */
@Data
public class BuyPrivateAttributeRequest {
    private String buyer;

    private String seller;

    private String attributeId;
}
