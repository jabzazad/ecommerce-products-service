package repositories

import (
	"ecommerce-product/internal/models"

	"gorm.io/gorm"
)

// VariantRepository variant repositories
type VariantRepository interface {
	Create(db *gorm.DB, i interface{}) error
	Update(db *gorm.DB, i interface{}) error
	FindOneObjectByIDUInt(database *gorm.DB, id uint, i interface{}) error
	FindAllByUintIDs(db *gorm.DB, ids []uint, i interface{}) error
	BulkUpsert(db *gorm.DB, uniqueKey string, columns []string, i interface{}, batchSize int) error
	FindOneByIDFullAssociations(db *gorm.DB, id uint64, i interface{}) error
	FindVariantOptionByProductMasterID(db *gorm.DB, pmID uint) ([]*models.ReturnVariant, error)
}

type variantRepository struct {
	Repository
}

// VaraintNewRepository new sql repository
func VariantNewRepository() VariantRepository {
	return &variantRepository{
		NewRepository(),
	}
}

// FindVariantOptionByProductMasterID find variant option by product master id
func (repo *variantRepository) FindVariantOptionByProductMasterID(db *gorm.DB, pmID uint) ([]*models.ReturnVariant, error) {
	entities := []*models.ReturnVariant{}
	err := db.Select("option_name,variant_name").Table("variants").
		Joins("join products pv on pv.id = variants.product_id").Joins("join products on products.id = pv.product_id and products.id = ?", pmID).
		Group("option_name,variant_name").Order("min(variants.id)").Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}
