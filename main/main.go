package main

import (
	"fmt"
	"os"

	"github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/models"
	"github.com/gin-gonic/gin"
)

func initializer() error {
	fmt.Println("Init...")
	if err := initializers.LoadEnvVariable(); err != nil {
		return err
	}
	if err := initializers.ConnectoToDB(); err != nil {
		return err
	}
	return nil
}

func pong(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"pong",
	})
}


func register(c *gin.Context){
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password")
		if err := models.CreateUser(initializers.DB, username, email, password); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User created successfully",
		})
}


func main() {
	initializer()
	r := gin.Default()
	r.GET("/ping",pong)
	r.POST("/register",register)
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		user, err := models.GetUser(initializers.DB, username)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		if user.Password != password {
			c.JSON(400, gin.H{
				"message": "Invalid password",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Logged in successfully",
		})
	})
	r.POST("/create", func(c *gin.Context) {
		linkly := c.PostForm("linkly")
		redirectURL := c.PostForm("redirect")
		user_id := c.PostForm("user_id")
		if err := models.CreateLink(initializers.DB, linkly, redirectURL, user_id); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Link created successfully",
		})
	})

	r.GET("/links", func(c *gin.Context) {
		links, err := models.GetLinkByUser(initializers.DB)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"links": links,
		})
	})
	r.GET("/links/:linkly", func(c *gin.Context) {
		linkly := c.Param("linkly")
		link, err := models.GetLink(initializers.DB, linkly)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"link": link,
		})
	})
	r.PUT("/links/:linkly", func(c *gin.Context) {
		linkly := c.Param("linkly")
		redirectURL := c.PostForm("redirect")
		if err := models.UpdateLink(initializers.DB, linkly, redirectURL); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Link updated successfully",
		})
	})
	r.DELETE("/links/:linkly", func(c *gin.Context) {
		linkly := c.Param("linkly")
		if err := models.DeleteLink(initializers.DB, linkly); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Link deleted successfully",
		})
	})

	r.Run(os.Getenv("PORT"))
}
