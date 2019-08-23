package service

import (
	"github.com/ltebes/GoTebes/src/domain"
)

var userManager *UserManager

type UserManager struct {
	newUser  *domain.User
	newUsers []*domain.User
}

func (u *UserManager) Register(user *domain.User) error {
	u.newUser = user
	u.newUsers = append(u.newUsers, user)
	return nil
	// Aca tengo que agregar una posicion del array de maps
}

func (u *UserManager) GetUser() *domain.User {
	return u.newUser
}

func GetInstance() *UserManager {
	if userManager == nil {
		userManager = new(UserManager)
	}

	return userManager
}
