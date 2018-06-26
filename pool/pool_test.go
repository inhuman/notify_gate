package pool

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"log"
	"testing"
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
	gormDB.LogMode(true)

	return gormDB.Set("gorm:update_column", true), mck
}

func endExpect(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// test with mock db, not work, because insert and select doing async
//func TestNotifyPool_AddMock(t *testing.T) {
//
//	dbm, mock := getMock(t)
//	defer dbm.Close()
//
//	n := 2
//	tt := time.Now()
//
//	i := 0
//
//	// Expect inserts
//	for i = 0; i < n; i++ {
//		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type", "message", "uids"}).
//			AddRow(i, tt, tt, nil, "test", "test message", nil)
//
//		mock.ExpectQuery(`INSERT INTO "notifies" (.+)`).
//			WithArgs(tt, tt, nil,"test", "test message", nil).
//			WillReturnRows(rows)
//	}
//
//	// Expects select
//	for i = 0; i < n; i++ {
//		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "type", "message", "uids"}).
//			AddRow(i, tt, tt, nil, "test", "test message", nil)
//
//		mock.ExpectQuery(`SELECT * FROM "notifies" (.+)`).
//			WithArgs(tt, tt, nil,"test", "test message", nil).
//			WillReturnRows(rows)
//	}
//
//	err := config.AppConf.Load()
//
//	if err != nil {
//		t.Log(err)
//	}
//
//	db.Stor.SetDb(dbm)
//
//	wpool := workerpool.NewPool(5)
//
//	go Saver(wpool)
//	go Sender()
//
//
//	senders.Providers["test"] = func(n *notify.Notify) error {
//		time.Sleep(500 * time.Millisecond)
//		fmt.Println("sent message " + n.Message)
//		return nil
//	}
//
//	for i := 0; i < n; i++ {
//		n := &notify.Notify{
//			Type:    "test",
//			Message: "test message",
//		}
//		n.CreatedAt = tt
//		n.UpdatedAt = tt
//
//		nPool.AddToSave(n)
//	}
//	<- time.After(1 * time.Second)
//
//	for {
//		n := notify.GetNotify()
//		if n.ID == 0 {
//			nPool.Done <- true
//			break
//		}
//		<- time.After(1 * time.Second)
//	}
//	endExpect(t, mock)
//}

// test with real db, works fine
//func TestNotifyPool_Add(t *testing.T) {
//
//	err := config.AppConf.Load()
//
//	if err != nil {
//		t.Log(err)
//	}
//
//	db.Init()
//	db.Stor.Db()
//	db.Stor.Migrate(notify.Notify{})
//
//
//	wpool := workerpool.NewPool(5)
//
//	go Saver(wpool)
//	go Sender()
//
//
//	senders.Providers["test"] = func(n *notify.Notify) error {
//		time.Sleep(500 * time.Millisecond)
//		fmt.Println("sent message " + n.Message)
//		return nil
//	}
//
//	for i := 0; i < 100; i++ {
//		n := &notify.Notify{
//			Type:    "test",
//			Message: "test message " + strconv.Itoa(i),
//		}
//		AddToSave(n)
//	}
//	<- time.After(1 * time.Second)
//
//	for {
//		n := notify.GetNotify()
//		if n.ID == 0 {
//			nPool.Done <- true
//			break
//		}
//		<- time.After(1 * time.Second)
//	}
//}
