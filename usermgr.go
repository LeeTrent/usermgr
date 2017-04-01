package usermgr

import (
	"errors"
	"fmt"
)

type User struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

type UserMgr struct {
	userMap map[string]User
}

func NewUserMgr() UserMgr {
	um := UserMgr{
		userMap: make(map[string]User),
	}
	return um
}

func (um UserMgr) GetUser(un string) (User, bool) {
	user, found := um.userMap[un]
	if found {
		return user, found
	} else {
		return User{}, found
	}
}

func (um UserMgr) UserNameIsTaken(un string) bool {
	_, ok := um.userMap[un]
	return ok
}

func (um UserMgr) RemoveUser(userName string) {
	delete(um.userMap, userName)
}

func (um UserMgr) CreateUser(pw []byte, un, fn, ln string) (User, error) {
	if _, ok := um.userMap[un]; ok {
		errMsg := fmt.Sprintf("Username '%s' already taken", un)
		return User{}, errors.New(errMsg)
	}
	um.userMap[un] = User{un, pw, fn, ln}
	return um.userMap[un], nil
}
