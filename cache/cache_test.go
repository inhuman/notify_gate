package cache

import (
	"testing"
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"log"
	"time"
	"jgit.me/tools/notify_gate/db"
	"github.com/stretchr/testify/assert"
)

func getMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	dbm, mck, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, gerr := gorm.Open("postgres", dbm)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}

	return gormDB.Set("gorm:update_column", true), mck
}

func endExpect(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestBuildTokenCache(t *testing.T) {

	dbm, mock := getMock(t)
	defer dbm.Close()

	db.Stor.SetDb(dbm)

	tt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "token"}).
		AddRow(1, tt, tt, nil, "test_service", "test_token")

	mock.ExpectQuery(`SELECT (.+) FROM "services" (.+)`).
		WillReturnRows(rows)

	BuildTokenCache()
	endExpect(t, mock)

	assert.Equal(t, "test_service", TokensCache["test_token"])

}
