package models

type Config struct {
	Rank  Price `json:"Rank"`
	Price Price `json:"Price"`
}

type Price struct {
	Max int `json:"max"`
}
