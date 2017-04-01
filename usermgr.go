package usermgr

import (
	"errors"
	"fmt"
)

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
}
type UserMgr map[string]User

func (um UserMgr) GetUser(un string) User {
	return um[un]
}

func (um UserMgr) UserNameIsTaken(un string) bool {
	_, ok := um[un]
	return ok
}

func (um UserMgr) RemoveUser(un string) {
	delete(um, un)
}

func (um UserMgr) CreateUser(pw []byte, un, fn, ln string) (User, error) {
	if _, ok := um[un]; ok {
		errMsg := fmt.Sprintf("Username '%s' already taken", un)
		return User{}, errors.New(errMsg)
	}
	um[un] = User{un, pw, fn, ln}
	return um[un], nil
}
