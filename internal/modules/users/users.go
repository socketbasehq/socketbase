package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/socketbase/socketbase/internal/pkg/server"

	"go.uber.org/fx"
)

func NewUsersModule() server.Module {
	return server.Module{
		Routes: server.RouteGroup{
			Prefix: "/users",
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/",
					Handler: listUsers,
				},
				{
					Method:  "POST",
					Path:    "/",
					Handler: createUser,
				},
			},
		},
	}
}

// Handler implementations
func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func createUser(c *gin.Context) {
	// ... existing implementation
}

var Users = fx.Annotate(
	NewUsersModule,
	fx.ResultTags(`group:"routes"`),
)
