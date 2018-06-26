package service

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"jgit.me/tools/notify_gate/db"
	"log"
)

// Service is used for manage services
type Service struct {
	gorm.Model
	Name  string `gorm:"not null;unique"`
	Token string `gorm:"not null;unique"`
}

// Register is used for create token and save to db service model
func Register(srv *Service) (*Service, error) {

	srv.Token = srv.GenerateToken()

	tx := db.Stor.Db().Begin()
	if err := tx.Create(&srv).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return srv, tx.Commit().Error
}

// Unregister is used for remove service from db
func Unregister(srv *Service) error {

	log.Printf("%+v\n", srv)

	if err := db.Stor.Db().Unscoped().Where("token = ?", srv.Token).Delete(Service{}).Error; err != nil {
		return err
	}

	return nil
}

// GetAll is used for receive all services from db
func GetAll() ([]Service, error) {
	srvs := []Service{}
	db.Stor.Db().Find(&srvs)
	return srvs, nil

}

// GenerateToken is used for generating service token
func (u *Service) GenerateToken() string {
	h := sha1.New()
	h.Write([]byte(u.Name))
	return hex.EncodeToString(h.Sum(nil))
}
