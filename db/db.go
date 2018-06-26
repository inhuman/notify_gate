package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // exporting postgres dialect
	"jgit.me/tools/notify_gate/config"
	"time"
)

type storage struct {
	db *gorm.DB
}

// Stor is used for access db
var Stor storage

// Init is used for initialize and connect to db
func Init() {
	Stor = storage{}
	Stor.Db()
	fmt.Println("Db connected")
}

func (s *storage) Connect() error {
	if s.db == nil {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				config.AppConf.Postgres.User,
				config.AppConf.Postgres.Password,
				config.AppConf.Postgres.Host,
				config.AppConf.Postgres.Port,
				config.AppConf.Postgres.DbName,
			))
		if err != nil {
			return err
		}

		//db.LogMode(true)
		s.db = db
	}
	err := s.db.DB().Ping()
	if err != nil {
		s.db = nil
		s.Connect()
	}
	return nil
}

func (s *storage) Db() *gorm.DB {
	err := s.Connect()

	if err != nil {
		fmt.Println("Lost db connection. Reconnecting..")
	}

	if s.db == nil {
		<-time.After(5 * time.Second)
		return s.Db()
	}

	return s.db
}

func (s *storage) Close() {
	s.Db().Close()
}

func (s *storage) Migrate(object interface{}) {
	s.Db().AutoMigrate(object)
}

func (s *storage) SetDb(db *gorm.DB) {
	s.db = db
}
