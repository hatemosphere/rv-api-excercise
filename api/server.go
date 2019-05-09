package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hatemosphere/rv-api-excercise/api/db"
	"github.com/hatemosphere/rv-api-excercise/api/handlers/users"
	"github.com/hatemosphere/rv-api-excercise/api/middlewares"
	_ "go.uber.org/automaxprocs"
)

const (
	port      = "3000"
	prefix = "/hello"
)

func init() {
	db.Connect()
}

func main() {
	binding.Validator = new(defaultValidator)

	router := gin.Default()

	router.Use(middlewares.Connect)
	router.GET(prefix + "/:username", users.GetOne)
	router.PUT(prefix + "/:username", users.CreateOrUpdate)

	router.Run(":" + port)
}
