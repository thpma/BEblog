package models

type Users struct {
	Id 					int			`xorm:"INT(11) pk autoincr unique"`
	Name				string	`xorm:"VARCHAR(50)"`
	Password		string	`xorm:"VARCHAR(255)"`
	Token				string	`xorm:"VARCHAR(255)"`
	ExpireTime	int64		`xorm:"INT(11)"`
}