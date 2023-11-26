package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/holdennguyen/vietcslab/services/backend/controllers"
	"github.com/holdennguyen/vietcslab/services/backend/database"
	"github.com/holdennguyen/vietcslab/services/backend/middleware"
	"github.com/holdennguyen/vietcslab/services/backend/routes"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := []byte(`{"status": "ok"}`)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		w.Write(resp)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
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
	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}