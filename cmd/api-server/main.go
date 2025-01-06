package main

import (
	"fmt"
	"internal/conf"
	"internal/mysql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	config, err := conf.LoadConf("./goconf.json")
	if err != nil {
		slog.Error(err.Error())
	}

	db, err := mysql.Connect(config.GetMySQLUser(), config.GetMySQLPass())
	if err != nil {
		slog.Error(err.Error())
	}

	defer db.Close()

	r := createNewRouter()

	r.Run(fmt.Sprint(":", config.GetServerPort()))
}

func createNewRouter() *gin.Engine {
	r := gin.Default()

	localFile := static.LocalFile("../../client/build", true)

	r.Use(static.Serve("/", localFile))

	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
