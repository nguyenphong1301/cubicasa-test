package controllers

import (
	"cubicasa/models"
	"cubicasa/utils"
	"cubicasa/variables"
	"github.com/gin-gonic/gin"
	"gopkg.in/dealancer/validate.v2"
)

type TeamController struct {
}

// Creat Team
// @Summary Create New Team
// @Description Create New Team
// @Tags team
// @ID create-team
// @Accept  json
// @Produce  json
// @Param jsonBody body models.TeamRequest true "Team Type is one of (Platform, Fin, Music, Publishing, Studio), HubId is optional, maybe assign later"
// @Success 200 {object} models.Response "include team data, see models.Team"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /team [post]
func (c *TeamController) Create(ctx *gin.Context) {

	// parse json
	params := models.TeamRequest{}
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

	// validate team type
	if !c.isValidTeamType(params.TeamType) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Team type is invalid",
		})
		return
	}

	// validate hub
	if params.HubID != 0 && !c.validateHub(params.HubID) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Hub is not exist",
		})
		return
	}

	team := models.Team{
		Name:     params.Name,
		TeamType: params.TeamType,
		HubID:    params.HubID,
	}

	// save database
	err := variables.PostgresDB.Create(&team).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    team,
		})
	}

}

// Team Hub Assign
// @Summary Team Hub Assign
// @Description Team Hub Assign
// @Tags team
// @ID assign-team-hub
// @Accept  json
// @Produce  json
// @Param jsonBody body models.HubAssign true "assigned data"
// @Success 200 {object} models.Response "new team data"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /team/assign-to-hub [post]
func (c *TeamController) AssignHub(ctx *gin.Context) {

	// parse json
	params := models.HubAssign{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// find Team
	var team models.Team
	err := variables.PostgresDB.Model(&team).Where("id = ?", params.TeamID).First(&team).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Team is not exist",
		})
		return
	}

	// validate hub
	if !c.validateHub(params.HubID) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Hub is not exist",
		})
		return
	}

	// update database
	team.AssignHub(params.HubID)
	err = variables.PostgresDB.Save(&team).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    team,
		})
	}

}

func (c *TeamController) isValidTeamType(teamType string) bool {

	// check team type in in defined Array
	if !utils.InArray(teamType, []string{"Platform", "Fin", "Music", "Publishing", "Studio"}) {
		return false
	}
	return true
}

func (c *TeamController) validateHub(hubID int64) bool {

	// check if hub exist in database
	var hub models.Hub
	err := variables.PostgresDB.Where("id = ?", hubID).First(&hub).Error
	if err != nil {
		return false
	}
	return true
}
