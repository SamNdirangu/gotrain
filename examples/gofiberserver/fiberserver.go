package gofiberserver

import (
	"log"
	"training/examples/gofiberserver/apis"
	"training/examples/gofiberserver/apis/books"
	"training/examples/gofiberserver/apis/products"
	"training/examples/gofiberserver/apis/test"
	"training/examples/gofiberserver/config"
	"training/examples/gofiberserver/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	*fiber.App
}

func StartServer() {
	//Get our environment config varirables ************************************
	config := config.New()

	// Connect and Init our database ************************************************
	db, err := database.ConnectDB(config)
	if err != nil {
		//No need to start server if we cant connect to the database
		log.Fatal("failed to connect to database:", err.Error())
	}
	initDB(db) //Initialize db models and automigration

	// Setup our server app ***************************************************
	app := App{
		fiber.New(*config.GetFiberConfig()),
	}
	// Add our middlewares as needed for the environment
	app.registerMiddlewares(config)

	// Register API Endpoints **************************************************
	app.Get("/test", test.HeartBeat) //Testing endpoint if server is alive

	// v1 API Endpoints
	apiV1 := app.Group("/api/v1") //version 1.0 API route
	apis.RegisterV1APIs(apiV1)    //Register API 1.0 routes

	// Register Error handlers
	app.Use(errorHandlers.NotFoundHandler) // 404 not found error handler

	// Start the Server *******************************************************
	log.Fatal(app.Listen(config.GetString("APP_PORT")))
}

func initDB(db *gorm.DB) {
	db := database.GetDB()

	db.AutoMigrate(&books.Book{})
	db.AutoMigrate(&products.Product{})
}
