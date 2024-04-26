package transactionFlowEnums

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type TransactionFlow string

const (
	Debit    TransactionFlow = `DEBIT`
	Credit   TransactionFlow = `CREDIT`
	Reversal TransactionFlow = `REVERSAL`
)

func (s *TransactionFlow) String() string {
	switch *s {
	case Debit:
		return "DEBIT"
	case Credit:
		return "CREDIT"
	case Reversal:
		return "REVERSAL"
	default:
		return "undefined"
	}
}

func (s *TransactionFlow) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux TransactionFlow
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case Debit, Credit, Reversal:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for TransactionStatus type, must Debit, Credit, Reversal")
	}
}
