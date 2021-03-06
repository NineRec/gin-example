package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/NineRec/gin-example/handlers/todo"
)

// Route generates all the router for http server
func Route(r *gin.Engine) *gin.Engine {
	// Get list
	r.GET("/todo", todo.GetTodo)

	r.POST("/todo", todo.AddTodo)

	return r
}
