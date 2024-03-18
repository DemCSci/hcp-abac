package model

type Policy struct {
	Id                       string   `json:"id"`
	ResourceId               string   `json:"resource_id"`
	RequesterAttributeKeys   []string `json:"requesterAttributeKeys,omitempty"`
	RequesterAttributeValues []string `json:"requester_attribute_values,omitempty"`
	ResourceAttributeKeys    []string `json:"resourceAttributeKeys,omitempty"`
	ResourceAttributeValues  []string `json:"resource_attribute_values,omitempty"`
	PrivateKeys              []string `json:"private_keys,omitempty"`
	PrivateValues            []string `json:"private_values,omitempty"`
}
