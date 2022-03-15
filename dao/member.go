package dao

import (
	"FitnessSpace/model"
)

func GetOneMember(id string) (model.Member, error) {
	var m model.Member
	tx := db.Find(&m, id)
	return m, tx.Error
}

func GetAllMember() ([]model.Member, error) {
	var ms []model.Member
	tx := db.Find(&ms)
	return ms, tx.Error
}

func AddOneMember(name string, card string, phone string, wx string, image string) error {
	m := &model.Member{Name: name, Card: card, Phone: phone, WeiXin: wx, Image: image}
	tx := db.Create(&m)
	return tx.Error
}
