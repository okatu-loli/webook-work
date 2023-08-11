package domain

import (
	"errors"
	"time"
)

// User 领域对象，是 DDD 中的 entity
// BO(business object)
type User struct {
	Id       int64
	Email    string
	Nickname *string
	Birthday *time.Time
	Bio      *string
	Password string
	Ctime    time.Time
}

func (u *User) Validate() error {
	if u.Nickname != nil && (len(*u.Nickname) < 3 || len(*u.Nickname) > 20) {
		return errors.New("昵称必须在3到20个字符之间")
	}
	if u.Bio != nil && len(*u.Bio) > 200 {
		return errors.New("个人简介必须少于200个字符")
	}
	return nil
}

//type Address struct {
//}
