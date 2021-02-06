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
		ID:           2,
		Name:         "Programacion con GO",
		Observations: "Programming with go",
		Price:        125,
	}

	err := serviceProduct.Update(m)

	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
