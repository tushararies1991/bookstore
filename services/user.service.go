package service

import (
	"bookstore/domain/user"
	"bookstore/utils/appUtils"
	"bookstore/utils/error"
)

// "UserService is"
var UserService iUserService = &userService{}

type userService struct {
}

type iUserService interface {
	CreateUser(user.User) (*user.User, *error.AppErr)
	GetUser(int64) (*user.User, *error.AppErr)
	UpdateUser(bool, user.User) (*user.User, *error.AppErr)
	DeleteUser(int64) *error.AppErr
	FindByStatus(string) (user.Users, *error.AppErr)
}

func (us *userService) CreateUser(user user.User) (*user.User, *error.AppErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = "active"
	user.CreatedAt = appUtils.GetNowDBFormat()
	user.Password = appUtils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userService) GetUser(userID int64) (*user.User, *error.AppErr) {
	userInfo := &user.User{Id: userID}

	if err := userInfo.Get(); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (us *userService) UpdateUser(isPartial bool, user user.User) (*user.User, *error.AppErr) {
	crntUsr, err := us.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			crntUsr.FirstName = user.FirstName
		}
		if user.LastName != "" {
			crntUsr.LastName = user.LastName
		}
		if user.Email != "" {
			crntUsr.Email = user.Email
		}
		if user.Phone != "" {
			crntUsr.Phone = user.Phone
		}
	} else {
		crntUsr.FirstName = user.FirstName
		crntUsr.LastName = user.LastName
		crntUsr.Email = user.Email
		crntUsr.Phone = user.Phone
	}
	if err := crntUsr.UpdateUser(); err != nil {
		return nil, err
	}

	return crntUsr, nil
}

func (us *userService) DeleteUser(userID int64) *error.AppErr {
	user := &user.User{Id: userID}
	return user.Delete()
}

func (us *userService) FindByStatus(status string) (user.Users, *error.AppErr) {
	user := &user.User{Status: status}
	return user.FindByStatus()
}
