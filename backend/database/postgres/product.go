package postgres

import (
	"context"
	"errors"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func (db *serviceDB) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	result := db.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (db *serviceDB) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	var product model.Product
	result := db.DB.Where("id = ?", productID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (db *serviceDB) UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	result := db.DB.Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (db *serviceDB) DeleteProduct(ctx context.Context, productID string) error {
	result := db.DB.Where("id = ?", productID).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (db *serviceDB) ListProducts(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	result := db.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
