package dao

import (
	"FitnessSpace/model"
)

func GetOneCoach(id string) (model.Coach, error) {
	var c model.Coach
	tx := db.Find(&c, id)
	return c, tx.Error
}

func GetAllCoach() ([]model.Coach, error) {
	var cs []model.Coach
	tx := db.Find(&cs)
	return cs, tx.Error
}

func AddOneCoach(name string, nick string, pass string) error {
	tx := db.Create(&model.Coach{Name: name, Nick: nick, Pass: pass})
	return tx.Error
}

func DeleteCoach(id string) error {
	tx := db.Delete(&model.Coach{}, id)
	return tx.Error
}

func AddBatchesCoach(cs []model.Coach) error {
	tx := db.CreateInBatches(cs, len(cs))
	return tx.Error
}
