package com.lei.controller.request;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

/**
 * @author lei
 * @since 2023-03-03
 */

@Data
@ToString
@NoArgsConstructor
@AllArgsConstructor
public class AddResourceControllerRequest {
    private String resourceId;
    private String controllerId;
}
