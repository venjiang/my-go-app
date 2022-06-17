package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var addr string = ":8080"

func init() {
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}
}

func main() {
	e := gin.Default()
	e.GET("/home", func(c *gin.Context) {
		c.String(200, "Gin Web App on Cloudflare @%v", time.Now().UnixMilli())
	})
	fmt.Println("Gin Web App running @", addr)
	e.Run(addr)
}
