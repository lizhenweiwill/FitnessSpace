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
// TODO : 数据库主键回填
func AddCourse(c *gin.Context) {
	//if name, ok := c.GetPostForm("name"); ok {
	//mc := model.Course{
	//	Id:   5, // 数据库回填主键
	//	Name: name,
	//}
	//mCourse["5"] = mc
	//}
	//c.String(http.StatusOK, "添加成功")
}

// DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	//if id, ok := c.GetPostForm("id"); ok {
	//	if _, ok := mCourse[id]; ok {
	//		delete(mCourse, id)
	//	}
	//}
	//c.String(http.StatusOK, "删除成功")
}
