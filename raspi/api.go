package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type data struct {
	gorm.Model
	message string
}

func main() {
	r := gin.Default()
	db, err := gorm.Open("sqlite3", "./db/test.db")
	if err != nil {
		log.Fatal("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&data{})
	r.GET("/ping", func(c *gin.Context) {
		result := db.Create(&data{message: "pong"})
		fmt.Println(result.Error)
		if result.Error == nil {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "すまん",
			})
		}
	})
	r.Run()
}
