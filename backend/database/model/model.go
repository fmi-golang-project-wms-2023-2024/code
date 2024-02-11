package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New().String()
	return nil
}

type Role string

const (
	Admin Role = "admin"
	Staff Role = "staff"
)

type User struct {
	Base
	FullName string `gorm:"type:text;not null"`
	Username string `gorm:"type:text;not null;unique"`
	Password string `gorm:"type:text;not null"`
	Role     Role   `gorm:"type:text;not null"`
}

type Product struct {
	Base
	SKU      string `gorm:"type:text;not null"`
	Title    string `gorm:"type:text;not null"`
	Price    string `gorm:"type:text;not null"`
	Image    string `gorm:"type:text;not null"`
	Quantity int    `gorm:"not null;default:0"`
}

type OrderStatus int

const (
	StatusReceived OrderStatus = iota + 1
	StatusProcessing
	StatusPicking
	StatusPacked
	StatusReadyShipment
	StatusInTransit
	StatusDelivered
	StatusCancelled
	StatusOnHold
	StatusBackordered
	StatusReturned
	StatusCompleted
)

type Order struct {
	Base
	RecipientFullName string      `gorm:"type:text;not null"`
	EmailAddress      string      `gorm:"type:text;not null"`
	DeliveryAddress   string      `gorm:"type:text;not null"`
	Phone             string      `gorm:"type:text;not null"`
	Status            OrderStatus `gorm:"not null;default:0"`

	OrderLines []OrderLine `gorm:"foreignKey:OrderID"`
}

type OrderLine struct {
	Base
	OrderID   string `gorm:"type:text;not null"`
	ProductID string `gorm:"type:text;not null"`
	Price     string `gorm:"type:text;not null"`
	Quantity  int    `gorm:"not null"`

	Product Product `gorm:"foreignKey:ProductID"`
}

type Invoice struct {
	Base
	OrderID     string        `gorm:"type:text;not null;unique"`
	TotalAmount string        `gorm:"type:text;not null"`
	PaidAmount  string        `gorm:"type:text;not null;default:0"`
	DueDate     time.Time     `gorm:"not null"`
	PaymentDate *time.Time    `gorm:"default:null"`
	Paid        bool          `gorm:"not null;default:false"`
	Items       []InvoiceItem `gorm:"foreignKey:InvoiceID"`

	Order Order `gorm:"foreignKey:OrderID"`
}

type InvoiceItem struct {
	Base
	InvoiceID string `gorm:"type:text;not null"`
	ProductID string `gorm:"type:text;not null"`
	Quantity  int    `gorm:"not null"`
	UnitPrice string `gorm:"type:text;not null"`
	TotalCost string `gorm:"type:text;not null"`

	Product Product `gorm:"foreignKey:ProductID"`
}
