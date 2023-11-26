package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/holdennguyen/vietcslab/services/backend/controllers"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users/signup", controllers.SignUp())
	route.POST("/users/login", controllers.Login())
	route.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	route.GET("/users/product", controllers.SearchProduct())
	route.GET("/users/search", controllers.SearchProductByQuery())
}