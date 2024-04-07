package urlshortener

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bluele/gcache"
	"github.com/walidsi/go-projects/go-url-shortener-api/package/sqlitedb"
	"gorm.io/gorm"
)

const hashLength int = 16

type UrlMap struct {
	gorm.Model
	Url  string
	Hash string
}

var gc gcache.Cache

func Init() error {

	gc = gcache.New(20).LRU().Build()

	migrated, _ := os.LookupEnv("MIGRATED")

	if migrated == "no" {
		db, err := sqlitedb.Open()

		if err != nil {
			return err
		}

		db.AutoMigrate(&UrlMap{})

		sqlitedb.Close(db)

		os.Setenv("MIGRATED", "yes")
	}

	return nil
}

func generateHashString(input string, length int) string {
	//rand.Seed(time.Now().UnixNano())
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(gen.Intn(26) + 65) // A-Z characters
	}
	return string(bytes)
}

func ShortenUrl(url string) (string, error) {
	// Generate random string
	hash := generateHashString(url, hashLength)

	// Store url and the random string
	db, err := sqlitedb.Open()

	if err != nil {
		return "", err
	}

	urlMap := UrlMap{Url: url, Hash: hash}

	result := db.Create(&urlMap)

	if result.Error != nil {
		return "", result.Error
	}

	domain, _ := os.LookupEnv("DOMAIN")
	shortUrl := domain + "/" + hash

	// return the random string
	return shortUrl, nil
}

func GetURL(hash string) (string, error) {
	start := time.Now()

	fromCache := false

	defer func() {
		elapsed := time.Since(start)
		if fromCache {
			log.Println("Time taken from cache:", elapsed)
		} else {
			log.Println("Time taken from db:", elapsed)
		}
	}()

	// Check if hash is in cache
	url, err := gc.Get(hash)
	if err == nil {
		fromCache = true
		return url.(string), nil
	}

	// Open the db
	db, err := sqlitedb.Open()

	if err != nil {
		return "Error opening db", err
	}

	record := UrlMap{}

	// Get the record containing the hash
	result := db.Where("hash = ?", hash).First(&record)

	if result.Error != nil {
		return "Error retrieving url for given hash", result.Error
	}

	// Store in cache
	gc.Set(hash, record.Url)

	// return the random string
	return record.Url, nil
}
