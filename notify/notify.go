package notify

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"jgit.me/tools/notify_gate/db"
)

// Notify is used for managing notifies
type Notify struct {
	gorm.Model
	Type    string         `json:"type"`
	Message string         `json:"message"`
	UIDs    pq.StringArray `json:"uids" gorm:"type:varchar(64)[]"`
}

// Save is used for saving notify to db
func (n *Notify) Save() {
	db.Stor.Db().Save(n)
}

// Execute is used for implement WorkerPool.Task
func (n *Notify) Execute() {
	n.Save()
}

// Delete is used for deleting notify from db
func (n *Notify) Delete() {
	db.Stor.Db().Unscoped().Delete(n)
}

// GetNotify is used for receive one notify from db
func GetNotify() *Notify {
	ns := &Notify{}
	db.Stor.Db().First(ns)
	return ns
}
