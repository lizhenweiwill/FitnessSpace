package main

import (
	"FitnessSpace/service"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	group := engine.Group("/api/v1")
	{
		// 教练相关
		group.GET("/coach/:id", service.GetCoach)
		group.GET("/coach/list", service.ListCoach)
		group.POST("/coach", service.AddCoach)
		group.DELETE("/coach", service.DeleteCoach)
		// 课程相关
		group.GET("/course/:id", service.GetCourse)
		group.GET("/course/list", service.ListCourse)
		group.POST("/course", service.AddCourse)
		group.DELETE("/course", service.DeleteCourse)
		// 会员相关
		group.GET("/member/:id", service.GetMember)
		group.GET("/member/list", service.ListMember)
		group.POST("/member", service.AddMember)
		//group.PUT("/member") // 修改信息
		// 记录相关
		group.POST("/record")
		group.PUT("/record")
	}
	err := engine.Run(":8080")
	if err != nil {
		panic("服务启动失败！")
	}
}
