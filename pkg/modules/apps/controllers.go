package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAppInput struct {
	Name string `json:"name" binding:"required"`
}

func handleCreateApp(c *gin.Context) {
	var input CreateAppInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserID").(string)

	app, err := CreateApp(input.Name, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": app})
}

func handleListApps(c *gin.Context) {
	userID := c.MustGet("UserID").(string)

	apps, err := ListApps(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": apps})
}

func handleGetApp(c *gin.Context) {
	id := c.Param("id")

	app, err := GetAppByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": app})
}
