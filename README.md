# go-db

## Migrar tabla de products

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
    log.Fatalf("product.Migrate: %v", err)
	}
```

## Migrar tabla de invoiceheaders

```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
    log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

## Migrar tabla de invoiceitems
```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
    log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

# Crear un producto
```go
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
```