package router

import "github.com/gofiber/fiber/v2"

func (r *Router) getAllPokemons(ctx *fiber.Ctx) error {

	pokemons, err := r.db.GetAllPokemons()

	if err != nil {
		fiber.NewError(400, err.Error())
	}

	return ctx.JSON(pokemons)
}
