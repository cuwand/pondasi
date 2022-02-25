package transactionStatusEnums

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type TransactionStatus string

const (
	Success TransactionStatus = `SUCCESS`
	Pending TransactionStatus = `PENDING`
	Failed  TransactionStatus = `FAILED`
)

func (s *TransactionStatus) String() string {
	switch *s {
	case Success:
		return "SUCCESS"
	case Pending:
		return "PENDING"
	case Failed:
		return "FAILED"
	default:
		return "undefined"
	}
}

func (s *TransactionStatus) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux TransactionStatus
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case Success, Pending, Failed:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for TransactionStatus type, must Success, Pending, Failed")
	}
}
