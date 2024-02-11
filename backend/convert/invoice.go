package convert

import (
	invoicev1 "github.com/nikola-enter21/wms/backend/api/invoice/v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nikola-enter21/wms/backend/database/model"
)

func InvoiceProtoToModel(proto *invoicev1.Invoice) (*model.Invoice, error) {
	invoiceModel := &model.Invoice{
		OrderID:     proto.OrderId,
		TotalAmount: proto.TotalAmount,
		PaidAmount:  proto.PaidAmount,
		DueDate:     proto.DueDate.AsTime(),
		Paid:        proto.Paid,
	}

	for _, item := range proto.Items {
		invoiceModel.Items = append(invoiceModel.Items, model.InvoiceItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
			UnitPrice: item.UnitPrice,
			TotalCost: item.TotalCost,
		})
	}

	return invoiceModel, nil
}

func InvoiceModelToProto(model *model.Invoice) (*invoicev1.Invoice, error) {
	items := make([]*invoicev1.InvoiceItem, len(model.Items))
	for i, item := range model.Items {
		items[i] = &invoicev1.InvoiceItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			UnitPrice: item.UnitPrice,
			TotalCost: item.TotalCost,
		}
	}

	var paymentDate *timestamppb.Timestamp
	if model.PaymentDate != nil {
		paymentDate = timestamppb.New(*model.PaymentDate)
	}

	return &invoicev1.Invoice{
		Id:          model.Base.ID,
		OrderId:     model.OrderID,
		TotalAmount: model.TotalAmount,
		PaidAmount:  model.PaidAmount,
		DueDate:     timestamppb.New(model.DueDate),
		PaymentDate: paymentDate,
		Paid:        model.Paid,
		Items:       items,
	}, nil
}

func InvoicesModelToProto(models []*model.Invoice) ([]*invoicev1.Invoice, error) {
	protos := make([]*invoicev1.Invoice, len(models))
	for i, m := range models {
		proto, err := InvoiceModelToProto(m)
		if err != nil {
			return nil, err
		}
		protos[i] = proto
	}
	return protos, nil
}

func UpdateInvoiceRequestToModel(proto *invoicev1.UpdateInvoiceRequest) (*model.Invoice, error) {
	return &model.Invoice{
		Base: model.Base{
			ID: proto.Invoice.Id,
		},
		OrderID:     proto.Invoice.OrderId,
		TotalAmount: proto.Invoice.TotalAmount,
		PaidAmount:  proto.Invoice.PaidAmount,
		DueDate:     proto.Invoice.DueDate.AsTime(),
		Paid:        proto.Invoice.Paid,
	}, nil
}
