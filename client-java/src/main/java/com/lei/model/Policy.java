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
public class Policy {
    private String id;

    private String resourceId;

    private String ownerId;

    private String content;
}
