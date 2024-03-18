package model

type Policy struct {
	Id                       string   `json:"id"`
	ResourceId               string   `json:"resource_id"`
	RequesterAttributeKeys   []string `json:"requesterAttributeKeys"`
	RequesterAttributeValues []string `json:"requester_attribute_values"`
	ResourceAttributeKeys    []string `json:"resourceAttributeKeys"`
	ResourceAttributeValues  []string `json:"resource_attribute_values"`
	PrivateKeys              []string `json:"private_keys"`
	PrivateValues            []string `json:"private_values"`
}
