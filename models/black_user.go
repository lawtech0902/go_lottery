package models

import (
	"time"
)

type BlackUser struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"VARCHAR(50)"`
	BlackTime  time.Time `xorm:"DATETIME"`
	RealName   string    `xorm:"VARCHAR(50)"`
	Mobile     string    `xorm:"VARCHAR(50)"`
	Address    string    `xorm:"VARCHAR(255)"`
	SysCreated time.Time `xorm:"DATETIME"`
	SysUpdated time.Time `xorm:"DATETIME"`
	SysIp      string    `xorm:"VARCHAR(50)"`
}
