package socketbase

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	SubscribeEvent = "pusher:subscribe"
)

type PusherSubscribeEvent struct {
	Event string `json:"event"`
	Data  struct {
		Channel string `json:"channel"`
	} `json:"data"`
}

// clients[app_id + channel] = connections[]
var clients = make(map[string][]*websocket.Conn)
var channels = make(map[string][]string)

func handleAppRoute(c *gin.Context) {
	id := c.Param("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		println(err.Error())
		return
	}

	conn.WriteJSON(map[string]interface{}{
		"event": "pusher:connection_established",
		"data": map[string]string{
			"socket_id":        id,
			"activity_timeout": "120",
		},
	})

	for {
		mt, message, err := conn.ReadMessage()

		println(string(message))

		if err != nil {
			break
		}

		if mt == websocket.TextMessage {
			var payload map[string]interface{}
			json.Unmarshal(message, &payload)

			if payload["event"] == "pusher:ping" {
				conn.WriteJSON(map[string]interface{}{
					"event": "pusher:pong",
				})
			} else if payload["event"].(string) == SubscribeEvent {
				channel := payload["data"].(map[string]interface{})["channel"].(string)

				key := id + channel

				if clients[key] == nil {
					clients[key] = make([]*websocket.Conn, 0)
					channels[key] = make([]string, 0)
				}

				clients[key] = append(clients[key], conn)
				channels[key] = append(channels[key], id)

				conn.WriteJSON(map[string]interface{}{
					"event":   "pusher_internal:subscription_succeeded",
					"data":    map[string]string{},
					"channel": channel,
				})
			} else if payload["channel"] != nil && payload["data"] != nil {
				event := payload["event"].(string)
				channel := payload["channel"].(string)
				data := payload["data"]

				key := id + channel

				for _, client := range clients[key] {
					client.WriteJSON(map[string]interface{}{
						"event":   event,
						"data":    data,
						"channel": channel,
					})
				}
			}
		}

		if mt == websocket.CloseMessage {
			break
		}
	}

	channels := channels[id]

	for _, channel := range channels {
		delete(clients, channel)
	}

	println("client disconnected")
}
