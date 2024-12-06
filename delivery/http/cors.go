package http

import (
	"time"

	"github.com/gin-contrib/cors"
)

func initCors() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

	return corsConfig
}
