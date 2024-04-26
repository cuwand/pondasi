package models

import (
	"github.com/cuwand/pondasi/helper/dateHelper/dateProperty"
)

type TransactionRequest struct {
	CustomerReferenceNumber string                `json:"customer_reference_number" form:"customer_reference_number" binding:"required"`
	ClientDateTime          dateProperty.DateTime `json:"client_date_time" form:"client_date_time" binding:"required"`
}
