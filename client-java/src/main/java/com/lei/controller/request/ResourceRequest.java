package com.lei.controller.request;

import lombok.Data;

/**
 * @author lizhi
 */
@Data
public class ResourceRequest {
    private String id;

    private String owner;

    private String url;

    private String description;
}
