package config

import (
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
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
