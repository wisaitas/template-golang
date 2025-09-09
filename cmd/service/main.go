package main

import "github.com/wisaitas/template-golang/internal/service/initial"

func main() {
	app := initial.New()

	app.Run()

	app.Close()
}
