package handlers

import (

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/models"
	"github.com/Fransiscus-Xaverius/ToDoAPI/internal/db"
	"net/http"
	// "time"
)

func CreateTodoHandler(c *gin.Context) {
	
	//get data from body
	var body struct {
        UserID      string `json:"user_id" binding:"required"`
        Title       string `json:"title" binding:"required"`
        Description string `json:"description" binding:"required"`
    }
	//bind data
    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	//fmt.Println(body)

	todo := models.Todo{
        UserID:      body.UserID,
        Title:       body.Title,
        Description: body.Description,
        Completed:   false,
        CompletedAt: nil,
    }
	fmt.Println("Creating New Todo")
    result := db.DB.Create(&todo)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully", "todo": todo})
}

func ListTodosHandler(c *gin.Context) {
	var body struct {
		UserID      string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil { //for some reason this doesn't work with Bind() ?? ok..
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(body)

	fmt.Println("Listing all todos")

	var todos []models.Todo

	db.DB.Where("user_id = ?", body.UserID).Find(&todos) //find todos of user by id

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func EditTodoHandler(c *gin.Context){
	var body struct {
		ID          uint `json:"id" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//define todo model
	var todo models.Todo

	//find todo, if not found then return 404
	err := db.DB.Where("id = ?", body.ID).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	//update todo
	todo.Title = body.Title
	todo.Description = body.Description
	db.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": todo})

}