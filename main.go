package main

import (
	"cubicasa/configs"
	"cubicasa/routes"
	"cubicasa/variables"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := variables.Init()
	if err != nil {
		log.Fatal("init application failed with error " + err.Error())
	}
	defer variables.DeInit()

	// load routes
	r := gin.New()
	r.Use(gin.Recovery())
	// allow origin
	r.Use(cors.Default())
	routes.Init(r)

	// start web application
	r.Run(fmt.Sprintf(":%v", configs.Port))
}
