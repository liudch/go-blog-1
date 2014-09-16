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

type Bookmarks struct {
	Id      int       `orm:"pk;auto"`
	Name    string    `orm:"unique;size(64)"`
	Count   uint      `orm:"null"`
	Special bool      `orm:"null"`
	MdfTime time.Time `orm:"auto_now;type(datetime)"`
}
