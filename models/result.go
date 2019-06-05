package models

import (
	"time"
)

type Result struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	GiftId     int       `xorm:"index INT(11)"`
	GiftName   string    `xorm:"VARCHAR(250)"`
	GiftType   int       `xorm:"INT(11)"`
	Uid        int       `xorm:"INT(11)"`
	Username   string    `xorm:"VARCHAR(50)"`
	PrizeCode  int       `xorm:"INT(11)"`
	GiftData   string    `xorm:"VARCHAR(50)"`
	SysCreated time.Time `xorm:"DATETIME"`
	SysIp      string    `xorm:"VARCHAR(50)"`
	SysStatus  int       `xorm:"SMALLINT(6)"`
}
