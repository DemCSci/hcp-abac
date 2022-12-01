package com.lei.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author lizhi
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class Attribute {
    /**
     * 私有属性id格式 attribute:resourceId:uuid
     */
    private String id;

    private String type;

    /**
     * 针对那一个资源
     */
    private String resourceId;

    /**
     * 属性发布者
     */
    private String ownerId;

    private double money;

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
