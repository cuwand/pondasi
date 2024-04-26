package mapperHelper

import (
	"encoding/json"
	"fmt"
	"github.com/cuwand/pondasi/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func InterfaceMongoToStruct(from interface{}, to interface{}) error {
	marshaledDoc, _ := bson.Marshal(from)

	if err := bson.Unmarshal(marshaledDoc, to); err != nil {
		return errors.InternalServerError("Cannot parsing string to structs")
	}

	return nil
}

func InterfaceToStruct(from interface{}, to interface{}) error {
	marshaledDoc, _ := json.Marshal(from)

	if err := json.Unmarshal(marshaledDoc, to); err != nil {
		return errors.InternalServerError("Cannot parsing string to structs")
	}

	return nil
}

func JsonStringToStruct(jsonString string, to interface{}) error {
	if err := json.Unmarshal([]byte(jsonString), to); err != nil {
		return errors.InternalServerError(fmt.Sprintf("Cannot parsing string to structs | %v", jsonString))
	}

	return nil
}
