package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/Fransiscus-Xaverius/ToDoAPI/db"
	"github.com/Fransiscus-Xaverius/ToDoAPI/models"
)

func main(){
	fmt.Println("Starting server...")
	router := gin.Default()
	db.Init()
	err:= models.Migrate(db.DB)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Server is running on port", port)
    log.Fatal(router.Run(":" + port))

}