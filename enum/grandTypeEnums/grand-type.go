package grandTypeEnums

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type GrandType string

const (
	IMPLICIT           GrandType = `implicit`
	AUTHORIZATION_CODE GrandType = `authorization_code`
	CLIENT_CREDENTIAL  GrandType = `client_credentials`
	PASSWORD           GrandType = `password`
	REFRESH_TOKEN      GrandType = `refresh_token`
)

func FromString(str string) GrandType {
	switch str {
	case "implicit":
		return IMPLICIT
	case "authorization_code":
		return AUTHORIZATION_CODE
	case "client_credentials":
		return CLIENT_CREDENTIAL
	case "password":
		return PASSWORD
	case "refresh_token":
		return REFRESH_TOKEN
	default:
		return "undefined"
	}
}

func (s *GrandType) String() string {
	switch *s {
	case IMPLICIT:
		return "implicit"
	case AUTHORIZATION_CODE:
		return "authorization_code"
	case CLIENT_CREDENTIAL:
		return "client_credentials"
	case PASSWORD:
		return "password"
	case REFRESH_TOKEN:
		return "refresh_token"
	default:
		return "undefined"
	}
}

func (s *GrandType) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux GrandType
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case IMPLICIT, AUTHORIZATION_CODE, CLIENT_CREDENTIAL, PASSWORD, REFRESH_TOKEN:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for IMPLICIT, AUTHORIZATION_CODE, CLIENT_CREDENTIAL, PASSWORD, REFRESH_TOKEN")
	}
}
