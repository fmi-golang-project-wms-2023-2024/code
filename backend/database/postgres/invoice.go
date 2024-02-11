package postgres

import (
	"context"
	"errors"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func (db *serviceDB) CreateInvoice(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	result := db.DB.Create(invoice)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoice, nil
}

func (db *serviceDB) GetInvoice(ctx context.Context, invoiceID string) (*model.Invoice, error) {
	var invoice model.Invoice
	result := db.DB.Preload("Order").First(&invoice, invoiceID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &invoice, nil
}

func (db *serviceDB) UpdateInvoice(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	result := db.DB.Save(invoice)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoice, nil
}

func (db *serviceDB) DeleteInvoice(ctx context.Context, invoiceID string) error {
	result := db.DB.Where("id = ?", invoiceID).Delete(&model.Invoice{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("invoice not found")
	}
	return nil
}

func (db *serviceDB) ListInvoices(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	result := db.DB.Preload("Order").Preload("Items").Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoices, nil
}
