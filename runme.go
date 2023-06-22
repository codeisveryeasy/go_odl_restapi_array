package main

import (
	"fmt"
	"restapi_array/routes"
)

func main() {
	fmt.Println("Create API")
	routes.SetupRoutes()
}
