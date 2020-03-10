package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// Ping controller
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type Score struct {
	Num int
}

func (s *Score) Do() {
	time.Sleep(1 * 5 * time.Second)
	fmt.Println("num:", s.Num)
}

// Concurrency controller
func Concurrency(c *gin.Context) {
	workerNum := 100 * 100 * 2
	p := models.CreateWorkerPool(workerNum)
	p.Run()

	data := 100 * 100 * 100
	go func() {
		for i := 1; i <= data; i++ {
			sc := &Score{Num: i}
			p.JobQueue <- sc //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
		}
	}()

}
