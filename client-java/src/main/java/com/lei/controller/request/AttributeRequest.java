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
public class AttributeRequest {

    private String id;

    private String type;

    private String resourceId;
    /**
     * 属性拥有者
     */
    private String ownerId;
    /**
     * 属性价格
     */
    private double money;
    /**
     * 属性的key
     */
    private String key;
    private String value;
    /**
     * 生效时间
     */
    private String notBefore;
    /**
     * 失效时间
     */
    private String notAfter;

}
