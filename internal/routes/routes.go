package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Fransiscus-Xaverius/ToDoAPI/internal/routes/todoRoutes"
)

// RegisterRoutes initializes all the routes for the application.
func RegisterRoutes(router *gin.Engine) {
    // Register todos endpoints.
    todoRoutes.RegisterTodoRoutes(router)
}
