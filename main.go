package main

import (
	"log"

	"github.com/JulianDavidGamboa/go-db/pkg/product"
	"github.com/JulianDavidGamboa/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(1)

	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
