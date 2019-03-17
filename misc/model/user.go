package model

type User struct {
	Id        int    `xorm:"not null pk autoincr unique INT(11)"`
	UserId    string `xorm:"not null default '' unique VARCHAR(20)"`
	CompanyId string `xorm:"not null default '' unique VARCHAR(20)"`
	Phone     string `xorm:"default '' VARCHAR(11)"`
	Passwd    string `xorm:"default '' VARCHAR(50)"`
	NickName  string `xorm:"default '' VARCHAR(30)"`
	Gender int `xorm:"default 0 comment('0: 男
1: 女
2: 保密') INT(11)"`
	WxOpenid   string `xorm:"default '' VARCHAR(50)"`
	WxUnionid  string `xorm:"default '' VARCHAR(50)"`
	Address    string `xorm:"default '' VARCHAR(100)"`
	Del        int    `xorm:"default 0 TINYINT(1)"`
	CreateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
