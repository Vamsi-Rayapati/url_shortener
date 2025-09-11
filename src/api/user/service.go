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

func (us *UserService) OnboardUser(userId string, userName string, user OnboardRequest) (*UserResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	userIdParsed, _ := uuid.Parse(userId)
	newUser := database.User{
		ID:        userIdParsed,
		Username:  userName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    database.Active,
	}
	result := db.Create(&newUser)

	if result.Error != nil {
		mysqlErr, ok := result.Error.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return nil, errors.ConfilctError("Username already exists")
		}
		return nil, errors.InternalServerError("Failed to create user")
	}

	return &UserResponse{
		ID:             newUser.ID.String(),
		Username:       newUser.Username,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		FullName:       newUser.FirstName + " " + newUser.LastName,
		PrimaryAddress: newUser.PrimaryAddress,
		Mobile:         newUser.Mobile,
		Role:           newUser.Role,
		Avatar:         newUser.Avatar,
		Status:         newUser.Status,
		CreatedAt:      newUser.CreatedAt.String(),
	}, nil

}

func (us *UserService) AddUser(user CreateUserRequest) (*UserResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	newUser := database.User{
		ID:        uuid.New(),
		Username:  user.Username,
		Mobile:    user.Mobile,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      database.UserRole(user.Role),
	}
	result := db.Create(&newUser)

	if result.Error != nil {
		mysqlErr, ok := result.Error.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return nil, errors.ConfilctError("Username already exists")
		}
		return nil, errors.InternalServerError("Failed to create user")
	}

	return &UserResponse{
		ID:             newUser.ID.String(),
		Username:       newUser.Username,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		FullName:       newUser.FirstName + " " + newUser.LastName,
		PrimaryAddress: newUser.PrimaryAddress,
		Avatar:         newUser.Avatar,
		Mobile:         newUser.Mobile,
		Role:           newUser.Role,
		Status:         newUser.Status,
	}, nil

}

func (us *UserService) GetUsers(request UsersRequest) (*UsersResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var users []database.User
	var total int64

	db.Model(&database.User{}).Count(&total)
	result := db.Order("created_at").Offset(request.PageSize * (request.PageNo - 1)).Limit(request.PageSize).Find(&users)

	if result.Error != nil {
		return nil, errors.InternalServerError("Failed to get users")
	}

	userList := utils.Map(users, func(user database.User) UserResponse {
		return UserResponse{
			ID:             user.ID.String(),
			Username:       user.Username,
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			FullName:       user.FirstName + " " + user.LastName,
			PrimaryAddress: user.PrimaryAddress,
			Mobile:         user.Mobile,
			Role:           user.Role,
			Status:         user.Status,
			CreatedAt:      user.CreatedAt.String(),
		}
	})

	return &UsersResponse{
		Users: userList,
		Total: total,
	}, nil

}

func (us *UserService) GetUser(id string) (*UserResponse, *errors.ApiError) {
	db := dbclient.GetCient()
	var user database.User
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		log.Println("GetUser: %+v", result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("User not found")
		}

		return nil, errors.InternalServerError("Failed to get user")
	}

	return &UserResponse{
		ID:             user.ID.String(),
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		FullName:       user.FirstName + " " + user.LastName,
		PrimaryAddress: user.PrimaryAddress,
		Mobile:         user.Mobile,
		Role:           user.Role,
		Avatar:         user.Avatar,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt.String(),
	}, nil
}

func (us *UserService) UpdateUser(id string, request UpdateUserRequest) (*UserResponse, *errors.ApiError) {
	db := dbclient.GetCient()

	var user database.User

	data, err := utils.StructToMap(request)
	if err != nil {
		return nil, errors.InternalServerError("Failed to read payload")
	}

	log.Println("UpdateUser: %+v", data)
	result := db.Model(database.User{}).Where("id = ?", id).Updates(data).First(&user)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NotFoundError("User not found")
		}

		return nil, errors.InternalServerError("Failed to update user")
	}

	return &UserResponse{
		ID:             user.ID.String(),
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		FullName:       user.FirstName + " " + user.LastName,
		PrimaryAddress: user.PrimaryAddress,
		Mobile:         user.Mobile,
		Role:           user.Role,
		Avatar:         user.Avatar,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt.String(),
	}, nil

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
