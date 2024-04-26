package models

import (
	"time"
)

type Audit struct {
	Id          string    `bson:"_id,omitempty" json:"id"`
	CreatedBy   *User     `bson:"created_by,omitempty" json:"created_by"`
	CreatedDate time.Time `bson:"created_date,omitempty" json:"created_date"`
	UpdatedBy   *User     `bson:"updated_by,omitempty" json:"updated_by"`
	UpdatedDate time.Time `bson:"updated_date,omitempty" json:"updated_date"`
	Version     uint64    `bson:"version,omitempty" json:"version"`
	Delete      bool      `bson:"delete" json:"delete"`
}
