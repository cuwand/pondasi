package models

type Client struct {
	Id         string `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	ClientIp   string `bson:"client_ip" json:"client_ip"`
	ClientCode string `bson:"client_code" json:"client_code"`
}
