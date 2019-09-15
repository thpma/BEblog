package models

import (
	"time"
)

type Articles struct {
	Id					int				`xorm:"INT(11) pk autoincr unique"`
	Title				string		`xorm:"VARCHAR(255)"`
	Author			string		`xorm:"VARCHAR(50)"`
	Content			string		`xorm:"LONGTEXT"`
	Poster			string		`xorm:"VARCHAR(255)"`
	Category		string		`xorm:"VARCHAR(255)"`
	Comments		int				`xorm:"INT(11)"`
	Views				int				`xorm:"INT(11)"`
	CreatedAt		time.Time	`xorm:"TIMESTAMP"`
	UpdatedAt		time.Time	`xorm:"TIMESTAMP"`
	DeletedAt		time.Time	`xorm:"TIMESTAMP"`
}