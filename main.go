package main

import (
	"fmt"
	"internal/conf"
	"internal/mysql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	config, err := conf.LoadConf("goconf.json")
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

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("/login", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
