package main

import (
	"flag"
	"fmt"
	"service-account/api/middleware"
	"service-account/config"
	"service-account/pkg/container"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	app := fiber.New(fiber.Config{
		Prefork:               true,
		CaseSensitive:         true,
		StrictRouting:         true,
		ReadTimeout:           time.Second * 5,
		WriteTimeout:          time.Second * 10,
		IdleTimeout:           time.Second * 30,
		BodyLimit:             4 * 1024 * 1024,
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
	})
	middleware.AppMiddleware(app)
	container.DependencyInjection(app)
	host := flag.String("host", "0.0.0.0", "REST API host")
	port := flag.Int("port", 8080, "REST API port")
	flag.Parse()
	fmt.Printf("Server running on %s:%d\n", *host, *port)
	app.Listen(fmt.Sprintf("%s:%d", *host, *port))
}
