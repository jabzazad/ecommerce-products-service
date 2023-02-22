// Package product enpoint
package product

import (
	"ecommerce-product/internal/core/config"
	"ecommerce-product/internal/handlers"
	"ecommerce-product/internal/request"

	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	FindAll(c *fiber.Ctx) error
	FindOne(c *fiber.Ctx) error
	BulkUpdateProducts(c *fiber.Ctx) error
	FindAllByIDs(c *fiber.Ctx) error
}

type endpoint struct {
	config  *config.Configs
	result  *config.ReturnResult
	service Service
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:  config.CF,
		result:  config.RR,
		service: NewService(),
	}
}

// FindAll find all
// @Tags Product
// @Summary FindAll
// @Description FindAll
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request query request.FindProductRequest true "query for get all"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /products [get]
func (ep *endpoint) FindAll(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.FindAll, &request.FindProductRequest{})
}

// FindOne find one
// @Tags Product
// @Summary FindOne
// @Description FindOne
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param id path uint true "ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /products/{id} [get]
func (ep *endpoint) FindOne(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.FindOne, &request.GetOne{})
}

// BulkUpdateProducts bulk update products
// @Tags Product
// @Summary BulkUpdateProducts
// @Description BulkUpdateProducts
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body request.BulkUpdateProducts true "request for bulk update products"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /products/bulk [post]
func (ep *endpoint) BulkUpdateProducts(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.BulkUpdateProducts, &request.BulkUpdateProducts{})
}

// FindAllByIDs find all by ids
// @Tags Product
// @Summary FindAllByIDs
// @Description FindAllByIDs
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request query request.FindProductIDsRequest true "request for find all by ids"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /products/ids [get]
func (ep *endpoint) FindAllByIDs(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.FindAllByIDs, &request.FindProductIDsRequest{})
}
