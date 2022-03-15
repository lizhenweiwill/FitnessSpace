package dao

import (
	"FitnessSpace/model"
	"testing"
)

func TestGetOneCoach(t *testing.T) {
	coach, err := GetOneCoach("1")
	if err != nil {
		return
	}
	t.Log(coach)
}

func TestGetAllCoach(t *testing.T) {
	coach, err := GetAllCoach()
	if err != nil {
		return
	}
	t.Log(coach)
}

func TestAddOneCoach(t *testing.T) {
	err := AddOneCoach("振威", "will", "000000")
	if err != nil {
		return
	}
	t.Log("操作成功")
}

func TestDeleteCoach(t *testing.T) {
	err := DeleteCoach("10")
	if err != nil {
		return
	}
	t.Log("操作成功")
}

func TestAddBatchesCoach(t *testing.T) {
	cs := []model.Coach{
		{Name: "张发", Nick: "ZF", Pass: "111111"},
		{Name: "王吉", Nick: "WJ", Pass: "222222"},
	}
	err := AddBatchesCoach(cs)
	if err != nil {
		return
	}
	t.Log("操作成功")
}
