package repositories

import (
	"hris/models"

	"gorm.io/gorm"
)

type Employee interface {
	GetAllEmployee() ([]models.M_Employee, error)
	GetEmployeeByID(id uint) (models.M_Employee, error)
	CreateEmployee(employee models.M_Employee) (models.M_Employee, error)
	UpdateEmployee(employee models.M_Employee) error
	DeleteEmployee(id uint) error
	GetEmployeeByEmployeeID(employeeID string) (models.M_Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAllEmployee() ([]models.M_Employee, error) {
	var employees []models.M_Employee
	err := r.db.
		Preload("GenderRef").
		Preload("NationalityRef").
		Preload("BloodTypeRef").
		Preload("DepartmentRef").
		Preload("MaritalStatusRef").
		Preload("ReligionRef").
		Find(&employees).
		Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) GetEmployeeByID(id uint) (models.M_Employee, error) {
	var employee models.M_Employee
	err := r.db.
		Preload("GenderRef").
		Preload("NationalityRef").
		Preload("BloodTypeRef").
		Preload("DepartmentRef").
		Preload("MaritalStatusRef").
		Preload("ReligionRef").
		Find(&employee, id).Error
	if err != nil {
		return models.M_Employee{}, err
	}
	return employee, nil
}

func (r *employeeRepository) CreateEmployee(employee models.M_Employee) (models.M_Employee, error) {
	err := r.db.Create(&employee).Error
	if err != nil {
		return models.M_Employee{}, err
	}
	return employee, nil
}

func (r *employeeRepository) UpdateEmployee(employee models.M_Employee) error {
	err := r.db.Save(&employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *employeeRepository) DeleteEmployee(id uint) error {
	err := r.db.Delete(&models.M_Employee{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *employeeRepository) GetEmployeeByEmployeeID(employeeID string) (models.M_Employee, error) {
	var employee models.M_Employee
	err := r.db.
		Preload("GenderRef").
		Preload("NationalityRef").
		Preload("BloodTypeRef").
		Preload("DepartmentRef").
		Preload("MaritalStatusRef").
		Preload("ReligionRef").
		Where("EmployeeID = ?", employeeID).
		Find(&employee).Error
	if err != nil {
		return models.M_Employee{}, err
	}
	return employee, nil
}
