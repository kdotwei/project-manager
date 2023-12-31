package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	username string
	password string
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("Username can not be empty"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("Password can not be empty"),
		})
		return
	}
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "Login Success",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Login Failed, Username or Password is not correct",
		})
		return
	}
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")

	// Read the static files
	server.Static("/assets", "./template/assets")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":8888")
}
