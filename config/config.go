package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"log"
	"os"
)

var Port = func() string {
	return "0.0.0.0:" + os.Getenv("PORT")
}()

var Cors = cors.New(cors.Config{
	AllowMethods: "GET, POST, OPTIONS, PUT, DELETE, PATCH",
	AllowOrigins: "*",
})

var Db = func() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}()

var ApiSecret = func() string {
	return os.Getenv("API_SECRET")
}()

var PublicDir = func() string {
	path := os.Getenv("PUBLIC_DIR")
	if path == "" {
		return "public"
	}
	return path
}()

var BaseUrl = func() string {
	url := os.Getenv("BASE_URL")
	if url == "" {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}
		hostname = fmt.Sprintf("http://%s", hostname)
		log.Printf("BASE_URL is not set, default to %s", hostname)
		return hostname
	}
	return url
}()
