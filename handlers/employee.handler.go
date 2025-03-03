package handlers

import (
	"hris/models"
	"hris/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	service services.Employee
}

func NewEmployeeHandler(service services.Employee) *EmployeeHandler {
	return &EmployeeHandler{service}
}

// Get all employees
func (h *EmployeeHandler) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := h.service.GetAllEmployee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employees)
}

// Get employee by ID
func (h *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	employee, err := h.service.GetEmployeeByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}
	return c.JSON(employee)
}

// Create a new employee
func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var employee models.M_Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newEmployee, err := h.service.CreateEmployee(employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(newEmployee)
}

// Update an employee
func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	var employee models.M_Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdateEmployee(employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Employee updated successfully"})
}

// Delete an employee
func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.service.DeleteEmployee(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}

func (h *EmployeeHandler) GetEmployeeByEmployeeID(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	employee, err := h.service.GetEmployeeByEmployeeID(employeeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Employee not found"})
	}
	return c.JSON(employee)
}
