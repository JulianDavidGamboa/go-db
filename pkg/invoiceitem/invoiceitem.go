package invoiceitem

import "time"

// Model if invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
