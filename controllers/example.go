package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

type respBody struct {
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

// Ping controller
func Ping(c *gin.Context) {
	resMap := map[string]interface{}{
		"message": "pong",
	}
	c.JSON(http.StatusOK, resMap)
	// res := &respBody{
	// 	Message: "hello",
	// 	Data:    []string{"apple", "banana"},
	// }
	// jsonRes, _ := json.Marshal(res)
	// c.Writer.WriteHeader(http.StatusOK)
	// c.Writer.Write(jsonRes)
}

// Pong controller
func Pong(c *gin.Context) {
	id := c.Param("id")
	name := c.DefaultQuery("name", "osir")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
	// res := &respBody{
	// 	Message: "hello",
	// 	Data:    []string{"apple", "banana"},
	// }
	// jsonRes, _ := json.Marshal(res)
	// c.Writer.WriteHeader(http.StatusOK)
	// c.Writer.Write(jsonRes)
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
