package main

import (
	"08-go-solidity-when/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routers.ApiRoutersInit(r)
	_ = r.Run()
}
