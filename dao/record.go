package dao

import (
	"FitnessSpace/model"
	"errors"
	"gorm.io/gorm"
)

func GetAllRecord() ([]model.Record, error) {
	var rs []model.Record
	tx := db.Find(&rs)
	return rs, tx.Error
}

// AddBuyRecord 事务控制
func AddBuyRecord(mId uint, cId uint, number uint, money uint) error {
	// 开启事务控制
	return db.Transaction(func(tx *gorm.DB) error {
		// 新增购买记录
		mb := model.BuyRecord{MemberId: mId, CourseId: cId, Number: number, Money: money}
		tx.Table("records_buy").Create(&mb)
		// 查询统计记录
		var mr = model.Record{}
		find := tx.Where("member_id = ? AND course_id = ?", mId, cId).Find(&mr)
		if find.RowsAffected == 0 {
			// 没有该记录，新建记录
			tx.Create(&model.Record{MemberId: mId, CourseId: cId, Count: number})
		} else {
			// 修改统计记录
			tx.Model(&model.Record{}).Where("member_id = ? AND course_id = ?", mId, cId).Update("count", mr.Count+number)
		}
		// 自动提交事务
		return nil
	})
}

// AddEndRecord 事务控制
func AddEndRecord(mId uint, cId uint, number uint, tId uint) error {
	// 开启事务控制
	return db.Transaction(func(tx *gorm.DB) error {
		// 新增销课记录
		me := &model.EndRecord{MemberId: mId, CourseId: cId, Number: number, CoachId: tId}
		tx.Table("records_end").Create(&me)
		// 查询统计记录
		var mr = model.Record{}
		find := tx.Where("member_id = ? AND course_id = ?", mId, cId).Find(&mr)
		if find.RowsAffected != 0 && mr.Count > number {
			// 修改统计记录
			tx.Model(&model.Record{}).Where("member_id = ? AND course_id = ?", mId, cId).Update("count", mr.Count-number)
		} else {
			// 课程无法核销 - 回退事务
			return errors.New("课程无法核销")
		}
		// 自动提交事务
		return nil
	})
}
