package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func New() *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			GETOnly:      false,
			ErrorHandler: nil,
			AppName:      "api-gateway",
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}
