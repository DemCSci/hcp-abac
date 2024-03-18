package model

type BuyAttributeRequest struct {
	Buyer       string `json:"buyer"`
	Seller      string `json:"seller"`
	AttributeId string `json:"attributeId"`
}
type Attribute struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	ResourceId string  `json:"resourceId"`
	Owner      string  `json:"ownerId"`
	Key        string  `json:"key"`
	Value      string  `json:"value"`
	NotBefore  string  `json:"notBefore"`
	NotAfter   string  `json:"notAfter"`
	Money      float64 `json:"money"`
}
