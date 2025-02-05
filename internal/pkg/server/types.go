package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Route represents a single HTTP route
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// RouteGroup represents a group of routes with a common prefix
type RouteGroup struct {
	Prefix string
	Routes []Route
}

// Module is the common type for all modules
type Module struct {
	fx.Out
	Routes RouteGroup `group:"routes"`
}
