package main

import (
	"session-auth/model"
	"session-auth/routes"
)

func main() {
	model.Setup()
	routes.Setup()
}
