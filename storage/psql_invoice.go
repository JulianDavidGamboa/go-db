package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JulianDavidGamboa/go-db/pkg/invoice"
	"github.com/JulianDavidGamboa/go-db/pkg/invoiceheader"
	"github.com/JulianDavidGamboa/go-db/pkg/invoiceitem"
)

// PsqlInvoice used for work with postgres - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice return a new pointer of PsqlInvouce
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}

	log.Printf("Factura creada con id: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}

	log.Printf("Items creados: %d \n", len(m.Items))

	return tx.Commit()

}
