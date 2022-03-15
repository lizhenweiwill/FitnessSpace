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
	course, err := dao.GetOneCourse(id)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(course)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListCourse 查询列表
func ListCourse(c *gin.Context) {
	list, err := dao.GetAllCourse()
	if err != nil {
		return
	}
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
	err := dao.AddOneCourse(name)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		return
	}
	err := dao.DeleteCourse(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
