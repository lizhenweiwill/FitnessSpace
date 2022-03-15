package dao

import (
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
	err := AddOneMember(
		"振威", "No.222222", "123000000", "wx_123", "http://www.xx.png",
	)
	if err != nil {
		return
	}
	t.Log("操作成功")
}
