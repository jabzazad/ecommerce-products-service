package product

import (
	"ecommerce-product/internal/core/config"
	"ecommerce-product/internal/core/context"
	"ecommerce-product/internal/models"
	"ecommerce-product/internal/repositories"
	"ecommerce-product/internal/request"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	FindAll(c *context.Context, request *request.FindProductRequest) (*models.Page, error)
	FindOne(c *context.Context, request *request.GetOne) (*models.Product, error)
	FindAllByIDs(c *context.Context, request *request.FindProductIDsRequest) ([]*models.Product, error)
	BulkUpdateProducts(c *context.Context, request *request.BulkUpdateProducts) error
}

type service struct {
	config            *config.Configs
	result            *config.ReturnResult
	repository        repositories.ProductRepository
	fileRepository    repositories.FileRepository
	variantRepository repositories.VariantRepository
}

// NewService new service
func NewService() Service {
	return &service{
		config:            config.CF,
		result:            config.RR,
		repository:        repositories.ProductNewRepository(),
		fileRepository:    repositories.FileNewRepository(),
		variantRepository: repositories.VariantNewRepository(),
	}
}

// FindAll find all
func (s *service) FindAll(c *context.Context, request *request.FindProductRequest) (*models.Page, error) {
	db := c.GetDatabase()
	page, err := s.repository.FindAll(db, request)
	if err != nil {
		logrus.Errorf("find product error: %s", err)
		return nil, s.result.Internal.DatabaseNotFound
	}

	imageIDs := models.Int64Array{}
	for _, product := range page.Entities.([]*models.Product) {
		imageIDs = append(imageIDs, product.CoverImageIDs...)
	}

	images := []*models.File{}
	err = s.fileRepository.FindAllByIDs(db, imageIDs, &images)
	if err != nil {
		return nil, err
	}

	for _, p := range page.Entities.([]*models.Product) {
		for _, imageID := range p.CoverImageIDs {
			for _, image := range images {
				if imageID == int64(image.ID) {
					p.CoverImages = append(p.CoverImages, image)
				}
			}
		}
	}

	return page, nil
}

// FindAll find all
func (s *service) FindAllByIDs(c *context.Context, request *request.FindProductIDsRequest) ([]*models.Product, error) {
	db := c.GetDatabase()
	entities, err := s.repository.FindAllByProductIDs(db, request)
	if err != nil {
		logrus.Errorf("find product error: %s", err)
		return nil, s.result.Internal.DatabaseNotFound
	}

	imageIDs := models.Int64Array{}
	for _, product := range entities {
		imageIDs = append(imageIDs, product.CoverImageIDs...)
	}

	images := []*models.File{}
	err = s.fileRepository.FindAllByIDs(db, imageIDs, &images)
	if err != nil {
		return nil, err
	}

	for _, p := range entities {
		for _, imageID := range p.CoverImageIDs {
			for _, image := range images {
				if imageID == int64(image.ID) {
					p.CoverImages = append(p.CoverImages, image)
				}
			}
		}
	}

	return entities, nil
}

// FindOne find one by id
func (s *service) FindOne(c *context.Context, request *request.GetOne) (*models.Product, error) {
	db := c.GetDatabase()
	product, err := s.repository.FindOneByID(db, request.ID)
	if err != nil {
		logrus.Errorf("find product by id error: %s", err)
		return nil, s.result.Internal.DatabaseNotFound
	}

	imageIDs := product.CoverImageIDs
	if product.IsVariant {
		for _, productVaraint := range product.ProductVariants {
			imageIDs = append(imageIDs, productVaraint.CoverImageIDs...)
		}
	}

	files := []*models.File{}
	err = s.fileRepository.FindAllByIDs(db, imageIDs, &files)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		for _, imageID := range product.CoverImageIDs {
			if file.ID == uint(imageID) {
				product.CoverImages = append(product.CoverImages, file)
			}

		}
	}

	if product.IsVariant {
		entities, err := s.variantRepository.FindVariantOptionByProductMasterID(db, product.ID)
		if err != nil {
			logrus.Errorf("find variant option by product master id error: %s", err)
			return nil, s.result.Internal.DatabaseNotFound
		}

		product.VariantOptions = entities
		jobs := make(chan *models.Product, len(product.ProductVariants))
		wg := new(sync.WaitGroup)
		for w := 0; w < runtime.NumCPU()*2; w++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				MapImageToProduct(jobs, files, wg)
			}()
		}

		for _, productVariant := range product.ProductVariants {
			jobs <- productVariant
			if len(productVariant.CoverImages) == 0 {
				for _, imageID := range productVariant.CoverImageIDs {
					for _, image := range files {
						if imageID == int64(image.ID) {
							productVariant.CoverImages = append(productVariant.CoverImages, image)
						}
					}
				}
			}
		}

		close(jobs)
		wg.Wait()
	}

	return product, nil
}

// BulkUpdateProducts bulk update products
func (s *service) BulkUpdateProducts(c *context.Context, request *request.BulkUpdateProducts) error {
	db := c.GetDatabase()
	err := s.repository.BulkUpsert(db, "id", []string{}, request, 100)
	if err != nil {
		logrus.Errorf("bulk update product error: %s", err)
		return s.result.Internal.DatabaseNotFound
	}

	return nil
}
