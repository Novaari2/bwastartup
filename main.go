package main

import (
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//if err != nil {
	//	panic(err)
	//}
	//
	//var users []user.User
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//}

	router := gin.Default()
	router.GET("/users", handler)

	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
