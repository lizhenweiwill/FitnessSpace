package dao

import (
	"testing"
)

func TestGetOneCourse(t *testing.T) {
	course := GetOneCourse("1")
	t.Log(course)
}

func TestGetAllCourse(t *testing.T) {
	course := GetAllCourse()
	t.Log(course)
}

func TestAddOneCourse(t *testing.T) {
	course := AddOneCourse("常规")
	t.Log(course)
}

func TestDeleteCourse(t *testing.T) {
	course := DeleteCourse("10")
	t.Log(course)
}

func TestAddBatchesCourse(t *testing.T) {
	cs := []string{"拉伸", "搏击", "体态纠正"}
	AddBatchesCourse(cs)
}
