package dao

import (
	"FitnessSpace/model"
	"FitnessSpace/provider"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = provider.NewDB()
	println("init database ~")
	migrator := db.Migrator()
	if migrator.HasTable(&model.Coach{}) {
		err := migrator.CreateTable(&model.Coach{})
		if err != nil {
			return
		}
	}
	//if migrator.HasTable(&model.Course{}) {
	//	err := migrator.CreateTable(&model.Course{})
	//	if err != nil {
	//		return
	//	}
	//}
	//if migrator.HasTable(&model.Member{}) {
	//	err := migrator.CreateTable(&model.Member{})
	//	if err != nil {
	//		return
	//	}
	//}
}

func GetOneCourse(id string) model.Course {
	course := model.Course{}
	find := db.Find(&course, id)
	if find.Error != nil {
		println(find.Error.Error())
	}
	return course
}

func GetAllCourse() []model.Course {
	result := db.Find(&model.Course{})
	list := make([]model.Course, result.RowsAffected)
	rows, err := result.Rows()
	if err != nil {
		return nil
	}
	for rows.Next() {
		//rows.
	}
	return list
}
