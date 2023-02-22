package request

import "ecommerce-product/internal/models"

type BulkUpdateProducts struct {
	Products []*models.Product `json:"products"`
}

type FindProductIDsRequest struct {
	ProductIDs []uint64 `query:"product_ids"`
}

type FindProductRequest struct {
	IsPublish *bool `query:"is_publish"`
	models.PageForm
}
