package http

import (
	"fmt"
	"message-processor/internal/settings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.New()

	r.Use(cors.Default(),
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	r.GET("/health", Health)
	v1(r)

	settings := settings.GetSettings()
	address := fmt.Sprintf("%s:%s", settings.Host, settings.Port)

	r.Run(address)
}

func v1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
	}
}
