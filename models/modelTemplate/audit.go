package modelTemplate

import (
	"github.com/cuwand/pondasi/helper/dateHelper"
	"github.com/cuwand/pondasi/helper/idHelper"
	"github.com/cuwand/pondasi/models"
)

func ToAuditSystem() models.Audit {
	timeNow := dateHelper.TimeNow()
	user := ToUserSystem()

	return models.Audit{
		Id:          idHelper.UUID(),
		CreatedBy:   &user,
		CreatedDate: timeNow,
		UpdatedBy:   &user,
		UpdatedDate: timeNow,
		Version:     1,
		Delete:      false,
	}
}

func ToAudit(user models.User) models.Audit {
	timeNow := dateHelper.TimeNow()

	return models.Audit{
		Id:          idHelper.UUID(),
		CreatedBy:   &user,
		CreatedDate: timeNow,
		UpdatedBy:   &user,
		UpdatedDate: timeNow,
		Version:     1,
		Delete:      false,
	}
}

func ToUserSystem() models.User {
	return models.User{
		Identity: "system",
		Username: "system",
		FullName: "system",
	}
}
