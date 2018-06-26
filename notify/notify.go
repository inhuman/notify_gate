package notify

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"jgit.me/tools/notify_gate/db"
)

type Notify struct {
	gorm.Model
	Type    string         `json:"type"`
	Message string         `json:"message"`
	UIDs    pq.StringArray `json:"uids" gorm:"type:varchar(64)[]"`
}

func (n *Notify) Save() {
	db.Stor.Db().Save(n)
}

func (n *Notify) Execute() {
	n.Save()
}

func (n *Notify) Delete() {
	db.Stor.Db().Unscoped().Delete(n)
}

func GetNotify() *Notify {
	ns := &Notify{}
	db.Stor.Db().First(ns)
	return ns
}
