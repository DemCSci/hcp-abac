package com.lei.request;

import lombok.*;

/**
 * @author lizhi
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@ToString
public class DecideRequest {
    private String id;

    private String requesterId;

    private String resourceId;


}
