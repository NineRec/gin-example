package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/NineRec/gin-example/storage"
)

// AddTodo ...
func AddTodo(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	var todo storage.Todo
	if c.BindJSON(&todo) != nil {
		c.JSON(406, gin.H{"message": "Invalid data", "data": todo})
		c.Abort()
		return
	}

	db.Create(&todo)

	c.JSON(200, gin.H{"message": "Todo created"})
}

// GetTodo ...
func GetTodo(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	var todos []storage.Todo
	db.Where("done = ?", 0).Find(&todos)
	c.JSON(200, gin.H{"todos": &todos})
}
