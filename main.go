package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/holdennguyen/vietcslab/controllers"
	"github.com/holdennguyen/vietcslab/database"
	"github.com/holdennguyen/vietcslab/middleware"
	"github.com/holdennguyen/vietcslab/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	route := gin.New()
	route.Use(gin.Logger())

	routes.UserRoutes(route)
	route.Use(middleware.AuthMiddleware())

	route.GET("/addtocart", app.AddToCart())
	route.GET("removeitem", app.RemoveItem())
	route.GET("/cartcheckout", app.BuyFromCart())
	route.GET("/instantbuy", app.InstantBuy())

	log.Fatal(route.Run(":" + port))
}