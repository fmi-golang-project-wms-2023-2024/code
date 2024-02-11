package grpcservice

import (
	"context"

	invoicev1 "github.com/nikola-enter21/wms/backend/api/invoice/v1"
	"github.com/nikola-enter21/wms/backend/convert"
)

func (s *Server) CreateInvoice(ctx context.Context, in *invoicev1.CreateInvoiceRequest) (*invoicev1.CreateInvoiceResponse, error) {
	invoiceModel, err := convert.InvoiceProtoToModel(in.Invoice)
	if err != nil {
		return nil, err
	}

	createdInvoice, err := s.DB.CreateInvoice(ctx, invoiceModel)
	if err != nil {
		return nil, err
	}

	protoInvoice, err := convert.InvoiceModelToProto(createdInvoice)
	if err != nil {
		return nil, err
	}

	return &invoicev1.CreateInvoiceResponse{
		Invoice: protoInvoice,
	}, nil
}

func (s *Server) GetInvoice(ctx context.Context, in *invoicev1.GetInvoiceRequest) (*invoicev1.GetInvoiceResponse, error) {
	invoice, err := s.DB.GetInvoice(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	protoInvoice, err := convert.InvoiceModelToProto(invoice)
	if err != nil {
		return nil, err
	}

	return &invoicev1.GetInvoiceResponse{
		Invoice: protoInvoice,
	}, nil
}

func (s *Server) UpdateInvoice(ctx context.Context, in *invoicev1.UpdateInvoiceRequest) (*invoicev1.UpdateInvoiceResponse, error) {
	invoiceModel, err := convert.UpdateInvoiceRequestToModel(in)
	if err != nil {
		return nil, err
	}

	updatedInvoice, err := s.DB.UpdateInvoice(ctx, invoiceModel)
	if err != nil {
		return nil, err
	}

	protoInvoice, err := convert.InvoiceModelToProto(updatedInvoice)
	if err != nil {
		return nil, err
	}

	return &invoicev1.UpdateInvoiceResponse{
		Invoice: protoInvoice,
	}, nil
}

func (s *Server) DeleteInvoice(ctx context.Context, in *invoicev1.DeleteInvoiceRequest) (*invoicev1.DeleteInvoiceResponse, error) {
	err := s.DB.DeleteInvoice(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &invoicev1.DeleteInvoiceResponse{}, nil
}

func (s *Server) ListInvoices(ctx context.Context, in *invoicev1.ListInvoicesRequest) (*invoicev1.ListInvoicesResponse, error) {
	invoices, err := s.DB.ListInvoices(ctx)
	if err != nil {
		return nil, err
	}

	var protoInvoices []*invoicev1.Invoice
	for _, invoice := range invoices {
		protoInvoice, err := convert.InvoiceModelToProto(invoice)
		if err != nil {
			return nil, err
		}
		protoInvoices = append(protoInvoices, protoInvoice)
	}

	return &invoicev1.ListInvoicesResponse{
		Invoices: protoInvoices,
	}, nil
}
