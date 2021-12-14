package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/heroku/go-getting-started/mappings"
)

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run()
}
