package main

import (
	// "fmt"
	// "log"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gusti3111/web-api/accounts"
	"github.com/gusti3111/web-api/handler"
	// "gorm.io/driver/postgres"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "host=localhost user=payment password=payment dbname=payment port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB conn Error")
	}
	// database automiigration
	db.AutoMigrate(&accounts.Person{})
	personRepo := accounts.NewRepo(db)
	// personRepo.Create(person)
	// fileRepo := accounts.NewfileRepo(personRepo)
	personService := accounts.NewServices(personRepo)

	personHandler := handler.NewUserHandler(personService)

	//
	//

	// persons, err := personRepo.FindAll()
	// // for _, person := range persons {
	// // 	fmt.Println("Name:", person.Name)
	// // }

	router := gin.Default()

	// API versioning
	v1 := router.Group("/v1")

	// v1.GET("/", personHandler.RootHandler)
	// v1.GET("/index", personHandler.IndexHandler)
	// v1.GET("/books/:id/:name", personHandler.PathHandler)
	// v1.GET("/query", personHandler.QueryHandler)
	v1.GET("/accounts", personHandler.GetHandler)
	v1.POST("/accounts", personHandler.PostHandler)
	v1.GET("/accounts/:id", personHandler.GetHandlerid)
	v1.PUT("/accounts/:id", personHandler.UpdateHandler)
	v1.DELETE("/accounts/:id", personHandler.DeleteHandler)

	router.Run(":8080")

}
