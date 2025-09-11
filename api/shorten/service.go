package user

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/smartbot/account/database"
	"github.com/smartbot/account/pkg/dbclient"
	"github.com/smartbot/account/pkg/errors"
	"github.com/smartbot/account/pkg/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

func (us *UserService) DeleteUser(userId string) *errors.ApiError {
	db := dbclient.GetCient()
	result := db.Where("id = ?", userId).Delete(&database.User{})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.NotFoundError("User not found")
		}
		return errors.InternalServerError("Failed to get users")
	}

	return nil
}
