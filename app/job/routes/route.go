package routes

import (
	"syncerland/app/job/controller"
	"syncerland/core/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(router *fiber.App) {
	userGroup := router.Group("/job")

	userGroup.Post("/", middlewares.AuthMiddleware, controller.CreateJobHandler)
	userGroup.Patch("/", middlewares.AuthMiddleware, controller.UpdateJobHandler)
	userGroup.Delete("/", middlewares.AuthMiddleware, controller.DeleteJobHandler)
	userGroup.Get("/", middlewares.AuthMiddleware, controller.FindAllJobsHandler)
	userGroup.Get("/:id", middlewares.AuthMiddleware, controller.FindOneHandler)
}
