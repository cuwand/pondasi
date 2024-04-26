package models

import "fmt"

type Duration struct {
	Seconds float64 `json:"seconds" bson:"seconds"`
	Minutes float64 `json:"minutes" bson:"minutes"`
	Hours   float64 `json:"hours" bson:"hours"`
}

type Distance struct {
	Meters     int     `json:"meters,omitempty" bson:"meters"`
	Kilometers float64 `json:"kilometers,omitempty" bson:"kilometers"`
}

type Coordinate struct {
	Latitude  float64 `bson:"latitude" json:"latitude" form:"latitude" binding:"required"`
	Longitude float64 `bson:"longitude" json:"longitude" form:"longitude" binding:"required"`
}

func (c Coordinate) ToCoordinates() []float64 {
	return []float64{c.Longitude, c.Latitude}
}

func (c Coordinate) ToString() string {
	return fmt.Sprintf("%v,%v", c.Latitude, c.Longitude)
}

func (c Coordinate) ToReversedString() string {
	return fmt.Sprintf("%v,%v", c.Longitude, c.Latitude)
}
