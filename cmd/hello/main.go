package main

import (
	"github.com/iotassss/saitamarental/internal/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", hello.HelloHandler)
	r.Run()
}
