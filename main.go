package main

import (
	"log"

	"github.com/JulianDavidGamboa/go-db/pkg/invoice"
	"github.com/JulianDavidGamboa/go-db/pkg/invoiceheader"
	"github.com/JulianDavidGamboa/go-db/pkg/invoiceitem"
	"github.com/JulianDavidGamboa/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Julian Gamboa",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
			&invoiceitem.Model{ProductID: 3},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)

	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
