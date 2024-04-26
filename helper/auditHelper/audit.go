package auditHelper

import (
	"github.com/cuwand/pondasi/helper/dateHelper"
	"github.com/cuwand/pondasi/helper/idHelper"
	"github.com/cuwand/pondasi/models"
)

func ToAudit(user models.User) models.Audit {
	currentTime := dateHelper.TimeNow()

	return models.Audit{
		Id:          idHelper.UUID(),
		CreatedBy:   &user,
		CreatedDate: currentTime,
		UpdatedBy:   &user,
		UpdatedDate: currentTime,
		Version:     0,
		Delete:      false,
	}
}

func ToAuditCustomId(id string, user models.User) models.Audit {
	currentTime := dateHelper.TimeNow()

	return models.Audit{
		Id:          id,
		CreatedBy:   &user,
		CreatedDate: currentTime,
		UpdatedBy:   &user,
		UpdatedDate: currentTime,
		Version:     0,
		Delete:      false,
	}
}
