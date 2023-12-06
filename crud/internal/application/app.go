package application

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/walidsi/go-projects/crud/models"
	"github.com/walidsi/go-projects/crud/mysqldb"
)

func readPost(c *gin.Context) {

	id := c.Param("id")

	db, err := mysqldb.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to the db",
		})
		return
	}

	var blog models.Blog
	result := db.First(&blog, id)
	mysqldb.Close(db)

	if result.Error != nil {
		message := fmt.Sprintf("could not find record with id = %v", id)
		c.JSON(http.StatusNotFound, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             blog.Id,
		"title":          blog.Title,
		"description":    blog.Description,
		"date_published": blog.Date_published,
	})
}

func createPost(c *gin.Context) {

	type RequestBody struct {
		Title          string
		Description    string
		Date_published time.Time
	}

	requestJson := RequestBody{}

	c.Bind(&requestJson)

	db, err := mysqldb.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to the db",
		})
		return
	}

	post := models.Blog{Title: requestJson.Title, Description: requestJson.Description, Date_published: requestJson.Date_published}
	result := db.Create(&post)
	mysqldb.Close(db)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             post.Id,
		"title":          post.Title,
		"description":    post.Description,
		"date_published": post.Date_published,
	})
}

func listPosts(c *gin.Context) {

	db, err := mysqldb.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to the db",
		})
		return
	}

	var posts []models.Blog
	result := db.Find(&posts)
	mysqldb.Close(db)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No records found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func deletePost(c *gin.Context) {

	id := c.Param("id")

	db, err := mysqldb.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to the db",
		})
		return
	}

	result := db.Delete(&models.Blog{}, id)

	if result.Error != nil {
		message := fmt.Sprintf("could not delete record with id=%v", id)
		log.Print(message)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
		})
		return
	}

	mysqldb.Close(db)

	if result.RowsAffected > 0 {
		message := fmt.Sprintf("Successfully deleted record with id=%v", id)
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	} else {
		message := fmt.Sprintf("Record not found, id=%v", id)
		c.JSON(http.StatusNotFound, gin.H{ // http.StatusNoContent could work as well
			"message": message,
		})
	}

}

func updatePost(c *gin.Context) {

	type RequestBody struct {
		Title          string
		Description    string
		Date_published time.Time
	}

	requestJson := RequestBody{}

	c.Bind(&requestJson)

	id := c.Param("id")

	db, err := mysqldb.Open()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error connecting to the db",
		})
		return
	}

	post := models.Blog{}

	result := db.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		mysqldb.Close(db)
		return
	}

	// Update the post with te rquest data
	post.Title = requestJson.Title
	post.Description = requestJson.Description
	post.Date_published = requestJson.Date_published

	result = db.Save(&post)
	mysqldb.Close(db)

	if result.Error != nil {
		message := fmt.Sprintf("could not update record with id = %v", id)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":             post.Id,
		"title":          post.Title,
		"description":    post.Description,
		"date_published": post.Date_published,
	})
}

func Run() error {
	// Force log's color
	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/blog/:id", readPost)
	router.POST("/blog/create", createPost)
	router.GET("/blogs", listPosts)
	router.DELETE("blog/:id", deletePost)
	router.PUT("/blog/:id", updatePost)

	err := router.Run()

	return err
}
