package main

import (
	"gotemplate/app/models"
	"gotemplate/routes"
)

func main() {
	models.GetInstance()
	routes.Run()
}
