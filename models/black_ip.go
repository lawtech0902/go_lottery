package models

import (
	"time"
)

type BlackIp struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Ip         string    `xorm:"VARCHAR(50)"`
	BlackTime  time.Time `xorm:"DATETIME"`
	SysCreated time.Time `xorm:"DATETIME"`
	SysUpdated time.Time `xorm:"DATETIME"`
}
