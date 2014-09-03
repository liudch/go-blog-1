package db

import (
	"time"
)

type User struct {
	Id        int       `orm:"pk;auto"`
	Nickname  string    `orm:"size(32)"`
	Email     string    `orm:"unique;size(255)"`
	Password  string    `orm:"size(255)"`
	Aphorism  string    `orm:"null;size(255)"`
	AvatarPth string    `orm:"null;size(255)"`
	Admin     bool      `orm:"null"`
	LgiTime   time.Time `orm:"null;type(datetime)"`
	GenTime   time.Time `orm:"auto_now_add;type(datetime)"`
	MdfTime   time.Time `orm:"auto_now;type(datetime)"`
}
