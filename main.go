package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/taylankalkan01/golang-rate-limiter/limiter"
)

func main() {
	app := fiber.New()

	// Create a token bucket rate limiter with a capacity of 5 and rate of 2 request per second
	capacity := 5
	rate := 2
	limiter := limiter.NewTokenBucket(capacity, rate)

	app.Use(func(c *fiber.Ctx) error {
		if !limiter.TakeTokens(1) {
			errorMessage := fmt.Sprintf("Too many requests. Your max request number is %d. Please try again in %d seconds later", capacity, rate)
			return c.Status(fiber.StatusTooManyRequests).SendString(errorMessage)
		}
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Token bucket algorithm rate limiter in Golang.")
	})

	port := 3000
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}
}
