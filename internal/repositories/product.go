package repositories

import (
	"ecommerce-product/internal/models"
	"ecommerce-product/internal/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// productRepository repo interface
type ProductRepository interface {
	Create(db *gorm.DB, i interface{}) error
	Update(db *gorm.DB, i interface{}) error
	FindOneObjectByIDUInt(db *gorm.DB, id uint, i interface{}) error
	FindOneByIDFullAssociations(db *gorm.DB, id uint64, i interface{}) error
	FindAll(db *gorm.DB, request *request.FindProductRequest) (*models.Page, error)
	BulkUpsert(db *gorm.DB, uniqueKey string, columns []string, i interface{}, batchSize int) error
	FindAllByProductIDs(db *gorm.DB, request *request.FindProductIDsRequest) ([]*models.Product, error)
	FindOneByID(db *gorm.DB, id uint) (*models.Product, error)
}

type productRepository struct {
	Repository
}

// ProductNewRepository new sql repository
func ProductNewRepository() ProductRepository {
	return &productRepository{
		NewRepository(),
	}
}

// FindAll find all
func (repo *productRepository) FindAll(db *gorm.DB, request *request.FindProductRequest) (*models.Page, error) {
	var entities []*models.Product
	page, err := repo.FindAllAndPageInformation(
		repo.query(db, request).Preload(clause.Associations).Preload("ProductVariants.Variants"), &request.PageForm, &entities,
	)
	if err != nil {
		return nil, err
	}

	return models.NewPage(page, entities), nil
}

func (repo *productRepository) query(db *gorm.DB, request *request.FindProductRequest) *gorm.DB {
	query := db.Model(&models.Product{})
	if request.IsPublish != nil {
		query = query.Where("is_publish = ?", request.IsPublish)
	}

	return query
}

// FindAllByProductIDs find all by product ids
func (repo *productRepository) FindAllByProductIDs(db *gorm.DB, request *request.FindProductIDsRequest) ([]*models.Product, error) {
	var entities []*models.Product
	err := db.Where("id IN ?", request.ProductIDs).Preload(clause.Associations).Preload("ProductVariants.Variants").
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// FindAll find all
func (repo *productRepository) FindOneByID(db *gorm.DB, id uint) (*models.Product, error) {
	var entity *models.Product
	err := db.Where("id = ?", id).Preload(clause.Associations).Preload("ProductVariants.Variants").
		First(&entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
