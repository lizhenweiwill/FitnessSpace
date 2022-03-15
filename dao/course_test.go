package dao

import (
	"testing"
)

func TestGetOneCourse(t *testing.T) {
	course, err := GetOneCourse("1")
	if err != nil {
		return
	}
	t.Log(course)
}

func TestGetAllCourse(t *testing.T) {
	course, err := GetAllCourse()
	if err != nil {
		return
	}
	t.Log(course)
}

func TestAddOneCourse(t *testing.T) {
	err := AddOneCourse("常规")
	if err != nil {
		return
	}
	t.Log("操作成功")
}

func TestDeleteCourse(t *testing.T) {
	err := DeleteCourse("10")
	if err != nil {
		return
	}
	t.Log("操作成功")
}

func TestAddBatchesCourse(t *testing.T) {
	cs := []string{"拉伸", "搏击", "体态纠正"}
	err := AddBatchesCourse(cs)
	if err != nil {
		return
	}
	t.Log("操作成功")
}
