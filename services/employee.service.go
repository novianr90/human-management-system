package services

import (
	"hris/models"
	"hris/repositories"
)

type Employee interface {
	GetAllEmployee() ([]models.M_Employee, error)
	GetEmployeeByID(id uint) (models.M_Employee, error)
	CreateEmployee(employee models.M_Employee) (models.M_Employee, error)
	UpdateEmployee(employee models.M_Employee) error
	DeleteEmployee(id uint) error
	GetEmployeeByEmployeeID(employeeID string) (models.M_Employee, error)
}

type employeeService struct {
	employeeRepository repositories.Employee
}

func NewEmployeeService(employeeRepository repositories.Employee) *employeeService { {
	return &employeeService{employeeRepository}
}

func (s *employeeService) GetAllEmployee() ([]models.M_Employee, error) {
	return s.employeeRepository.GetAllEmployee()
}

func (s *employeeService) GetEmployeeByID(id uint) (models.M_Employee, error) {
	return s.employeeRepository.GetEmployeeByID(id)
}

func (s *employeeService) CreateEmployee(employee models.M_Employee) (models.M_Employee, error) {
	return s.employeeRepository.CreateEmployee(employee)
}

func (s *employeeService) UpdateEmployee(employee models.M_Employee) error {
	return s.employeeRepository.UpdateEmployee(employee)
}

func (s *employeeService) DeleteEmployee(id uint) error {
	return s.employeeRepository.DeleteEmployee(id)
}

func (s *employeeService) GetEmployeeByEmployeeID(employeeID string) (models.M_Employee, error) {
	return s.employeeRepository.GetEmployeeByEmployeeID(employeeID)
}
