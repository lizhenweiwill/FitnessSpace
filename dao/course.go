package dao

import (
	"FitnessSpace/model"
)

func GetOneCourse(id string) model.Course {
	c := model.Course{}
	db.Find(&c, id)
	return c
}

func GetAllCourse() (cs []model.Course) {
	result := db.Find(&cs)
	if result.Error != nil {
		println(result.Error.Error())
	}
	return cs
}

func AddOneCourse(name string) model.Course {
	c := model.Course{Name: name}
	db.Create(&c)
	return c
}

func DeleteCourse(id string) model.Course {
	c := model.Course{}
	db.Delete(&c, id)
	return c
}

func AddBatchesCourse(cs []string) {
	db.CreateInBatches(cs, len(cs))
}
