package main

import (
	"github.com/Babiel09/configs"
	"github.com/Babiel09/routes"
)

func main() {
	configs.Load()
	routes.HandlerRequests()
}
