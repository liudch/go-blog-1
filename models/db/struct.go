package db

import (
	"time"
)

type Accounts struct {
	Id       int       `orm:"pk;auto"`
	Nickname string    `orm:"unique;size(32)"`
	Email    string    `orm:"unique;size(160)"`
	Password string    `orm:"size(32)"` //MD5
	Motto    string    `orm:"null;size(255)"`
	Admin    bool      `orm:"null"`
	GenTime  time.Time `orm:"auto_now_add;type(datetime)"`
	MdfTime  time.Time `orm:"auto_now;type(datetime)"`
}
