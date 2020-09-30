package db

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

var (
    db  *gorm.DB
    err error
)

// Init is initialize db from main function
func Init() {
	// FIXME: 環境変数化とポートの値確認
    db, err = gorm.Open("postgres", "host=0.0.0.0 port=8888 user=go-api dbname=go-api password=go-api sslmode=disable")
    if err != nil {
        panic(err)
    }
}

// GetDB is called in models
func GetDB() *gorm.DB {
    return db
}

// Close is closing db
func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}
