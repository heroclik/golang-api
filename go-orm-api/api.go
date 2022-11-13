package main

import (
	"example/go-orm-api/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	user()
}

func user() {
	//dbconnect
	dsn := "root@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//get users data
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(200, users)
	})
	//get users data by id
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user []model.User
		db.First(&user, id)
		c.JSON(200, user)
	})
	//create user data
	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&user)
		c.JSON(200, gin.H{"RowAffected": result.RowsAffected})
	})
	//delete users data
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user []model.User
		db.First(&user, id)
		db.Delete(&user)
	})
	//update users data
	r.PUT("/users", func(c *gin.Context) {
		var user model.User
		var updatedUser model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updatedUser, user.ID)
		updatedUser.Username = user.Username
		updatedUser.Fname = user.Fname
		updatedUser.Lname = user.Lname
		db.Save(updatedUser)
		c.JSON(200, updatedUser)
	})
	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
