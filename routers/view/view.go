package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowIndexPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, nil, "index")
}

func ShowRoulettePage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, nil, "roulette")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	//loggedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
