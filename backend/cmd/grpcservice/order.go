package grpcservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	orderV1 "github.com/nikola-enter21/wms/backend/api/order/v1"
	"github.com/nikola-enter21/wms/backend/database/model"
	"github.com/nikola-enter21/wms/backend/integrations/sender"

	"github.com/nikola-enter21/wms/backend/convert"
)

func (s *Server) CreateOrder(ctx context.Context, in *orderV1.CreateOrderRequest) (*orderV1.CreateOrderResponse, error) {
	if len(in.Order.OrderLines) <= 0 {
		return nil, errors.New("cannot create an empty order")
	}

	orderModel, err := convert.CreateOrderRequestToModel(in)
	if err != nil {
		return nil, err
	}

	createdOrder, err := s.DB.CreateOrder(ctx, orderModel)
	if err != nil {
		return nil, err
	}

	protoOrder, err := convert.OrderModelToProto(createdOrder)
	if err != nil {
		return nil, err
	}

	return &orderV1.CreateOrderResponse{
		Order: protoOrder,
	}, nil
}

func (s *Server) GetOrder(ctx context.Context, in *orderV1.GetOrderRequest) (*orderV1.GetOrderResponse, error) {
	order, err := s.DB.GetOrder(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	protoOrder, err := convert.OrderModelToProto(order)
	if err != nil {
		return nil, err
	}

	return &orderV1.GetOrderResponse{
		Order: protoOrder,
	}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, in *orderV1.UpdateOrderRequest) (*orderV1.UpdateOrderResponse, error) {
	orderModel, err := convert.UpdateOrderRequestToModel(in)
	if err != nil {
		return nil, err
	}

	updatedOrder, err := s.DB.UpdateOrder(ctx, orderModel)
	if err != nil {
		return nil, err
	}

	go func() {
		err := s.orderHook(context.Background(), updatedOrder)
		log.Errorw("UpdateOrder", "Error", err)
	}()

	protoOrder, err := convert.OrderModelToProto(updatedOrder)
	if err != nil {
		return nil, err
	}

	return &orderV1.UpdateOrderResponse{
		Order: protoOrder,
	}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, in *orderV1.DeleteOrderRequest) (*orderV1.DeleteOrderResponse, error) {
	err := s.DB.DeleteOrder(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &orderV1.DeleteOrderResponse{}, nil
}

func (s *Server) ListOrders(ctx context.Context, in *orderV1.ListOrdersRequest) (*orderV1.ListOrdersResponse, error) {
	orders, err := s.DB.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	protoOrders, err := convert.OrdersModelToProto(orders)
	if err != nil {
		return nil, err
	}

	return &orderV1.ListOrdersResponse{
		Orders: protoOrders,
	}, nil
}

func (s *Server) orderHook(ctx context.Context, order *model.Order) error {
	// We might want to trigger some events on order_status change.
	// These are some examples:
	switch order.Status {
	case model.StatusCompleted:
		_, err := s.createInvoiceForCompletedOrder(ctx, order)
		return err
	case model.StatusCancelled:
		htmlBody := fmt.Sprintf("<p>Dear Customer,</p><p>Your order with ID %s has been cancelled. We apologize for any inconvenience caused.</p><p>Sincerely,<br>The WMS Team</p>", order.ID)
		textBody := fmt.Sprintf("Dear Customer,\n\nYour order with ID %s has been cancelled. We apologize for any inconvenience caused.\n\nSincerely,\nThe WMS Team", order.ID)
		return s.EmailSender.SendEmail(ctx, &sender.Email{
			RecipientAddress: order.EmailAddress,
			HTMLBody:         htmlBody,
			TextBody:         textBody,
		})
	case model.StatusDelivered:
		htmlBody := fmt.Sprintf("<p>Dear Customer,</p><p>Your order with ID %s has been successfully delivered. Thank you for choosing us!</p><p>Sincerely,<br>The WMS Team</p>", order.ID)
		textBody := fmt.Sprintf("Dear Customer,\n\nYour order with ID %s has been successfully delivered. Thank you for choosing us!\n\nSincerely,\nThe WMS Team", order.ID)
		return s.EmailSender.SendEmail(ctx, &sender.Email{
			RecipientAddress: order.EmailAddress,
			HTMLBody:         htmlBody,
			TextBody:         textBody,
		})
	default:
		// No action taken for the other statuses
	}
	return nil
}

func (s *Server) createInvoiceForCompletedOrder(ctx context.Context, order *model.Order) (*model.Invoice, error) {
	totalAmount, err := calculateTotalAmount(order.OrderLines)
	if err != nil {
		return nil, err
	}

	invoiceItems, err := orderLinesToInvoiceItems(order.OrderLines)
	if err != nil {
		return nil, err
	}

	invoice := &model.Invoice{
		OrderID:     order.ID,
		Items:       invoiceItems,
		TotalAmount: totalAmount,
		DueDate:     time.Now().AddDate(0, 0, 30), // Some due date
		PaidAmount:  totalAmount,
		Paid:        true,
	}

	createdInvoice, err := s.DB.CreateInvoice(ctx, invoice)
	if err != nil {
		return nil, err
	}

	return createdInvoice, nil
}

// orderLinesToInvoiceItems returns invoice items based on order lines
func orderLinesToInvoiceItems(orderLines []model.OrderLine) ([]model.InvoiceItem, error) {
	var invoiceItems []model.InvoiceItem
	for _, line := range orderLines {
		unitPrice, err := strconv.Atoi(line.Price)
		if err != nil {
			return nil, err
		}
		totalCost := line.Quantity * unitPrice
		invoiceItem := model.InvoiceItem{
			ProductID: line.ProductID,
			Quantity:  line.Quantity,
			UnitPrice: strconv.Itoa(unitPrice),
			TotalCost: strconv.Itoa(totalCost),
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}
	return invoiceItems, nil
}

func calculateTotalAmount(orderLines []model.OrderLine) (string, error) {
	total := 0
	for _, line := range orderLines {
		price, err := strconv.Atoi(line.Price)
		if err != nil {
			return "", err
		}
		total += line.Quantity * price
	}
	return strconv.Itoa(total), nil
}
