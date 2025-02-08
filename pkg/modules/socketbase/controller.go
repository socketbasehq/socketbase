package socketbase

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"

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
var clientsMutex = sync.RWMutex{}

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

func handleAppEvent(c *gin.Context) {
	appId := c.Param("id")
	jsonData, err := io.ReadAll(c.Request.Body)
	query := c.Request.URL.Query()

	if err != nil {
		println(err.Error())
		return
	}

	authSignature := query["auth_signature"][0]
	delete(query, "auth_signature")

	app, err := GetApp(appId)

	if err != nil {
		println(err.Error())
		return
	}

	md5 := md5.Sum(jsonData)
	md5String := hex.EncodeToString(md5[:])
	query["body_md5"] = []string{md5String}

	method := c.Request.Method
	path := c.Request.URL.Path
	sortedQuery := toOrderedArray(query)

	payload := method + "\n" + path + "\n" + strings.Join(sortedQuery, "&")

	token := sign(app.AppSecret, payload)

	if token != authSignature {
		println("invalid signature")
		return
	}

	bodyMap := make(map[string]interface{})
	json.Unmarshal(jsonData, &bodyMap)

	// loop channels
	for _, channel := range bodyMap["channels"].([]interface{}) {
		clientsMutex.RLock()
		clientsList := clients[app.AppKey+channel.(string)]
		clientsMutex.RUnlock()

		for _, client := range clientsList {
			clientsMutex.Lock()
			err := client.WriteJSON(map[string]interface{}{
				"event":   bodyMap["name"],
				"data":    bodyMap["data"],
				"channel": channel,
			})
			clientsMutex.Unlock()
			if err != nil {
				println("Error writing to client:", err.Error())
			}
		}
	}

	println(bodyMap)
}

func sign(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func toOrderedArray(m url.Values) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	result := make([]string, len(keys))
	for i, key := range keys {
		result[i] = key + "=" + m[key][0]
	}

	return result
}
