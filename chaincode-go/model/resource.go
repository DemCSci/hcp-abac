package model

type Resource struct {
	Id          string `json:"id"`
	Owner       string `json:"owner"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
