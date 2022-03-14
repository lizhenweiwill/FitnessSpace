package model

// Coach 教练
type Coach struct {
	Id   int
	Name string
	Nick string
	Pass string
}

// Course 课程
type Course struct {
	Id   int    `json:"code"`
	Name string `json:"name"`
}
