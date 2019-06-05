package models

import (
	"time"
)

type UserDay struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Uid        int       `xorm:"INT(11)"`
	Day        string    `xorm:"VARCHAR(8)"`
	Num        int       `xorm:"INT(11)"`
	SysCreated time.Time `xorm:"DATETIME"`
	SysUpdated time.Time `xorm:"DATETIME"`
}
