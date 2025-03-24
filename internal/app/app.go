package app

import (
	"log"

	"github.com/badprinter/task_manager/internal/config"
	"github.com/badprinter/task_manager/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	store *storage.Storage
	api   *fiber.App
}

func NewApp() (*App, error) {
	store, err := storage.NewStorageAndConnect()
	if err != nil {
		return nil, err
	}

	api := fiber.New()

	return &App{
		store,
		api,
	}, nil
}

func (a *App) Close() {
	a.store.Close()
}

func (a *App) regusterRouters() {
	taskGroup := a.api.Group("/tasks")

	taskGroup.Post("/", a.createTask)
	taskGroup.Get("/", a.getAll)
	taskGroup.Get("/:id", a.getTaskById)
	taskGroup.Delete("/:id", a.DeleteTaskById)
	taskGroup.Put("/", a.updateTask)
}

func (a *App) Start() {
	a.regusterRouters()
	log.Fatal(a.api.Listen(config.HTTP_HOST + ":" + config.HTTP_PORT))
}
