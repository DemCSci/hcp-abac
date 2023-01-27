package com.lei.model;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
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
@ApiModel("策略实体类")
public class Policy {
    @ApiModelProperty("策略id")
    private String id;

    @ApiModelProperty("资源id")
    private String resourceId;

    @ApiModelProperty("策略拥有者")
    private String ownerId;

    @ApiModelProperty("策略内容")
    private String content;
}
