package router

import (
	"github.com/gin-gonic/gin"

	u "ilyaryabchinski/gotask/src/users"
)

func BindRoutes() *gin.Engine {
	r := gin.Default()
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", u.GetUsers)
		userRoutes.GET("/:personalCode", u.GetUser)
		userRoutes.POST("/", u.CreateUser)
		userRoutes.PUT("/:personalCode", u.UpdateUser)
		userRoutes.DELETE("/:personalCode", u.DeleteUser)
	}
	return r
}
