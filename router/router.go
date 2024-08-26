package router

import "github.com/gin-gonic/gin"

func Router() {
	engine := gin.Default()
	engine.Run("")
}
