package postgres

import (
	"context"
	"errors"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func (db *serviceDB) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	tx := db.DB.Begin()

	result := tx.Create(order)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	for _, orderLine := range order.OrderLines {
		product, err := db.GetProduct(ctx, orderLine.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if product.Quantity < orderLine.Quantity {
			tx.Rollback()
			return nil, errors.New("insufficient quantity in stock")
		}

		product.Quantity -= orderLine.Quantity
		if _, err := db.UpdateProduct(ctx, product); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return order, nil
}

func (db *serviceDB) GetOrder(ctx context.Context, orderID string) (*model.Order, error) {
	var order model.Order
	result := db.DB.Preload("OrderLines.Product").First(&order, orderID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (db *serviceDB) UpdateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	result := db.DB.Save(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (db *serviceDB) DeleteOrder(ctx context.Context, orderID string) error {
	result := db.DB.Where("id = ?", orderID).Delete(&model.Order{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (db *serviceDB) ListOrders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	result := db.DB.Preload("OrderLines.Product").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
