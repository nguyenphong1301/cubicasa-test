package controllers

import (
	"cubicasa/models"
	"cubicasa/utils"
	"cubicasa/variables"
	"github.com/gin-gonic/gin"
	"gopkg.in/dealancer/validate.v2"
)

type UserController struct {
}

// Creat User
// @Summary Create New User
// @Description Create New User
// @Tags user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param jsonBody body models.UserRequest true "Role is one of (Intern, S1, S2, S3, M1, M2, Head), team_id is optional, may be assign later"
// @Success 200 {object} models.Response "include user data, see models.User"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /user [post]
func (c *UserController) Create(ctx *gin.Context) {

	// parse json
	params := models.UserRequest{}
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
	if !c.isValidRole(params.Role) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Role is invalid",
		})
		return
	}

	// validate team
	if params.TeamID != 0 && !c.validateTeamID(params.TeamID) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Team is not exist",
		})
		return
	}

	// save database
	user := models.User{
		Email:  params.Email,
		Role:   params.Role,
		TeamID: params.TeamID,
	}
	err := variables.PostgresDB.Create(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    user,
		})
	}
}

// User Team Assign
// @Summary User Team Assign
// @Description User Team Assign
// @Tags user
// @ID assign-user-team
// @Accept  json
// @Produce  json
// @Param jsonBody body models.TeamAssign true "assigned data"
// @Success 200 {object} models.Response "new user data"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /user/assign-to-team [post]
func (c *UserController) AssignTeam(ctx *gin.Context) {

	// parse json
	params := models.TeamAssign{}
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	// find Team
	var user models.User
	err := variables.PostgresDB.Model(&user).Where("id = ?", params.TeamID).First(&user).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "user is not exist",
		})
		return
	}

	// validate hub
	if !c.validateTeamID(params.TeamID) {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "Team is not exist",
		})
		return
	}

	// update database
	user.AssignTeam(params.TeamID)
	err = variables.PostgresDB.Save(&user).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data":    user,
		})
	}

}

func (c *UserController) isValidRole(role string) bool {
	// check role is in defined array
	if !utils.InArray(role, []string{"Intern", "S1", "S2", "S3", "M1", "M2", "Head"}) {
		return false
	}
	return true
}

func (c *UserController) validateTeamID(teamID int64) bool {
	// check if team exist in database
	var team models.Team
	err := variables.PostgresDB.Where("id = ?", teamID).First(&team).Error
	if err != nil {
		return false
	}
	return true
}
