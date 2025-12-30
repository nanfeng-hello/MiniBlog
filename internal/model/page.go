package model

type Page struct {
	Total  int         `json:"total"`
	Page   int         `json:"page"`
	Size   int         `json:"size"`
	Record interface{} `json:"record"`
}
