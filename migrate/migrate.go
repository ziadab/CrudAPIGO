package main

import (
	"api/initializers"
	"api/models"
)

func init() {
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
