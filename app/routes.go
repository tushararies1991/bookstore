package app

import (
	"bookstore/controller/ping"
	"bookstore/controller/route"
	"bookstore/controller/user"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	r.POST("/user/create", user.CreateUser)
	r.GET("/user/:user_id", user.GetUser)
	r.PUT("/user/:user_id", user.UpdateUser)
	r.PATCH("/user/:user_id", user.UpdateUser)
	r.DELETE("/user/:user_id", user.DeleteUser)
	r.GET("/internal/users/search", user.FindByStatus)

	r.POST("/route/search", route.SearchRoute)
}
