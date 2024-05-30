package server

import (
	"api-server/config"
	"api-server/middlewares"
	"github.com/gin-gonic/gin"

	cUser "api-server/controllers/user"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			if config.ServerDebug {
				debug := v1.Group("/debug")
				{
					debug.POST("/token", cUser.TemporaryTokenRouter)
				}
			}

			user := v1.Group("/user")
			{
				user.Use(middlewares.JwtAuthMiddleware)

				user.GET("/", cUser.GetLoggedInUserRouter)
				user.GET("/:id", cUser.GetUserRouter)

				user.POST("/", cUser.CreateUserRouter)

				user.PATCH("/", cUser.UpdateUserRouter)
			}
		}
	}

	return r
}
