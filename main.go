package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var server string
var port int
var help bool

func main() {
	go GracefulShutdown()

	flag.Bool("help", false, "Show help")
	flag.StringVar(&server, "server", "127.0.0.1", "server IPV4 address")
	flag.IntVar(&port, "port", 7029, "server port")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	go Startup()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{server})

	router.GET("/healthcheck", healthCheck)
	router.POST("/create", createRecord)
	router.GET("/get", getRecord)
	router.DELETE("/delete", deleteRecord)
	router.GET("/flush", flushKeys)
	router.Run(":" + fmt.Sprint(port))
}
