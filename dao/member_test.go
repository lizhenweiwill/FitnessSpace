package dao

import (
	"FitnessSpace/model"
	"testing"
)

func TestGetOneMember(t *testing.T) {
	member, err := GetOneMember("1")
	if err != nil {
		return
	}
	t.Log(member)
}

func TestGetAllMember(t *testing.T) {
	member, err := GetAllMember()
	if err != nil {
		return
	}
	t.Log(member)
}

func TestAddOneMember(t *testing.T) {
	err := AddOneMember(&model.MemberRO{
		Name: "振威", Card: "No.222222", Phone: "123000000", WeiXin: "wx_123", Image: "http://www.xx.png",
	})
	if err != nil {
		return
	}
	t.Log("操作成功")
}
