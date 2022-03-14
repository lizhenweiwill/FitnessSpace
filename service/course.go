package service

import (
	"FitnessSpace/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// mock
var m = map[string]model.Course{
	"1": {1, "常规"},
	"2": {2, "拉伸"},
	"3": {3, "搏击"},
	"4": {4, "体态纠正"},
}

// GetCourse 查询课程
func GetCourse(c *gin.Context) {
	id := c.Param("id")
	course := m[id]
	marshal, err := json.Marshal(course)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListCourse 查询列表
func ListCourse(c *gin.Context) {
	var list []model.Course
	for _, v := range m {
		list = append(list, v)
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddCourse 新增课程
// TODO : 数据库主键回填
func AddCourse(c *gin.Context) {
	if name, ok := c.GetPostForm("name"); ok {
		mc := model.Course{
			Id:   5, // 数据库回填主键
			Name: name,
		}
		m["5"] = mc
	}
	c.String(http.StatusOK, "添加成功")
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	if id, ok := c.GetPostForm("id"); ok {
		if _, ok := m[id]; ok {
			delete(m, id)
		}
	}
	c.String(http.StatusOK, "删除成功")
}