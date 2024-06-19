package router

import (
	"log"
	"os"

	"github.com/AlejandroWaiz/server/database"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	db database.DatabaseImplementation
}

type RouterImplementation interface {
	getAllPokemons(ctx *fiber.Ctx) error
	ListenAndServe()
}

func New(db database.DatabaseImplementation) RouterImplementation {

	return &Router{db: db}

}

func (r *Router) ListenAndServe() {

	router := fiber.New()

	r.prepareUrls(router)

	port := ":" + os.Getenv("PORT")

	log.Println(router.Listen(port))

}

func (router *Router) prepareUrls(r *fiber.App) {

	r.Get("/api/pokemons/", router.getAllPokemons)

}
