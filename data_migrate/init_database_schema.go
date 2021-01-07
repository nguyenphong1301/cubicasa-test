package main

import (
	"cubicasa/models"
	"cubicasa/variables"
	"log"
)

func main() {
	err := variables.Init()
	if err != nil {
		log.Fatal("init application failed with error " + err.Error())
	}
	defer variables.DeInit()

	hubModel := &models.Hub{}
	teamModel := &models.Team{}
	userModel := &models.User{}

	dbConn := variables.PostgresDB

	dbConn.DropTableIfExists(userModel)
	dbConn.DropTableIfExists(teamModel)
	dbConn.DropTableIfExists(hubModel)

	dbConn.CreateTable(hubModel)
	dbConn.CreateTable(teamModel)
	dbConn.CreateTable(userModel)
}
