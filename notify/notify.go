package notify

import (
	"github.com/inhuman/notify_gate/db"
	"github.com/jinzhu/gorm"
	"strings"
)

// Notify is used for managing notifies
type Notify struct {
	gorm.Model
	Type    string   `json:"type"`
	Message string   `json:"message"`
	UIDs    []string `json:"uids" gorm:"-" sql:"-"`
	UIDsStr string
}

// Save is used for saving notify to db
func (n *Notify) Save() {

	if len(n.UIDs) > 0 {
		n.UIDsStr = strings.Join(n.UIDs, ";")
	}

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
	ns.UIDs = strings.Split(ns.UIDsStr, ";")
	return ns
}
