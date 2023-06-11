package main

import (
	"yhkim/gowebserver/routes"
	"yhkim/gowebserver/services"
)

func main() {
	services.RegisterServices()
	routes.Run()
}
