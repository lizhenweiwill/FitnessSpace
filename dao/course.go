package dao

import (
	"FitnessSpace/model"
)

func GetOneCourse(id string) (model.Course, error) {
	c := model.Course{}
	tx := db.Find(&c, id)
	return c, tx.Error
}

func GetAllCourse() ([]model.Course, error) {
	var cs []model.Course
	tx := db.Find(&cs)
	return cs, tx.Error
}

func AddOneCourse(name string) error {
	tx := db.Create(&model.Course{Name: name})
	return tx.Error
}

func DeleteCourse(id string) error {
	tx := db.Delete(&model.Course{}, id)
	return tx.Error
}

func AddBatchesCourse(cs []string) error {
	tx := db.CreateInBatches(cs, len(cs))
	return tx.Error
}
