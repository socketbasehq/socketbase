package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/socketbasehq/socketbase/pkg/server"
	"github.com/socketbasehq/socketbase/pkg/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWSModule() types.Module {
	return types.Module{
		Routes: server.RouteGroup{
			Prefix: "/ws",
			Routes: []server.Route{
				{
					Method:  "GET",
					Path:    "/",
					Handler: handleWebSocket,
				},
			},
		},
	}
}

func handleWebSocket(c *gin.Context) {
	upgrader.Upgrade(c.Writer, c.Request, nil)
}
