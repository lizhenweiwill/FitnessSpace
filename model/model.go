package model

import "time"

// Coach 教练
type Coach struct {
	ID        uint   `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string `json:"name" gorm:"not null;unique;comment:教练名称"`
	Nick      string `json:"nick" gorm:"not null;unique;comment:教练昵称"`
	Pass      string `json:"-" gorm:"not null;unique;comment:登陆密码"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Course 课程`gorm:"index"`
type Course struct {
	ID        uint   `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string `json:"name" gorm:"not null;unique;comment:课程名称"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Member 会员
type Member struct {
	ID        uint   `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string `json:"name" gorm:"not null;unique;comment:会员名称"`
	Card      string `json:"card" gorm:"not null;unique;uniqueIndex;comment:身份证号"`
	Phone     string `json:"phone" gorm:"not null;unique;uniqueIndex;comment:手机号码"`
	WeiXin    string `json:"wx" gorm:"not null;unique;comment:微信号码"`
	Image     string `json:"image" gorm:"not null;unique;comment:头像地址"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
