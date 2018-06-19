package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"jgit.me/tools/notify_gate/config"
)

type Storage struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	db       *gorm.DB
}

var Stor Storage

func Init() {
	Stor = Storage{
		Host:     config.AppConf.Postgre.Host,
		Port:     config.AppConf.Postgre.Port,
		DbName:   config.AppConf.Postgre.DbName,
		User:     config.AppConf.Postgre.User,
		Password: config.AppConf.Postgre.Password,
	}

}

func (s *Storage) Connect() error {
	if s.db == nil {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				s.User, s.Password, s.Host, s.Port, s.DbName))
		if err != nil {
			return err
		}

		s.db = db
	}
	err := s.db.DB().Ping()
	if err != nil {
		s.db = nil
		s.Connect()
	}
	return nil
}

func (s *Storage) Db() *gorm.DB {
	err := s.Connect()
	if err != nil {
		panic(err)
	}

	return s.db
}

func (s *Storage) Close() {
	s.Db().Close()
}

func (s *Storage) Migrate(object interface{}) {
	s.Db().AutoMigrate(object)
}