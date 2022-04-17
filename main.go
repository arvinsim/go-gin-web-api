package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"runtime"
)

type Item struct {
	gorm.Model
	Name string
}

func main() {
	// Setup database
	dsn := "host=localhost user=metheuser password=mysecretpassword dbname=mydb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("There is an error connecting to the database")
	}

	// Setup Router
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})

	router.GET("/echo/:val", func(c *gin.Context) {
		val := c.Param("val")
		c.String(200, val)
	})

	router.GET("/items/all", func(c *gin.Context) {
		var items []Item
		db.Find(&items)

		c.JSON(200, gin.H{
			"data": items,
		})
	})

	router.Run(":8000")
}
