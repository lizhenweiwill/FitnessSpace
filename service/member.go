package service

import (
	"FitnessSpace/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// mock
var mMember = map[string]model.Member{
	"1": {1, "张三", "", "", "", ""},
	"2": {2, "李四", "", "", "", ""},
}

// GetMember 查询会员
func GetMember(c *gin.Context) {
	id := c.Param("id")
	course := mMember[id]
	marshal, err := json.Marshal(course)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListMember 查询列表
func ListMember(c *gin.Context) {
	var list []model.Member
	for _, v := range mMember {
		list = append(list, v)
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddMember 新增课程
// TODO : 数据库主键回填
func AddMember(c *gin.Context) {
	if name, ok := c.GetPostForm("name"); ok {
		mc := model.Member{
			Id:   3, // 数据库回填主键
			Name: name,
		}
		mMember["3"] = mc
	}
	c.String(http.StatusOK, "添加成功")
}
