package controllers

import (
	"cubicasa/models"
	"cubicasa/variables"
	"github.com/gin-gonic/gin"
	"gopkg.in/dealancer/validate.v2"
)

type HubController struct {
}

// Creat Hub
// @Summary Create New Hub
// @Description Create New Hub
// @Tags hub
// @ID create-hub
// @Accept  json
// @Produce  json
// @Param jsonBody body models.HubRequest true "hub data"
// @Success 200 {object} models.Response "include hub data in this response, see models.Hub"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /hub [post]
func (c *HubController) Create(ctx *gin.Context) {
	// parse json
	params := models.HubRequest{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// validate params
	if err := validate.Validate(&params); err != nil {
		// values not valid, deal with errors here
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	hub := models.Hub{
		Name:        params.Name,
		GeoLocation: params.GeoLocation,
	}
	err := variables.PostgresDB.Create(&hub).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    hub,
		})
	}

}
