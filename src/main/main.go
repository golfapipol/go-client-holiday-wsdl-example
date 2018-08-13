package main

import (
	"holiday/api"
)

func main() {
	route := api.SetupRoute()
	route.Start(":3000")
}
