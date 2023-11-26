package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/holdennguyen/vietcslab/controllers"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users/signup", controllers.SignUp())
	route.POST("/users/login", controllers.Login())
	route.POST("/admin/add", controllers.ProductViewerAdmin())
	route.GET("/users/product", controllers.SearchProduct())
	route.GET("/users/search", controllers.SearchProductByQuery())
}