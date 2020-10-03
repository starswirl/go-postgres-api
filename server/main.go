package main

import (
	"github.com/starswirl/go-postgres-api/server/src/db"
	"github.com/starswirl/go-postgres-api/server/src/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
