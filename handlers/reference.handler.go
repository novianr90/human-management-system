package handlers

import (
	"hris/services"
	"strconv"
)

type ReferenceHandler struct {
	service services.Reference
}

func NewReferenceHandler(service services.Reference) *ReferenceHandler {
	return &ReferenceHandler{service}
}

func (h *ReferenceHandler) GetAllReferences(c *fiber.Ctx) error {
	references, err := h.service.GetAllReference()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(references)
}

func (h *ReferenceHandler) GetReferenceByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	reference, err := h.service.GetReferenceByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reference not found"})
	}
	return c.JSON(reference)
}

func (h *ReferenceHandler) CreateReference(c *fiber.Ctx) error {
	var reference models.M_Reference
	if err := c.BodyParser(&reference); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newReference, err := h.service.CreateReference(reference)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(newReference)
}

func (h *ReferenceHandler) UpdateReference(c *fiber.Ctx) error {
	var reference models.M_Reference
	if err := c.BodyParser(&reference); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.service.UpdateReference(reference); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reference updated successfully"})
}

func (h *ReferenceHandler) DeleteReference(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.service.DeleteReference(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Reference deleted successfully"})
}

func (h *ReferenceHandler) GetReferenceByName(c *fiber.Ctx) error {
	name := c.Params("name")
	reference, err := h.service.GetReferenceByName(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reference not found"})
	}
	return c.JSON(reference)
}
