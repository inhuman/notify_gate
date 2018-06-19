package user

import (
	"jgit.me/tools/notify_gate/db"
)

type User struct {
	Name  string
	Token string
}

func Register(user User) (*User, error) {

	db.Stor.Db()
	tx := db.Stor.Db().Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error
}
