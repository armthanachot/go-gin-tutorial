package employee

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	conn "user_api/db"
	"user_api/model"
	utils "user_api/utils"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Position  string `json:"position"`
}

type book3color struct {
	Customer_type string `json:"customer_type"`
	Customer_info struct {
		Customer_code    string `json:"customer_code"`
		Customer_road    string `json:"customer_road"`
		Customer_address string `json:"customer_address"`
		Customer_family  struct {
			Mother_name string `json:"mother_name"`
			Dad_name    string `json:"dad_name"`
		} `json:"customer_family"`
		Phone []string `json:"phone"`
	} `json:"customer_info"`
}

// ถ้าจะส่ง data เข้า struct book3color ให้ใช้ key ที่ระบุไว้ใน json
/*
{
    "customer_type":"shop",
    "customer_info":{
        "customer_code":"123123",
        "customer_road":"suwannasorn",
        "customer_address":"bangkok",
        "customer_family":{
            "mother_name":"xx",
            "dad_name":"yyyy"
        },
        "phone":["0894493088","0627422305"]
    }
}
*/

func FindAll(c *gin.Context) {
	var employees []model.Employee
	conn.DB.Select([]string{"id", "firstname", "lastname", "position"}).Find(&employees)
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func Create(c *gin.Context) {
	input := utils.ReadBody(c.Request.Body)
	var employee employee
	json.Unmarshal([]byte(input), &employee)
	initialized := model.Employee{Firstname: employee.Firstname, Lastname: employee.Lastname, Position: employee.Position}
	conn.DB.Create(&initialized)
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func Update(c *gin.Context) {
	input := utils.ReadBody(c.Request.Body)
	var employee_data employee
	json.Unmarshal([]byte(input), &employee_data)
	id := c.Param("id")
	conn.DB.Model(&model.Employee{}).Where("id = ?", id).Updates(employee{Firstname: employee_data.Firstname, Lastname: employee_data.Lastname, Position: employee_data.Position})
	c.JSON(http.StatusOK, gin.H{"messgae": "OK"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	conn.DB.Where("id = ?", id).Delete(&model.Employee{})
}

func Test(c *gin.Context) {
	input := utils.ReadBody(c.Request.Body)
	println(input)
	var book3color_data book3color
	json.Unmarshal([]byte(input), &book3color_data)
	fmt.Println(book3color_data.Customer_info)
	fmt.Println(book3color_data.Customer_info.Customer_family)
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func UploadFile(c *gin.Context) {
	c.Request.ParseForm()
	println(c.PostForm("name"))
	println(c.PostForm("lname"))
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	file_name := utils.FirstString(strings.Split(filename, "."))
	file_type := utils.LastString(strings.Split(filename, "."))
	filename = time.Now().Format("20060102150405") + "_" + file_name + "." + file_type
	out, err := os.Create("public/" + filename)
	// ต้องสร้าง folder public ก่อน
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
