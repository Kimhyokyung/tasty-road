package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tasty-road/models"
	"github.com/tasty-road/render"
)

func GetLunchs(c *gin.Context) {
	lunchList := models.GetLunchList()
	c.Request.Header.Set("Accept", "application/json")
	render.Render(c, gin.H{"payload": lunchList}, "")
}
