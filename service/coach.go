package service

import (
	"github.com/gin-gonic/gin"
)

// GetCoach 查询教练
func GetCoach(c *gin.Context) {
	//id := c.Param("id")
	//course := mCoach[id]
	//marshal, err := json.Marshal(course)
	//if err != nil {
	//	return
	//}
	//c.String(http.StatusOK, string(marshal))
}

// ListCoach 查询列表
func ListCoach(c *gin.Context) {
	//var list []model.Coach
	//for _, v := range mCoach {
	//	list = append(list, v)
	//}
	//marshal, err := json.Marshal(list)
	//if err != nil {
	//	return
	//}
	//c.String(http.StatusOK, string(marshal))
}

// AddCoach 新增教练
// TODO : 数据库主键回填
func AddCoach(c *gin.Context) {
	//name, _ := c.GetPostForm("name")
	//mc := model.Coach{
	//	Id:   3, // 数据库回填主键
	//	Name: name,
	//}
	//mCoach["3"] = mc
	//c.String(http.StatusOK, "添加成功")
}

// DeleteCoach 删除教练
func DeleteCoach(c *gin.Context) {
	//if id, ok := c.GetPostForm("id"); ok {
	//	if _, ok := mCoach[id]; ok {
	//		delete(mCoach, id)
	//	}
	//}
	//c.String(http.StatusOK, "删除成功")
}
