package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.Group("/api/v1")
	{
		// 教练相关
		engine.GET("/coach/:id")
		engine.POST("/coach")
		engine.DELETE("/coach")
		// 课程相关
		engine.GET("/course/:id")
		engine.POST("/course")
		engine.DELETE("/course")
		// 会员相关
		engine.GET("/member/:id")
		engine.POST("/member")
		engine.PUT("/member") // 修改信息
		// 记录相关
		engine.POST("/record")
		engine.PUT("/record")
	}
	err := engine.Run(":8080")
	if err != nil {
		panic("服务启动失败！")
	}
}
