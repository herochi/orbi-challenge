package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	rest "github/herochi/orbi/service-a/adapter/http/rest"
)

func (g *ginServer) MapRoutes() {
	g.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization", "Cache-Control", "Pragma", "Expires"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	apiV1 := g.Router.Group("orbi-api/v1/a")

	g.userRoutes(apiV1)

}

/*func (g *ginServer) loginRoutes(api *gin.RouterGroup) {
	authRoutes := api.Group("/login")
	{
		authHandler := rest.NewLoginHandler(
			g.Container.Service.AuthService,
			g.Container.Service.UserService)

		authRoutes.POST("/", authHandler.Login)
		authRoutes.POST("/token/refresh", authHandler.RefreshToken)
	}
}*/

func (g *ginServer) userRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		userHandler := rest.NewUserHandler(g.Container.Service.UserService)

		userRoutes.POST("/", userHandler.Create)
		userRoutes.GET("/:id", userHandler.GetById)
		userRoutes.PATCH("/:id", userHandler.UpdateUser)
	}
}
