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

func AddOneMember(mr *model.MemberRO) error {
	m := &model.Member{Name: mr.Name, Card: mr.Card, Phone: mr.Phone, WeiXin: mr.WeiXin, Image: mr.Image}
	tx := db.Create(&m)
	return tx.Error
}
