package main_controller

import (
	// "fmt"
	// "io"
	// "log"
	"net/http"
	// "os"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

// ค้องขึ้นต้นชื่อ function ด้วย FindAll เพื่อให้ไฟล์อื่น import ไปใช้ได้ ถ้าใช้ใน file เดียวกันถึงจะใช้พิมพ์เล็กได้