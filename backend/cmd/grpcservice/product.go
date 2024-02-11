package grpcservice

import (
	"context"

	productv1 "github.com/nikola-enter21/wms/backend/api/product/v1"
	"github.com/nikola-enter21/wms/backend/convert"
)

func (s *Server) CreateProduct(ctx context.Context, in *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	productModel, err := convert.CreateProductRequestToModel(in)
	if err != nil {
		return nil, err
	}

	createdProduct, err := s.DB.CreateProduct(ctx, productModel)
	if err != nil {
		return nil, err
	}

	protoProduct, err := convert.ProductModelToProto(createdProduct)
	if err != nil {
		return nil, err
	}

	return &productv1.CreateProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *Server) GetProduct(ctx context.Context, in *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	product, err := s.DB.GetProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	protoProduct, err := convert.ProductModelToProto(product)
	if err != nil {
		return nil, err
	}

	return &productv1.GetProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, in *productv1.UpdateProductRequest) (*productv1.UpdateProductResponse, error) {
	productModel, err := convert.UpdateProductRequestToModel(in)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := s.DB.UpdateProduct(ctx, productModel)
	if err != nil {
		return nil, err
	}

	protoProduct, err := convert.ProductModelToProto(updatedProduct)
	if err != nil {
		return nil, err
	}

	return &productv1.UpdateProductResponse{
		Product: protoProduct,
	}, nil
}

func (s *Server) ListProducts(ctx context.Context, in *productv1.ListProductsRequest) (*productv1.ListProductsResponse, error) {
	products, err := s.DB.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	protoProducts, err := convert.ProductsModelToProto(products)
	if err != nil {
		return nil, err
	}

	return &productv1.ListProductsResponse{
		Products: protoProducts,
	}, nil
}

func (s *Server) DeleteProduct(ctx context.Context, in *productv1.DeleteProductRequest) (*productv1.DeleteProductResponse, error) {
	err := s.DB.DeleteProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &productv1.DeleteProductResponse{}, nil
}
