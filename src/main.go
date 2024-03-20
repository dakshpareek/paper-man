package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daksh-pareek/paperman/src/config"
	_ "github.com/daksh-pareek/paperman/src/config"
	"github.com/daksh-pareek/paperman/src/repositories"
	"github.com/daksh-pareek/paperman/src/routes/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create server instance
	app := fiber.New()
	// Grouping routes
	apiV1 := app.Group("/api/v1")

	// Db initialization
	dbInstance := &repositories.DbQuizRepository{}

	// Register routes
	apiV1.Route("/quiz", api.QuizzesRoutes(dbInstance))

	// Handling graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM)
		signal.Notify(quit, syscall.SIGINT)

		<-quit
		fmt.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(ctx); err != nil {
			fmt.Printf("Server forced to shutdown: %v\n", err)
		}
	}()

	// Start server
	err := app.Listen(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
