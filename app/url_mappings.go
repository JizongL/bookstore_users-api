package app

import (
	"github.com/JizongL/bookstore_users-api/controllers/ping"
	user "github.com/JizongL/bookstore_users-api/controllers/users"
)



func mapUrl(){
	router.GET("/ping",ping.Ping)		
	router.GET("/users/:user_id",user.GetUser)
	router.POST("/users",user.CreateUser)
}