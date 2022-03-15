package service

import (
	"FitnessSpace/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCourse 查询课程
func GetCourse(c *gin.Context) {
	id := c.Param("id")
	course := dao.GetOneCourse(id)
	marshal, err := json.Marshal(course)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListCourse 查询列表
func ListCourse(c *gin.Context) {
	list := dao.GetAllCourse()
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddCourse 新增课程
func AddCourse(c *gin.Context) {
	name, ok := c.GetPostForm("name")
	if !ok {
		return
	}
	dao.AddOneCourse(name)
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		return
	}
	dao.DeleteCourse(id)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
