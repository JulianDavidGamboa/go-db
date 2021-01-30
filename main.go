package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/JulianDavidGamboa/go-db/pkg/product"
	"github.com/JulianDavidGamboa/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(6)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println("No hay un producto con este id")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(m)
	}

}
