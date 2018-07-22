package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/NineRec/gin-example/handlers"
	"github.com/NineRec/gin-example/storage"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	connString := "user:password@/todolist?charset=utf8&parseTime=True"
	r.Use(storage.DB(connString))

	return r
}

func main() {
	r := handlers.Route(setupRouter())

	// Listen and Server in 0.0.0.0:8080
	endless.ListenAndServe(":8080", r)
}
