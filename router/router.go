package router

import (
	"os"

	"github.com/AlejandroWaiz/server/database"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	db database.Database
}

type RouterImplementation interface {
	ListenAndServe()
}

func New(db database.DatabaseImplementation) RouterImplementation {

	return &Router{}

}

func (r *Router) ListenAndServe() {

	router := fiber.New()

	router.Listen(os.Getenv("PORT"))

}

func (router *Router) prepareUrls(r *fiber.App) {

}
