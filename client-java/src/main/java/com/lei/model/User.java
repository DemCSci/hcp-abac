package com.lei.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

/**
 * @author lizhi
 */
@NoArgsConstructor
@AllArgsConstructor
@Data
@Builder
public class User {
    private String id;
    private double money;
    private List<Attribute> attributes;
}
