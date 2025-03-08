package main

import (
	"github.com/iotassss/saitamarent/internal/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", hello.HelloHandler)
	r.Run()
}
