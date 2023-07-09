package main

import (
	"ecs_test_cars/framework"
	"ecs_test_cars/game"
	"ecs_test_cars/loader"
)

func main() {
	res := loader.ResourceLoader{}
	ttf := res.LoadFont("default")
	gg := game.NewGame(ttf)
	g := framework.InitWindowGame(gg, 800, 600, gg.GetTitle(), ttf)

	if err := g.Run(); err != nil {
		panic(err)
	}
}
