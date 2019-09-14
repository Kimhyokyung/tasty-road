package view

import (
	"github.com/gin-gonic/gin"
	"github.com/tasty-road/render"
)

func ShowIndexPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render.Render(c, nil, "index")
}

func ShowRoulettePage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render.Render(c, nil, "roulette")
}
