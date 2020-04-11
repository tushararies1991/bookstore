package app

import (
	"bookstore_users/controller/ping"
	"bookstore_users/controller/route"
	"bookstore_users/controller/user"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	r.POST("/user/login", user.Login)

	r.POST("/user/create", user.CreateUser)
	r.GET("/user/:user_id", user.GetUser)
	r.PUT("/user/:user_id", user.UpdateUser)
	r.PATCH("/user/:user_id", user.UpdateUser)
	r.DELETE("/user/:user_id", user.DeleteUser)
	r.GET("/internal/users/search", user.FindByStatus)

	r.POST("/route/search", route.SearchRoute)
}
