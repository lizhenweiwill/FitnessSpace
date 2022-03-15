package service

import (
	"FitnessSpace/dao"
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
	name, ok := c.GetPostForm("name")
	if !ok {
		return
	}
	err := dao.AddOneMember(name, name, name, name, name)
	if err != nil {
		return
	}
	c.String(http.StatusOK, "添加成功")
}
