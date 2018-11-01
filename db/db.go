package db

import (
	"fmt"
	"github.com/inhuman/notify_gate/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // exporting postgres dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
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
	log.Println("Db connected")
}

func (s *storage) Connect() error {

	if s.db == nil {
		switch config.AppConf.DB.Type {
		case "postgres":

			//TODO: check config values if postgres
			db, err := gorm.Open("postgres",
				fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
					config.AppConf.DB.User,
					config.AppConf.DB.Password,
					config.AppConf.DB.Host,
					config.AppConf.DB.Port,
					config.AppConf.DB.Name,
				))
			if err != nil {
				return err
			}

			//db.LogMode(true)
			s.db = db
		case "sqlite3":

			db, err := gorm.Open("sqlite3", "notifies.db")
			if err != nil {
				return err
			}

			//db.LogMode(true)
			s.db = db

		default:
			fmt.Println("Database type:", config.AppConf.DB.Type, "not supported")
			fmt.Println("Available types: postgres, sqlite3")
			os.Exit(1)
		}
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
		log.Println("Lost db connection. Reconnecting..")
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

func (s *storage) Migrate(object interface{}) error {
	return s.Db().AutoMigrate(object).Error
}

func (s *storage) SetDb(db *gorm.DB) {
	s.db = db
}
