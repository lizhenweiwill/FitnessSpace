package service

import (
	"FitnessSpace/dao"
	"FitnessSpace/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMember 查询会员
func GetMember(c *gin.Context) {
	id := c.Param("id")
	member, err := dao.GetOneMember(id)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(member)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListMember 查询列表
func ListMember(c *gin.Context) {
	list, err := dao.GetAllMember()
	if err != nil {
		return
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddMember 新增会员
func AddMember(c *gin.Context) {
	var mr = model.MemberRO{}
	var err error
	err = c.ShouldBind(&mr)
	if err != nil {
		return
	}
	err = dao.AddOneMember(&mr)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}
