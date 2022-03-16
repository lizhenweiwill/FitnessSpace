package dao

import (
	"FitnessSpace/model"
	"testing"
)

func TestGetAllRecord(t *testing.T) {
	record, err := GetAllRecord()
	if err != nil {
		return
	}
	t.Log(record)
}

// 新用户购买新课程
func TestAddBuyRecord1(t *testing.T) {
	err := AddBuyRecord(&model.BuyRecordRO{MemberId: 1, CourseId: 9, Number: 10, Money: 3000})
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户加购该课程
func TestAddBuyRecord2(t *testing.T) {
	err := AddBuyRecord(&model.BuyRecordRO{MemberId: 1, CourseId: 9, Number: 20, Money: 6000})
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户加购其他课程
func TestAddBuyRecord3(t *testing.T) {
	err := AddBuyRecord(&model.BuyRecordRO{MemberId: 1, CourseId: 1, Number: 30, Money: 9000})
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户核销一节课程
func TestAddEndRecord1(t *testing.T) {
	err := AddEndRecord(&model.EndRecordRO{MemberId: 1, CourseId: 8, Number: 1, CoachId: 3})
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 非法核销 - 课程不存在
func TestAddEndRecord2(t *testing.T) {
	err := AddEndRecord(&model.EndRecordRO{MemberId: 1, CourseId: 2, Number: 1, CoachId: 3})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("操作成功")
}

// 非法核销 - 剩余课程数不够
func TestAddEndRecord3(t *testing.T) {
	err := AddEndRecord(&model.EndRecordRO{MemberId: 1, CourseId: 8, Number: 100, CoachId: 3})
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("操作成功")
}
