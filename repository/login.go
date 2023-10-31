package repository

import (
	"fishfishes_backend/common"
)

const PrivatKey string = "John"

func (r Repo) CheckLogin(user common.User) bool {

	if user.Name == "name" && user.Password == "1234" {
		return true
	}
	return false
}
