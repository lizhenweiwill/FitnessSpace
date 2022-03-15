package dao

import "testing"

func TestGetAllRecord(t *testing.T) {
	record, err := GetAllRecord()
	if err != nil {
		return
	}
	t.Log(record)
}

// 新用户购买新课程
func TestAddBuyRecord1(t *testing.T) {
	err := AddBuyRecord(1, 9, 10, 3000)
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户加购该课程
func TestAddBuyRecord2(t *testing.T) {
	err := AddBuyRecord(1, 9, 20, 6000)
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户加购其他课程
func TestAddBuyRecord3(t *testing.T) {
	err := AddBuyRecord(1, 8, 20, 6000)
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 该用户核销一节课程
func TestAddEndRecord1(t *testing.T) {
	err := AddEndRecord(1, 8, 1, 3)
	if err != nil {
		return
	}
	t.Log("操作成功")
}

// 非法核销 - 课程不存在
func TestAddEndRecord2(t *testing.T) {
	err := AddEndRecord(1, 2, 1, 3)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("操作成功")
}

// 非法核销 - 剩余课程数不够
func TestAddEndRecord3(t *testing.T) {
	err := AddEndRecord(1, 8, 100, 3)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log("操作成功")
}
