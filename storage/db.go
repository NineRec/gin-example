package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DB is for get the db connection
func DB(connString string) gin.HandlerFunc {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
