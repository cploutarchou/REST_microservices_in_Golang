package app

import (
	"blog/controllers/ping_controller"
	"blog/controllers/users_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)
	router.GET("/users/:user_id", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
}
