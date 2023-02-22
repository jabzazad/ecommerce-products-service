package product

import (
	"ecommerce-product/internal/models"
	"sync"
)

// MapImageToProduct map image to product
func MapImageToProduct(product <-chan *models.Product, images []*models.File, wg *sync.WaitGroup) {
	for p := range product {
		for _, productImageID := range p.CoverImageIDs {
			for _, image := range images {
				if uint(productImageID) == image.ID {
					p.CoverImages = append(p.CoverImages, image)
				}
			}
		}
	}
}
