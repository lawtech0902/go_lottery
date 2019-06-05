package models

import (
	"time"
)

type Gift struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Title        string    `xorm:"VARCHAR(255)"`
	PrizeNum     int       `xorm:"INT(11)"`
	LeftNum      int       `xorm:"INT(11)"`
	PrizeCode    string    `xorm:"VARCHAR(50)"`
	PrizeTime    int       `xorm:"INT(11)"`
	Img          string    `xorm:"VARCHAR(255)"`
	DisplayOrder int       `xorm:"INT(11)"`
	Gtype        int       `xorm:"INT(11)"`
	Gdata        string    `xorm:"VARCHAR(255)"`
	TimeBegin    time.Time `xorm:"DATETIME"`
	TimeEnd      time.Time `xorm:"DATETIME"`
	PrizeData    string    `xorm:"MEDIUMTEXT"`
	PrizeBegin   time.Time `xorm:"DATETIME"`
	PrizeEnd     time.Time `xorm:"DATETIME"`
	SysStatus    int       `xorm:"SMALLINT(6)"`
	SysCreated   time.Time `xorm:"DATETIME"`
	SysUpdated   time.Time `xorm:"DATETIME"`
	SysIp        string    `xorm:"VARCHAR(50)"`
}
