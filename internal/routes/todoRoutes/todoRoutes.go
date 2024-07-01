package todoRoutes

import (
    "github.com/gin-gonic/gin"
    "github.com/Fransiscus-Xaverius/ToDoAPI/internal/handlers"
)

// RegisterTodoRoutes sets up the routes for todo-related operations.
func RegisterTodoRoutes(router *gin.Engine) {
    // Define a group for todo routes
    todos := router.Group("/todo")
    {
        // Route for creating a new todo
        todos.POST("/", handlers.CreateTodoHandler)

        // Route for listing all todos
        todos.GET("/all", handlers.ListTodosHandler)

        // Route for retrieving a specific todo by ID
        // todos.GET("/:id", handlers.GetTodoHandler)

        // Route for updating a specific todo by ID
        // todos.PUT("/:id", handlers.UpdateTodoHandler)

        // Route for deleting a specific todo by ID
        // todos.DELETE("/:id", handlers.DeleteTodoHandler)
    }
}
