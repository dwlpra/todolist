package main

import (
	"fmt"
	"runtime"
	"time"

	dbConfig "github.com/dwlpra/todolist/infrastructure/database/mysql"
	"github.com/dwlpra/todolist/infrastructure/persistance"
	"github.com/dwlpra/todolist/infrastructure/rest"
	"github.com/dwlpra/todolist/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	db := dbConfig.CreateConnection()
	// Inisialisasi repository
	activityRepo := persistance.NewActivity(db)
	todoRepo := persistance.NewTodo(db)

	activityService := usecase.NewActvityService(activityRepo)
	todoService := usecase.NewTodoService(todoRepo)

	activityHandler := rest.NewActivityHandler(activityService)
	todoHandler := rest.NewTodoHandler(todoService)

	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // LevelBestSpeed or LevelBestCompression
	}))
	cacheMiddleware := cache.New(cache.Config{
		Next: nil, // Lanjutkan dengan middleware berikutnya jika kondisi ini terpenuhi
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(fmt.Sprintf("%s?id=%s", c.Path(), c.Params("id")))
		},
		Expiration: 5 * time.Minute, // Durasi cache akan expire (default: 1 menit)
	})

	// app.Use(cacheMiddleware)

	app.Get("/activity-groups", cacheMiddleware, activityHandler.GetAllActivities)
	app.Get("/activity-groups/:id", cacheMiddleware, activityHandler.GetActivityByID)
	app.Post("/activity-groups", cacheMiddleware, activityHandler.CreateActivity)
	app.Patch("/activity-groups/:id", cacheMiddleware, activityHandler.UpdateActivity)
	app.Delete("/activity-groups/:id", cacheMiddleware, activityHandler.DeleteActivity)
	app.Get("/todo-items", todoHandler.GetAllTodos)
	app.Get("/todo-items/:id", todoHandler.GetTodoByID)
	app.Post("/todo-items", cacheMiddleware, todoHandler.CreateTodo)
	app.Patch("/todo-items/:id", cacheMiddleware, todoHandler.UpdateTodo)
	app.Delete("/todo-items/:id", cacheMiddleware, todoHandler.DeleteTodo)
	// Jalankan server
	app.Listen(":3030")
}
