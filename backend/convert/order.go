package convert

import (
	orderV1 "github.com/nikola-enter21/wms/backend/api/order/v1"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func CreateOrderRequestToModel(in *orderV1.CreateOrderRequest) (*model.Order, error) {
	orderModel := &model.Order{
		RecipientFullName: in.Order.RecipientFullName,
		EmailAddress:      in.Order.EmailAddress,
		DeliveryAddress:   in.Order.DeliveryAddress,
		Phone:             in.Order.Phone,
		Status:            model.OrderStatus(in.Order.Status),
	}

	for _, ol := range in.Order.OrderLines {
		orderModel.OrderLines = append(orderModel.OrderLines, model.OrderLine{
			ProductID: ol.ProductId,
			Price:     ol.Price,
			Quantity:  int(ol.Quantity),
		})
	}

	return orderModel, nil
}

func OrderModelToProto(order *model.Order) (*orderV1.Order, error) {
	orderLines := make([]*orderV1.OrderLine, len(order.OrderLines))
	for i, ol := range order.OrderLines {
		orderLines[i] = &orderV1.OrderLine{
			Id:        ol.ID,
			OrderId:   ol.OrderID,
			ProductId: ol.ProductID,
			Price:     ol.Price,
			Quantity:  int32(ol.Quantity),
		}
	}

	return &orderV1.Order{
		Id:                order.Base.ID,
		RecipientFullName: order.RecipientFullName,
		EmailAddress:      order.EmailAddress,
		DeliveryAddress:   order.DeliveryAddress,
		Phone:             order.Phone,
		Status:            orderV1.OrderStatus(order.Status),
		OrderLines:        orderLines,
	}, nil
}

func OrdersModelToProto(dbOrders []*model.Order) ([]*orderV1.Order, error) {
	orders := make([]*orderV1.Order, len(dbOrders))
	for i, o := range dbOrders {
		protoOrder, err := OrderModelToProto(o)
		if err != nil {
			return nil, err
		}
		orders[i] = protoOrder
	}
	return orders, nil
}

func UpdateOrderRequestToModel(in *orderV1.UpdateOrderRequest) (*model.Order, error) {
	orderModel := &model.Order{
		Base: model.Base{
			ID: in.Order.Id,
		},
		RecipientFullName: in.Order.RecipientFullName,
		EmailAddress:      in.Order.EmailAddress,
		DeliveryAddress:   in.Order.DeliveryAddress,
		Phone:             in.Order.Phone,
		Status:            model.OrderStatus(in.Order.Status),
	}

	for _, ol := range in.Order.OrderLines {
		orderModel.OrderLines = append(orderModel.OrderLines, model.OrderLine{
			ProductID: ol.ProductId,
			Price:     ol.Price,
			Quantity:  int(ol.Quantity),
		})
	}

	return orderModel, nil
}
