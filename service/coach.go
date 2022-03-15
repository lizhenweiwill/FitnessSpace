package service

import (
	"FitnessSpace/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCoach 查询教练
func GetCoach(c *gin.Context) {
	id := c.Param("id")
	course, err := dao.GetOneCoach(id)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(course)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// ListCoach 查询列表
func ListCoach(c *gin.Context) {
	list, err := dao.GetAllCoach()
	if err != nil {
		return
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddCoach 新增教练
func AddCoach(c *gin.Context) {
	name, ok := c.GetPostForm("name")
	if !ok {
		return
	}
	err := dao.AddOneCoach(name, name, name)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// DeleteCoach 删除教练
func DeleteCoach(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		return
	}
	err := dao.DeleteCoach(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
