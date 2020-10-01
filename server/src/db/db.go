package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
	"github.com/starswirl/go-postgres-api/server/src/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	fmt.Println("init!!!")
	// FIXME: 環境変数化
	db, err = gorm.Open("postgres", "host=postgres user=go-api dbname=postgres password=go-api sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	fmt.Println("close!!!")
	if err := db.Close(); err != nil {
		panic(err)
	}
}
func autoMigration() {
	fmt.Println("autoMigration!!!")
	db.AutoMigrate(&entity.User{})
}
