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
