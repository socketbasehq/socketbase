package server

import (
	"bytes"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Route struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type RouteGroup struct {
	Prefix      string
	Routes      []Route
	Middlewares []gin.HandlerFunc
}

type ServerParams struct {
	fx.In
	RouteGroups []RouteGroup `group:"routes"`
}

//go:embed dist
var DistDir embed.FS

func NewServer(params ServerParams) *gin.Engine {
	app := gin.Default()

	for _, group := range params.RouteGroups {
		router := app.Group(group.Prefix)

		for _, route := range group.Routes {
			middlewares := []gin.HandlerFunc{}

			if len(group.Middlewares) > 0 {
				middlewares = append(middlewares, group.Middlewares...)
			}
			if len(route.Middlewares) > 0 {
				middlewares = append(middlewares, route.Middlewares...)
			}

			if len(middlewares) > 0 {
				router.Handle(route.Method, route.Path, append(middlewares, route.Handler)...)
			} else {
				router.Handle(route.Method, route.Path, route.Handler)
			}
		}
	}

	staticFiles, err := fs.Sub(DistDir, "dist")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServer(http.FS(staticFiles))

	app.NoRoute(func(c *gin.Context) {
		if exists(c.Request.URL.Path) {
			fileServer.ServeHTTP(c.Writer, c.Request)
		} else {
			index, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer index.Close()

			content, err := fs.ReadFile(staticFiles, "index.html")
			if err != nil {
				log.Fatal(err)
			}

			stat, err := index.Stat()
			if err != nil {
				log.Fatal(err)
			}

			http.ServeContent(c.Writer, c.Request, "index.html", stat.ModTime(), bytes.NewReader(content))
		}
	})

	return app
}

func exists(path string) bool {
	_, err := DistDir.Open("dist" + path)
	return err == nil
}
