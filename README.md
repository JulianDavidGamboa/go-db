# go-db

Proyecto básico para conectar una base de datos de postgresql y MySql con Go.

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

# Obtener todos los productos

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

ms, err := serviceProduct.GetAll()

if err != nil {
    log.Fatalf("product.GetAll: %v", err)
}

fmt.Println(ms)
```

# Obtener producto por un id

```go
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
```

# Actualizar producto

```go
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
```

# Eliminar un producto

```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

err := serviceProduct.Delete(1)

if err != nil {
	log.Fatalf("product.Delete: %v", err)
}

```

# Transactions

```go
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
```
