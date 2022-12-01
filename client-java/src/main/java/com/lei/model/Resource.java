package com.lei.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author lizhi
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Resource {
    private String id;

    private String owner;

    private String url;

    private String description;

}
