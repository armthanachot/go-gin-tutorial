package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	core_controller "user_api/controllers/core"
	employee_controller "user_api/controllers/employee"
	main_controller "user_api/controllers/main_controller"
	pingpong_controller "user_api/controllers/pingpong"
	user_controller "user_api/controllers/user"
	db "user_api/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "go_tutorial"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

// user_api คือ ชื่อ module ที่เราสร้างไว้ ตอน go mod init <module_name>

func main() {
	r := gin.Default()
	r.Use(RequestLogger())
	db.ConnectDatabase()

	r.GET("/main", main_controller.FindAll)
	userRoute := r.Group("/user")
	{
		userRoute.GET("", user_controller.FindAll)
		userRoute.POST("", user_controller.Create)
	}
	// context ใช้รับค่าต่างๆ
	pingpongRoute := r.Group("/pingpong")
	{
		pingpongRoute.POST("", testMiddle, pingpong_controller.CreatePingpong)
		pingpongRoute.PUT("/:id", pingpong_controller.UpdatePingPong)
	}

	employeeRoute := r.Group("/employee")
	{
		employeeRoute.GET("", employee_controller.FindAll)
		employeeRoute.POST("", employee_controller.Create)
		employeeRoute.POST("/test", employee_controller.Test)
		employeeRoute.POST("/fileupload", employee_controller.UploadFile)
		employeeRoute.PUT("/:id", employee_controller.Update)
		employeeRoute.DELETE("/:id", employee_controller.Delete)

	}

	coreRoute := r.Group("/core")
	{
		coreRoute.POST("/genToken", core_controller.Login)
		coreRoute.POST("/verify", core_controller.VerifyToken)
	}
	r.StaticFS("/file", http.Dir("public"))
	// ต้อง set static file หากไม่่ set จะไม่สามารถเข้าถึงไฟล์ได้
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.
		c.Request.Body = rdr
		c.Next()
	}
}

func testMiddle(c *gin.Context) {
	println("test middle")
	c.Next()
}

type UserHandler struct {
	DB *gorm.DB
}
