package app

import (
	"blog/controllers/pingcontroller"
	"blog/controllers/userscontroller"
)

func mapUrls() {
	router.GET("/ping", pingcontroller.Ping)
	router.GET("/users/:user_id", userscontroller.GetUser)
	router.POST("/users", userscontroller.CreateUser)
}
