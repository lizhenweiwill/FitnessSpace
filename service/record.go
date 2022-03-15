package service

import (
	"FitnessSpace/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListRecord 查询列表
func ListRecord(c *gin.Context) {
	list, err := dao.GetAllRecord()
	if err != nil {
		return
	}
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	c.String(http.StatusOK, string(marshal))
}

// AddBuyRecord 新增购买记录
func AddBuyRecord(c *gin.Context) {
	_, ok := c.GetPostForm("name")
	if !ok {
		return
	}
	err := dao.AddBuyRecord(1, 9, 10, 3000)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// AddEndRecord 新增销课记录
func AddEndRecord(c *gin.Context) {
	_, ok := c.GetPostForm("id")
	if !ok {
		return
	}
	err := dao.AddEndRecord(1, 9, 1, 3)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
