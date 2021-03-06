package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/MarsMan1994/gout"
	"time"
)

const (
	benchNumber     = 30000
	benchConcurrent = 20
)

func server() {
	router := gin.New()
	router.POST("/", func(c *gin.Context) {
		c.String(200, "12345")
	})

	router.Run()
}

func main() {
	go server()
	time.Sleep(time.Millisecond)

	err := gout.
		POST(":8080").
		SetJSON(gout.H{"hello": "world"}). //设置请求body内容
		Filter().                          //打开过滤器
		Bench().                           //选择bench功能
		Concurrent(benchConcurrent).       //并发数
		Number(benchNumber).               //压测次数
		Do()

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
