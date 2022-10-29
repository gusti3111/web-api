package main

import (
	"web-api-golang-myself/controllers"
	"web-api-golang-myself/model"

	"github.com/gin-gonic/gin"
)

// var products = []model.Input{
// 	{Name: "Bedak Johnson&Johnson", Price: 1000, Rating: 10, Stock: 100},
// }

func main() {
	model.Conn()

	// product := model.Product{}
	// product.Name = "Sabun Asepso"
	// product.Price = 10000
	// product.Stock = 100
	// product.Rating = 5
	// err = db.Create(&product).Error
	// if err != nil {
	// 	return
	// }

	router := gin.Default()
	router.GET("/user", controllers.GetAllUser)
	router.GET("/user/:id", controllers.GetUserById)
	router.POST("/user/", controllers.CreateUser)
	router.PUT("user/:id", controllers.UpdateUser)
	router.DELETE("user/:id", controllers.DeleteUser)

	// router.GET("/user/:id",get)
	router.Run(":9000")

}
