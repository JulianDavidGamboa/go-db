package main

import (
	"fmt"
	"log"

	"github.com/JulianDavidGamboa/go-db/pkg/product"
	"github.com/JulianDavidGamboa/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()

	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)

}
