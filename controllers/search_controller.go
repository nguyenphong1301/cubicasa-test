package controllers

import (
	"cubicasa/models"
	"cubicasa/variables"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SearchController struct {
}

// Search
// @Summary Search for hub/team information
// @Description This api search hub and team information are matched with the key word (name, location, Type)
// @Tags search
// @ID search
// @Accept  json
// @Produce  json
// @Param key query string  false "search keyword"
// @Success 200 {object} models.Response "include array of hub and team are matched the search key"
// @Failure 400 {object} models.Response ""
// @Failure 500 {object} models.Response ""
// @Router /search [get]
func (c *SearchController) Search(ctx *gin.Context) {
	key := ctx.Query("key")
	var hubs []models.Hub
	var teams []models.Team
	hubBuilder := variables.PostgresDB.Model(&models.Hub{})
	teamBuilder := variables.PostgresDB.Model(&models.Team{})

	if len(key) > 0 {
		paramsValue := fmt.Sprintf(`%%%v%%`, key)
		hubBuilder = hubBuilder.Where("name like ? OR geo_location like ? ", paramsValue, paramsValue)
		teamBuilder = teamBuilder.
			Joins("left join hubs on teams.hub_id = hubs.id").
			Where("teams.name like ? OR teams.team_type like ? OR hubs.name like ? OR hubs.geo_location like ?", paramsValue, paramsValue, paramsValue, paramsValue)
	}

	err1 := hubBuilder.Find(&hubs).Error
	err2 := teamBuilder.Find(&teams).Error
	if err1 != nil || err1 != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": fmt.Sprintf(`%v,%v`, err1, err2),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"hubs":  hubs,
				"teams": teams,
			},
		})
	}
}
