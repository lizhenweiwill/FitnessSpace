package model

import (
	"gorm.io/gorm"
	"time"
)

// Coach 教练
// 钩子方法 BeforeSave 被调用
// 钩子方法 BeforeCreate 被调用
// 钩子方法 AfterSave 被调用
// 钩子方法 AfterCreate 被调用
type Coach struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string    `json:"name" gorm:"not null;unique;comment:教练名称"`
	Nick      string    `json:"nick" gorm:"not null;unique;comment:教练昵称"`
	Pass      string    `json:"-" gorm:"not null;unique;comment:登陆密码"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (c *Coach) BeforeCreate(tx *gorm.DB) (err error) {
	println("钩子方法 BeforeCreate 被调用")
	return
}

func (c *Coach) BeforeSave(tx *gorm.DB) (err error) {
	println("钩子方法 BeforeSave 被调用")
	return
}

func (c *Coach) AfterCreate(tx *gorm.DB) (err error) {
	println("钩子方法 AfterCreate 被调用")
	return
}

func (c *Coach) AfterSave(tx *gorm.DB) (err error) {
	println("钩子方法 AfterSave 被调用")
	return
}

// Course 课程
type Course struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string    `json:"name" gorm:"not null;unique;comment:课程名称"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Member 会员
type Member struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	Name      string    `json:"name" gorm:"not null;unique;comment:会员名称"`
	Card      string    `json:"card" gorm:"not null;unique;uniqueIndex;comment:身份证号"`
	Phone     string    `json:"phone" gorm:"not null;unique;uniqueIndex;comment:手机号码"`
	WeiXin    string    `json:"wx" gorm:"not null;unique;comment:微信号码"`
	Image     string    `json:"image" gorm:"not null;unique;comment:头像地址"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Record 统计记录
//【哪位会员】的【什么课程】还剩【多少节】
type Record struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	MemberId  uint      `json:"mId" gorm:"not null;comment:会员编号"`
	CourseId  uint      `json:"cId" gorm:"not null;comment:课程编号"`
	Count     uint      `json:"count" gorm:"not null;comment:统计信息"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// BuyRecord 购买记录
//【哪位会员】在【什么时间】购买【什么课程】【多少节】实际付款【多少元】
type BuyRecord struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	MemberId  uint      `json:"mId" gorm:"not null;comment:会员编号"`
	CourseId  uint      `json:"cId" gorm:"not null;comment:课程编号"`
	Number    uint      `json:"number" gorm:"not null;comment:购买数量"`
	Money     uint      `json:"money" gorm:"not null;comment:实付金额"`
	CreatedAt time.Time `json:"time"`
	UpdatedAt time.Time `json:"-"`
}

// EndRecord 核销记录
//【哪位会员】在【什么时间】核销【什么课程】【多少节】实际【上课教练】
type EndRecord struct {
	ID        uint      `json:"code" gorm:"primaryKey;autoIncrement;comment:唯一主键"`
	MemberId  uint      `json:"mId" gorm:"not null;comment:会员编号"`
	CourseId  uint      `json:"cId" gorm:"not null;comment:课程编号"`
	Number    uint      `json:"number" gorm:"not null;comment:核销数量"`
	CoachId   uint      `json:"tId" gorm:"not null;comment:任课教练"`
	CreatedAt time.Time `json:"time"`
	UpdatedAt time.Time `json:"-"`
}

// MemberRO 注册会员的请求参数
type MemberRO struct {
	Name   string `form:"name" binding:"required"`
	Card   string `form:"card" binding:"required"`
	Phone  string `form:"phone" binding:"required"`
	WeiXin string `form:"wx" binding:"required"`
	Image  string `form:"image" binding:"required"`
}

// BuyRecordRO 新增记录的请求参数
type BuyRecordRO struct {
	MemberId uint `form:"mId" binding:"required"`
	CourseId uint `form:"cId" binding:"required"`
	Number   uint `form:"number" binding:"required"`
	Money    uint `form:"money" binding:"required"`
}

// EndRecordRO 核销记录的请求参数
type EndRecordRO struct {
	MemberId uint `form:"mId" binding:"required"`
	CourseId uint `form:"cId" binding:"required"`
	Number   uint `form:"number" binding:"required"`
	CoachId  uint `form:"tId" binding:"required"`
}
