package main

//belom config gin nya diterminal

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//inisialisasi gin router
	router := gin.Default()

	//middleware logger
	router.Use(gin.Logger())

	//middleware Recovery
	router.Use(gin.Recovery())

	//Route definition
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/halo/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "halo " + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"Password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		//disini bisa lakuin validasi login, kaya meriksa didatabase.
		//contoh sederhana meriksa apakah email dan password cocok.
		if loginData.Email == "xxx.gmail.com" && loginData.Password == "admin123pdn" {
			c.JSON(200, gin.H{
				"message": "Login succesful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}

	})

	//menambahkan endpoint  untuk mengambil parameter query
	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Name parameter is missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello " + name + " !",
		})
	})

	//menjalankan server
	router.Run(":8080")
}
