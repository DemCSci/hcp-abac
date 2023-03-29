package com.lei.controller.request;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author lizhi
 */
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class BuyPrivateAttributeRequest {
    private String buyer;

    private String seller;

    private String attributeId;
}
