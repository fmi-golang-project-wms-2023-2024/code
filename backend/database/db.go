package database

import (
	"context"

	"github.com/nikola-enter21/wms/backend/database/model"
)

type DB interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, userID string) (*model.User, error)
	GetUserByCredentials(ctx context.Context, username string, password string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userID string) error
	ListUsers(ctx context.Context) ([]*model.User, error)

	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetProduct(ctx context.Context, productID string) (*model.Product, error)
	UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, productID string) error
	ListProducts(ctx context.Context) ([]*model.Product, error)

	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, orderID string) (*model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
	ListOrders(ctx context.Context) ([]*model.Order, error)

	CreateInvoice(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error)
	GetInvoice(ctx context.Context, invoiceID string) (*model.Invoice, error)
	UpdateInvoice(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error)
	DeleteInvoice(ctx context.Context, invoiceID string) error
	ListInvoices(ctx context.Context) ([]*model.Invoice, error)
}
