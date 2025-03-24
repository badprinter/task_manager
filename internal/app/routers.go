package app

import (
	"strconv"

	"github.com/badprinter/task_manager/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (a *App) createTask(c *fiber.Ctx) error {
	var data *models.Task

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Если успешно
	if err := a.store.CreateTask(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.JSON(data)
}

func (a *App) getAll(c *fiber.Ctx) error {
	data, err := a.store.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.JSON(data)
}

func (a *App) getTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid task ID."})
	}
	data, err := a.store.GetTaskById(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(data)
}

func (a *App) DeleteTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Invalid task ID."})
	}
	err = a.store.DeleteTaskById(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(fiber.Map{
		"return": "true",
	})
}

func (a *App) updateTask(c *fiber.Ctx) error {

	var data *models.Task
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := a.store.UpdateTask(data); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(data)
}
