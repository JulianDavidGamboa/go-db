package main

import (
	"github.com/JulianDavidGamboa/go-db/storage"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}
