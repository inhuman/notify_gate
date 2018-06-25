package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"jgit.me/tools/notify_gate/config"
	"time"
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
	fmt.Println("Inializing db")

	Stor = Storage{
		Host:     config.AppConf.Postgres.Host,
		Port:     config.AppConf.Postgres.Port,
		DbName:   config.AppConf.Postgres.DbName,
		User:     config.AppConf.Postgres.User,
		Password: config.AppConf.Postgres.Password,
	}

	//TODO: remove
	fmt.Printf("%+v\n", Stor)

	Stor.Db()
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
		fmt.Println("Lost db connection. Reconnecting..")
	}

	if s.db == nil {
		<- time.After(5 * time.Second)
		return s.Db()
	}

	return s.db
}

func (s *Storage) Close() {
	s.Db().Close()
}

func (s *Storage) Migrate(object interface{}) {
	s.Db().AutoMigrate(object)
}

func (s *Storage) SetDb(db *gorm.DB) {
	s.db = db
}