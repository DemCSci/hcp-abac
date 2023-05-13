package com.lei.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * @author lei
 * @since 2023-05-13
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Record {
    private String Id;

    private String requesterId;

    private String resourceId;

    private String response;
}
