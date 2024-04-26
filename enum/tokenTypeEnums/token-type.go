package tokenTypeEnums

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type TokenType string

const (
	BEARER TokenType = `Bearer`
)

func (s *TokenType) String() string {
	switch *s {
	case BEARER:
		return "bearer"
	default:
		return "undefined"
	}
}

func (s *TokenType) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux TokenType
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case BEARER:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for BEARER")
	}
}
