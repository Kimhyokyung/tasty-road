package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tasty-road/routers/api/v1"
	"github.com/tasty-road/routers/view"
	_ "github.com/tasty-road/routers/view"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/layouts/*.html")

	r.GET("/", view.ShowIndexPage)

	// Set View Routing Rules
	viewPage := r.Group("/view")
	{
		viewPage.GET("/roulette", view.ShowRoulettePage)
	}

	// Set API v1.0 Routing Rules
	apiv1 := r.Group("/api/v1")
	{
		// Router for Restaurants model
		apiv1.GET("/restaurants", v1.GetRestaurants)
		apiv1.GET("/restaurant/:id", v1.GetRestaurant)
	}

	return r
}
