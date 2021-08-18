package http

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.New()

	r.Use(gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	r.GET("/health", Health)
	v1(r)

	r.Run()
}

func v1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
	}
}
