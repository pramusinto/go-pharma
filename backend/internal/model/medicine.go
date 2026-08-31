package model

type Medicine struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
	Unit     string `json:"unit"`
	Price    int    `json:"price"`
}
