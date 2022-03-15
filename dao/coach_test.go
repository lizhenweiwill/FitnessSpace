package dao

import (
	"FitnessSpace/model"
	"testing"
)

func TestGetOneCoach(t *testing.T) {
	course := GetOneCoach("1")
	t.Log(course)
}

func TestGetAllCoach(t *testing.T) {
	course := GetAllCoach()
	t.Log(course)
}

func TestAddOneCoach(t *testing.T) {
	course := AddOneCoach("振威", "will", "000000")
	t.Log(course)
}

func TestDeleteCoach(t *testing.T) {
	course := DeleteCoach("1")
	t.Log(course)
}

func TestAddBatchesCoach(t *testing.T) {
	cs := []model.Coach{
		{Name: "张发", Nick: "ZF", Pass: "111111"},
		{Name: "王吉", Nick: "WJ", Pass: "222222"},
	}
	AddBatchesCoach(cs)
}
