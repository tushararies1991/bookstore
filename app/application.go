package app

import (
	"github.com/gin-gonic/gin"
)

func StartApplication() {
	r := gin.Default()
	registerRoutes(r)
	r.Run()
}
