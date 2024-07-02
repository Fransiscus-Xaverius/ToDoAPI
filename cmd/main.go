package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/db"
	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/routes"
	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/migrations"
)

func main(){
	fmt.Println("Starting server...")
	router := gin.Default() //creates a new gin router.
	db.Init() //initialize database connection from the db/db.go file
	err:= migrations.Migrate(db.DB) //migrates models from the models/todo.go file
	if err != nil {  //check for migrate errors
		log.Fatalf("failed to migrate: %v", err)
	}

	router.Use(cors.New(cors.Config{ //add cors middleware to the router
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

	routes.RegisterRoutes(router) //register routes from the internal/routes/todoRoutes.go file


	port := os.Getenv("PORT") //get port from .env file
    if port == "" {
        port = "8080" //if port is missing, use 8080 as default
    }

    log.Println("Server is running on port", port) //print server running message
    log.Fatal(router.Run(":" + port)) //Start server at port, and print an error if it fails

}