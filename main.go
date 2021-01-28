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

	m := &product.Model{
		Name:         "Desarrollo de db con Go",
		Price:        70,
		Observations: "On Fire",
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

}
