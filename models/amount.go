package models

type Amount struct {
	Value     float64 `json:"value"`
	Rupiah    string  `json:"rupiah"`
	Terbilang string  `json:"terbilang"`
}

type TagAmount struct {
	Tag    string `json:"tag"`
	Amount Amount `json:"amount"`
}
