package convert

import (
	productv1 "github.com/nikola-enter21/wms/backend/api/product/v1"
	"github.com/nikola-enter21/wms/backend/database/model"
)

func CreateProductRequestToModel(in *productv1.CreateProductRequest) (*model.Product, error) {
	return &model.Product{
		SKU:      in.Product.Sku,
		Title:    in.Product.Title,
		Price:    in.Product.Price,
		Image:    in.Product.Image,
		Quantity: int(in.Product.Quantity),
	}, nil
}

func ProductModelToProto(product *model.Product) (*productv1.Product, error) {
	return &productv1.Product{
		Id:       product.Base.ID,
		Sku:      product.SKU,
		Title:    product.Title,
		Price:    product.Price,
		Image:    product.Image,
		Quantity: int32(product.Quantity),
	}, nil
}

func ProductsModelToProto(dbProducts []*model.Product) ([]*productv1.Product, error) {
	products := []*productv1.Product{}

	for _, v := range dbProducts {
		protoProduct, err := ProductModelToProto(v)
		if err != nil {
			return nil, err
		}

		products = append(products, protoProduct)
	}

	return products, nil
}

func UpdateProductRequestToModel(in *productv1.UpdateProductRequest) (*model.Product, error) {
	productModel := &model.Product{
		Base: model.Base{
			ID: in.Product.Id,
		},
		SKU:      in.Product.Sku,
		Title:    in.Product.Title,
		Price:    in.Product.Price,
		Image:    in.Product.Image,
		Quantity: int(in.Product.Quantity),
	}

	return productModel, nil
}
