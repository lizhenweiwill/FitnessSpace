package service

import (
	"FitnessSpace/dao"
	"FitnessSpace/model"
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
	var br = model.BuyRecordRO{}
	var err error
	err = c.ShouldBind(&br)
	if err != nil {
		return
	}
	err = dao.AddBuyRecord(&br)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功"})
}

// AddEndRecord 新增销课记录
func AddEndRecord(c *gin.Context) {
	var er = model.EndRecordRO{}
	var err error
	err = c.ShouldBind(&er)
	if err != nil {
		return
	}
	err = dao.AddEndRecord(&er)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
