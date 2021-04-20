package user

import (
	// "log"
	"encoding/json"
	"net/http"

	// "os"
	"fmt"
	conn "user_api/db"
	"user_api/model"
	utils "user_api/utils"

	"github.com/gin-gonic/gin"
)

type users struct {
	Name     string
	Age      int
	Favorite []string
}
type userInput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        string `json:"age"`
	Email      string `json:"email"`
}

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

// func FindAll(c *gin.Context) {
// 	users := []users{
// 		{
// 			Name:     "Thanachot",
// 			Age:      24,
// 			Favorite: []string{"art", "programming", "food"},
// 		},
// 		{
// 			Name:     "Ammara",
// 			Age:      24,
// 			Favorite: []string{"art", "programming", "food"},
// 		},
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }

func Create(c *gin.Context) {
	input := utils.ReadBody(c.Request.Body)

	// if err := c.ShouldBindJSON(input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	// 	return
	// }
	var read userInput
	res := json.Unmarshal([]byte(input), &read)
	fmt.Printf("Name: %s \n", read.First_name)
	fmt.Printf("%+v\n", read)
	println(res)
	user_data := model.User{FirstName: read.First_name, LastName: read.Last_name, Age: read.Age, Email: read.Email}
	conn.DB.Create(&user_data)
	// ต้อง insert ด้วย address ของ data
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func FindAll(c *gin.Context) {
	var users []model.User
	conn.DB.Select([]string{"id", "first_name", "last_name", "age", "email"}).Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
	return
}
