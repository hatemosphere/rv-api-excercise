package users

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/hatemosphere/rv-api-excercise/api/models"
)

// GetOne function fetches user by username from DB and tells him about his birthday
func GetOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var user models.User

	query := bson.M{"username": c.Param("username")}
	err := db.C(models.CollectionUsers).Find(query).One(&user)
	if err != nil {
		c.Error(err)
		if err == mgo.ErrNotFound {
			c.Status(404)
		} else {
			c.Status(400)
		}
		return
	}
	userDateOfBirth, ok := DateStringProcessor(dateLayout, user.DateOfBirth)
	if !ok {
		log.Printf("Problem converting date of birthday string for user %s:", user.Username)
		c.Status(400)
		return
	}

	message, ok := BirthDayPrinter(user.Username, userDateOfBirth)
	if !ok {
		log.Printf("Problem printing birthday message for user %s:", user.Username)
	}

	// c.JSON(200, message)
	c.JSON(200, gin.H{
		"message": message,
	})
}

// CreateOrUpdate function - non-CRUD upsert operation for user
func CreateOrUpdate(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		c.Status(400)
		return
	}

	userDateOfBirth, ok := DateStringProcessor(dateLayout, user.DateOfBirth)
	if !ok {
		log.Printf("Problem converting date of birthday string for user %s:", user.Username)
		c.Status(400)
		return
	}

	dateDifference := DateDiffer(userDateOfBirth)
	if dateDifference < 1 {
		log.Printf("%s is not born yet or born today", user.Username)
		c.Status(400)
		return
	}

	query := bson.M{"username": c.Param("username")}
	doc := bson.M{
		"username":    user.Username,
		"dateOfBirth": user.DateOfBirth,
	}
	_, err = db.C(models.CollectionUsers).Upsert(query, doc)
	if err != nil {
		c.Error(err)
		c.Status(400)
		return
	}

	c.Status(204)
}
