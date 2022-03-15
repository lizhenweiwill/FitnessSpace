package dao

import (
	"FitnessSpace/model"
)

func GetOneCoach(id string) model.Coach {
	c := model.Coach{}
	db.Find(&c, id)
	return c
}

func GetAllCoach() (cs []model.Coach) {
	result := db.Find(&cs)
	if result.Error != nil {
		println(result.Error.Error())
	}
	return cs
}

func AddOneCoach(name string, nick string, pass string) model.Coach {
	c := model.Coach{Name: name, Nick: nick, Pass: pass}
	db.Create(&c)
	return c
}

func DeleteCoach(id string) model.Coach {
	c := model.Coach{}
	db.Delete(&c, id)
	return c
}

func AddBatchesCoach(cs []model.Coach) {
	db.CreateInBatches(cs, len(cs))
}
