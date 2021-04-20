package pingpong
import (
	"fmt"
	"github.com/gin-gonic/gin"
	utils "user_api/utils"

)
type Pingpong struct {
	ID      string `json:"id" binding:"required" validate:min=10`
	Name    string `json:"name"`
	Message string `json:"message"`
}
var Pingpongs []Pingpong

func CreatePingpong(c *gin.Context) {
	body := utils.ReadBody(c.Request.Body)
	println(body)
	var input Pingpong
	if err := c.ShouldBindJSON(&input); err != nil {
		println(err.Error())
		dataType := fmt.Sprintf("%T", err)
		println(dataType)

		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	Pingpongs = append(Pingpongs, input)
	c.JSON(200, gin.H{
		"message": "pingpong",
		"user":    Pingpongs,
	})
}

func UpdatePingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "updated",
	})
}