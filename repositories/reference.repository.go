package repositories

import (
	"hris/models"

	"gorm.io/gorm"
)

type Reference interface {
	GetAllReference() ([]models.AD_Reference, error)
	GetReferenceByID(id uint) (models.AD_Reference, error)
	CreateReference(reference models.AD_Reference) (models.AD_Reference, error)
	UpdateReference(reference models.AD_Reference) error
	DeleteReference(id uint) error

	GetReferenceByName(name string) (models.AD_Reference, error)
}

type referenceRepository struct {
	db *gorm.DB
}

func NewReferenceRepository(db *gorm.DB) *referenceRepository {
	return &referenceRepository{db}
}

func (r *referenceRepository) GetAllReference() ([]models.AD_Reference, error) {
	var references []models.AD_Reference
	err := r.db.Find(&references).Error
	if err != nil {
		return nil, err
	}
	return references, nil
}

func (r *referenceRepository) GetReferenceByID(id uint) (models.AD_Reference, error) {
	var reference models.AD_Reference
	err := r.db.Find(&reference, id).Error
	if err != nil {
		return models.AD_Reference{}, err
	}
	return reference, nil
}

func (r *referenceRepository) CreateReference(reference models.AD_Reference) (models.AD_Reference, error) {
	err := r.db.Create(&reference).Error
	if err != nil {
		return models.AD_Reference{}, err
	}
	return reference, nil
}

func (r *referenceRepository) UpdateReference(reference models.AD_Reference) error {
	err := r.db.Save(&reference).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *referenceRepository) DeleteReference(id uint) error {
	err := r.db.Delete(&models.AD_Reference{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *referenceRepository) GetReferenceByName(name string) (models.AD_Reference, error) {
	var reference models.AD_Reference
	err := r.db.Where("ReferenceName = ?", name).First(&reference).Error
	if err != nil {
		return models.AD_Reference{}, err
	}
	return reference, nil
}
