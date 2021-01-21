package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pretty "github.com/inancgumus/prettyslice"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.Static("/css", "./templates/css")

	userRoutes := r.Group("/users")
	{

		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)

	}

	viewRoutes := r.Group("/index")
	{

		viewRoutes.GET("/", GetAllUser)

	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}

}

func GetUsers(c *gin.Context) {

	c.JSON(200, Users)

}

func GetAllUser(c *gin.Context) {

	pretty.Show("Users", Users)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "All Users",
		"Users": Users,
	})
}

func CreateUser(c *gin.Context) {

	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {

		c.JSON(422, gin.H{

			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)

	c.JSON(200, gin.H{

		"error": false,
	})
}

func EditUser(c *gin.Context) {

	id := c.Param("id")

	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{

				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})

}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {

			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"messgae": "invalid user id",
	})
}
