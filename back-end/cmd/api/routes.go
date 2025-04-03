package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {

	router := gin.Default()

	// CORS middleware
	router.Use(app.enableCORS())

	// Public routes
	router.GET("/", app.Home)
	router.POST("/authenticate", app.authenticate)
	//router.GET("/refresh", app.refreshToken)
	//router.GET("/logout", app.logout)

	router.GET("/users", app.AllUsers)
	router.GET("/restaurants", app.AllRestaurants)
	router.GET("/restaurants/:id", app.RestaurantDetails)
	router.POST("/restaurants", app.InsertRestaurant)
	//router.GET("/restaurants/:id", app.GetRestaurant)

	// Admin routes with authentication
	admin := router.Group("/admin")
	admin.Use(app.authRequired())
	{
		//admin.GET("/restaurants", app.RestaurantCatalog)
		//admin.GET("/restaurants/:id", app.RestaurantForEdit)
		//admin.PUT("/restaurants/0", app.InsertRestaurant)
		//admin.PATCH("/restaurants/:id", app.UpdateRestaurant)
		//admin.DELETE("/restaurants/:id", app.DeleteRestaurant)
	}

	return router
}
