package services

import "hris/models"

type Reference interface {
	GetAllReference() ([]models.AD_Reference, error)
	GetReferenceByID(id uint) (models.AD_Reference, error)
	CreateReference(reference models.AD_Reference) (models.AD_Reference, error)
	UpdateReference(reference models.AD_Reference) error
	DeleteReference(id uint) error

	GetReferenceByName(name string) (models.AD_Reference, error)
}

type referenceService struct {
	referenceRepository Reference
}

func NewReferenceService(referenceRepository Reference) *referenceService {
	return &referenceService{referenceRepository}
}

func (s *referenceService) GetAllReference() ([]models.AD_Reference, error) {
	return s.referenceRepository.GetAllReference()
}

func (s *referenceService) GetReferenceByID(id uint) (models.AD_Reference, error) {
	return s.referenceRepository.GetReferenceByID(id)
}

func (s *referenceService) CreateReference(reference models.AD_Reference) (models.AD_Reference, error) {
	return s.referenceRepository.CreateReference(reference)
}

func (s *referenceService) UpdateReference(reference models.AD_Reference) error {
	return s.referenceRepository.UpdateReference(reference)
}

func (s *referenceService) DeleteReference(id uint) error {
	return s.referenceRepository.DeleteReference(id)
}

func (s *referenceService) GetReferenceByName(name string) (models.AD_Reference, error) {
	return s.referenceRepository.GetReferenceByName(name)
}
