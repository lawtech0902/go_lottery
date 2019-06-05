package models

import (
	"time"
)

type Code struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	GiftId     int       `xorm:"index INT(11)"`
	Code       string    `xorm:"VARCHAR(255)"`
	SysCreated time.Time `xorm:"DATETIME"`
	SysUpdated time.Time `xorm:"DATETIME"`
	SysStatus  int       `xorm:"SMALLINT(6)"`
}
