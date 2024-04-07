package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/walidsi/go-projects/go-url-shortener-api/package/urlshortener"
)

func init() {
	godotenv.Load()
	urlshortener.Init()
}

func shortenUrl(c *gin.Context) {

	url := c.Query("url")

	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing url!"})
		return
	}

	// If url does not start with a protocol, we will add http as default
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	shortUrl, err := urlshortener.ShortenUrl(url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error shortening the given url",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": shortUrl,
	})
}

func redirectToUrl(c *gin.Context) {

	hash := c.Param("hash")

	if hash == "" {
		domain, _ := os.LookupEnv("DOMAIN")
		errMsg := fmt.Sprintf("No correct url given, should be in form %v/{hash}", domain)

		c.JSON(http.StatusBadRequest, gin.H{"message": errMsg})
		return
	}

	url, err := urlshortener.GetURL(hash)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error getting original url",
		})
		return
	}

	c.Redirect(301, url)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "URL Shortener!",
	})
}

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	router := gin.Default()

	corsConfig := cors.DefaultConfig() // default config should allow localhost
	corsConfig.AllowAllOrigins = true
	//corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	//corsConfig.AllowHeaders = []string{"Origin"}
	// Register the middleware
	router.Use(cors.New(corsConfig))

	router.GET("/", hello)
	router.GET("/shortenurl", shortenUrl)
	router.GET("/:hash", redirectToUrl)

	log.Print("Running on http://localhost:8080")
	router.Run()
}
