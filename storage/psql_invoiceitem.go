package storage

import (
	"database/sql"
	"log"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS go_db.invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
		(invoice_header_id) REFERENCES go_db.invoice_headers (id) ON UPDATE 
		RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_invoice_product_id_fk FOREIGN KEY (product_id) 
		REFERENCES go_db.products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

// PsqlInvoiceItem used for work with postgres - invoice_item
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer of PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implement the interface invoice_item.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migración de invoiceItem ejecutada correctamente")

	return nil
}
