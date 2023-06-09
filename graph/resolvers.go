// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graph

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __getProductsInput is used internally by genqlient
type __getProductsInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// GetLimit returns __getProductsInput.Limit, and is useful for accessing the field via an interface.
func (v *__getProductsInput) GetLimit() int { return v.Limit }

// GetOffset returns __getProductsInput.Offset, and is useful for accessing the field via an interface.
func (v *__getProductsInput) GetOffset() int { return v.Offset }

// getBrandsBrandsBrand includes the requested fields of the GraphQL type Brand.
type getBrandsBrandsBrand struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// GetName returns getBrandsBrandsBrand.Name, and is useful for accessing the field via an interface.
func (v *getBrandsBrandsBrand) GetName() string { return v.Name }

// GetCode returns getBrandsBrandsBrand.Code, and is useful for accessing the field via an interface.
func (v *getBrandsBrandsBrand) GetCode() string { return v.Code }

// getBrandsResponse is returned by getBrands on success.
type getBrandsResponse struct {
	Brands []getBrandsBrandsBrand `json:"brands"`
}

// GetBrands returns getBrandsResponse.Brands, and is useful for accessing the field via an interface.
func (v *getBrandsResponse) GetBrands() []getBrandsBrandsBrand { return v.Brands }

// getProductsProductsProductPaginate includes the requested fields of the GraphQL type ProductPaginate.
type getProductsProductsProductPaginate struct {
	Pagination getProductsProductsProductPaginatePagination        `json:"pagination"`
	Products   []getProductsProductsProductPaginateProductsProduct `json:"products"`
}

// GetPagination returns getProductsProductsProductPaginate.Pagination, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginate) GetPagination() getProductsProductsProductPaginatePagination {
	return v.Pagination
}

// GetProducts returns getProductsProductsProductPaginate.Products, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginate) GetProducts() []getProductsProductsProductPaginateProductsProduct {
	return v.Products
}

// getProductsProductsProductPaginatePagination includes the requested fields of the GraphQL type Pagination.
type getProductsProductsProductPaginatePagination struct {
	Limit  int `json:"limit"`
	Totals int `json:"totals"`
	Offset int `json:"offset"`
}

// GetLimit returns getProductsProductsProductPaginatePagination.Limit, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginatePagination) GetLimit() int { return v.Limit }

// GetTotals returns getProductsProductsProductPaginatePagination.Totals, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginatePagination) GetTotals() int { return v.Totals }

// GetOffset returns getProductsProductsProductPaginatePagination.Offset, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginatePagination) GetOffset() int { return v.Offset }

// getProductsProductsProductPaginateProductsProduct includes the requested fields of the GraphQL type Product.
type getProductsProductsProductPaginateProductsProduct struct {
	Id                         int64                                                                                                  `json:"id"`
	Code                       string                                                                                                 `json:"code"`
	Brand                      getProductsProductsProductPaginateProductsProductBrand                                                 `json:"brand"`
	ProductSpecificationValues []getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue `json:"productSpecificationValues"`
}

// GetId returns getProductsProductsProductPaginateProductsProduct.Id, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProduct) GetId() int64 { return v.Id }

// GetCode returns getProductsProductsProductPaginateProductsProduct.Code, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProduct) GetCode() string { return v.Code }

// GetBrand returns getProductsProductsProductPaginateProductsProduct.Brand, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProduct) GetBrand() getProductsProductsProductPaginateProductsProductBrand {
	return v.Brand
}

// GetProductSpecificationValues returns getProductsProductsProductPaginateProductsProduct.ProductSpecificationValues, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProduct) GetProductSpecificationValues() []getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue {
	return v.ProductSpecificationValues
}

// getProductsProductsProductPaginateProductsProductBrand includes the requested fields of the GraphQL type Brand.
type getProductsProductsProductPaginateProductsProductBrand struct {
	Name string `json:"name"`
}

// GetName returns getProductsProductsProductPaginateProductsProductBrand.Name, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProductBrand) GetName() string { return v.Name }

// getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue includes the requested fields of the GraphQL type ProductSpecificationValue.
type getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue struct {
	Specification getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification `json:"specification"`
	Value         string                                                                                                                                `json:"value"`
}

// GetSpecification returns getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue.Specification, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue) GetSpecification() getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification {
	return v.Specification
}

// GetValue returns getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue.Value, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValue) GetValue() string {
	return v.Value
}

// getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification includes the requested fields of the GraphQL type ProductSpecification.
type getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification struct {
	Code string `json:"code"`
}

// GetCode returns getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification.Code, and is useful for accessing the field via an interface.
func (v *getProductsProductsProductPaginateProductsProductProductSpecificationValuesProductSpecificationValueSpecificationProductSpecification) GetCode() string {
	return v.Code
}

// getProductsResponse is returned by getProducts on success.
type getProductsResponse struct {
	Products getProductsProductsProductPaginate `json:"products"`
}

// GetProducts returns getProductsResponse.Products, and is useful for accessing the field via an interface.
func (v *getProductsResponse) GetProducts() getProductsProductsProductPaginate { return v.Products }

// The query or mutation executed by getBrands.
const getBrands_Operation = `
query getBrands {
	brands {
		name
		code
	}
}
`

func getBrands(
	ctx context.Context,
	client graphql.Client,
) (*getBrandsResponse, error) {
	req := &graphql.Request{
		OpName: "getBrands",
		Query:  getBrands_Operation,
	}
	var err error

	var data getBrandsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by getProducts.
const getProducts_Operation = `
query getProducts ($limit: Int!, $offset: Int!) {
	products(pagination: {limit:$limit,offset:$offset}) {
		pagination {
			limit
			totals
			offset
		}
		products {
			id
			code
			brand {
				name
			}
			productSpecificationValues {
				specification {
					code
				}
				value
			}
		}
	}
}
`

func getProducts(
	ctx context.Context,
	client graphql.Client,
	limit int,
	offset int,
) (*getProductsResponse, error) {
	req := &graphql.Request{
		OpName: "getProducts",
		Query:  getProducts_Operation,
		Variables: &__getProductsInput{
			Limit:  limit,
			Offset: offset,
		},
	}
	var err error

	var data getProductsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
