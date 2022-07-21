package server

import (
	"errors"

	"github.com/apex/log"
	fiber "github.com/gofiber/fiber/v2"
)

// SetupFiber sets up a fiber app
func SetupFiber(s *Server, config *fiber.Config) *fiber.App {
	app := fiber.New(*config)

	// Error handling
	app.Use(s.Middleware())

	// Endpoint handlers
	// app.Get("/ping", s.HandlePing).Name("HandlePing")
	// app.Get("/directory/parents", s.HandleDirectoryParents).Name("HandleDirectoryParents")

	return app
}

// CreateFiberConfig creates a fiber config for the fiber app
func CreateFiberConfig() *fiber.Config {
	return &fiber.Config{
		AppName:        "vauthgo",
		Immutable:      true,
		ReadBufferSize: 8192, // 8kb, twice the default
	}
}

type httpError interface {
	StatusCode() int
	Message() string
	Unwrap() error
}

// Middleware returns a fiber.Handler that runs the handlers down the line,
// and catches any errors that implement httpStatus (e.g. httperrors.Error).
// It sets the status and returns the JSON of those errors in the following format:
// {
//   "error": string # result of Message()
// }
// Modified from go/httperrors
func (*Server) Middleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()

		var httpError httpError
		if !errors.As(err, &httpError) {
			// let fiber handle error since not formatted correctly
			if err != nil {
				log.Errorf("errorf processing request: %s", err)
			}
			return err
		}
		log.WithFields(
			log.Fields{
				"status":  httpError.StatusCode(),
				"message": httpError.Message(),
			}).Errorf("error processing request: %s", httpError.Unwrap())
		errObj := map[string]interface{}{
			"error": httpError.Message(),
		}

		return ctx.Status(httpError.StatusCode()).JSON(errObj)
	}
}
