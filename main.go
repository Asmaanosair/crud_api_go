// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	Id     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Status bool   `json:"status"`
// }

// var users = []User{
// 	{Id: 1, Name: "Asmaa", Status: true},
// 	{Id: 2, Name: "Mohamed", Status: true},
// 	{Id: 3, Name: "Fatma", Status: false},
// }

//	func main() {
//		router := gin.Default()
//		router.GET("/users", getUsers)
//		router.Run("localhost:9090")
//	}
//
//	func getUsers(context *gin.Context) {
//		context.IndentedJSON(http.StatusOK, users)
//	}
package main

import (
	"goTask/config"
	"goTask/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run("localhost:9090")
}
