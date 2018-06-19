package user

import (
	"jgit.me/tools/notify_gate/db"
	"github.com/jinzhu/gorm"
	"crypto/sha1"
	"encoding/hex"
)

type Service struct {
	gorm.Model
	Name  string `gorm:"not null;unique"`
	Token string `gorm:"not null;unique"`
}

func Register(srv *Service) (*Service, error) {

	srv.Token = srv.GenerateToken()

	db.Stor.Db()
	tx := db.Stor.Db().Begin()
	if err := tx.Create(&srv).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return srv, tx.Commit().Error
}

func GetAll() ([]Service, error) {

	srvs := []Service{}
	db.Stor.Db().Find(&srvs)
	return srvs, nil

}

func (u *Service) GenerateToken() string {
	h := sha1.New()
	h.Write([]byte(u.Name))
	return hex.EncodeToString(h.Sum(nil))
}
