package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hatemosphere/rv-api-excercise/api/db"
)

func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
