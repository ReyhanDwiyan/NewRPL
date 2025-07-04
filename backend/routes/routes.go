package routes

import (
    "backend/controllers"
    "backend/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // TAMBAHKAN: Route untuk debugging di root
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("API Server is running")
    })

    // Auth routes - di kedua lokasi untuk kompatibilitas
    // 1. Tanpa prefix /auth untuk frontend lama
    app.Post("/login", controllers.Login)
    app.Post("/register", controllers.Register)
    
    // 2. Dengan prefix /auth untuk frontend baru
    auth := app.Group("/auth")
    auth.Post("/login", controllers.Login)
    auth.Post("/register", controllers.Register)

    // TAMBAHKAN: Non-protected User endpoint
    app.Get("/users", controllers.GetUsers)
    app.Get("/users/:id", controllers.GetUserById)

    // Protected Api routes
    api := app.Group("/api")
    api.Use(middleware.Protected())

    // Protected User endpoints dalam group /api
    api.Get("/users", controllers.GetUsers)
    api.Get("/team-members", controllers.GetTeamMembers)
    api.Get("/users/:id", controllers.GetUserById)
    api.Put("/users/:id", controllers.UpdateUser)

    // Add these lines to your routes setup

    // Emotion routes
    app.Post("/api/emotions", controllers.SaveEmotion)
    app.Get("/api/emotions/stats", controllers.GetEmotionStats)
    app.Get("/api/emotions/user/:id", controllers.GetUserEmotions)

    // Upload profile image
    api.Post("/upload-profile-image", controllers.UploadProfileImage)

    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.Status(200).JSON(fiber.Map{
            "status": "ok",
            "message": "Server is running",
        })
    })
}