package models

type Asset struct {
	Url  string `bson:"url" json:"url"`
	Name string `bson:"name" json:"name"`
}
