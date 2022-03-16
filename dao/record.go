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
func AddBuyRecord(br *model.BuyRecordRO) error {
	// 开启事务控制
	return db.Transaction(func(tx *gorm.DB) error {
		// 新增购买记录
		mb := model.BuyRecord{MemberId: br.MemberId, CourseId: br.CourseId, Number: br.Number, Money: br.Money}
		tx.Table("records_buy").Create(&mb)
		// 查询统计记录
		var mr = model.Record{}
		find := tx.Where("member_id = ? AND course_id = ?", br.MemberId, br.CourseId).Find(&mr)
		if find.RowsAffected == 0 {
			// 没有该记录，新建记录
			tx.Create(&model.Record{MemberId: br.MemberId, CourseId: br.CourseId, Count: br.Number})
		} else {
			// 修改统计记录
			tx.Model(&model.Record{}).Where("member_id = ? AND course_id = ?", br.MemberId, br.CourseId).Update("count", mr.Count+br.Number)
		}
		// 自动提交事务
		return nil
	})
}

// AddEndRecord 事务控制
func AddEndRecord(er *model.EndRecordRO) error {
	// 开启事务控制
	return db.Transaction(func(tx *gorm.DB) error {
		// 新增销课记录
		me := &model.EndRecord{MemberId: er.MemberId, CourseId: er.CourseId, Number: er.Number, CoachId: er.CoachId}
		tx.Table("records_end").Create(&me)
		// 查询统计记录
		var mr = model.Record{}
		find := tx.Where("member_id = ? AND course_id = ?", er.MemberId, er.CourseId).Find(&mr)
		if find.RowsAffected != 0 && mr.Count > er.Number {
			// 修改统计记录
			tx.Model(&model.Record{}).Where("member_id = ? AND course_id = ?", er.MemberId, er.CourseId).Update("count", mr.Count-er.Number)
		} else {
			// 课程无法核销 - 回退事务
			return errors.New("课程无法核销")
		}
		// 自动提交事务
		return nil
	})
}
