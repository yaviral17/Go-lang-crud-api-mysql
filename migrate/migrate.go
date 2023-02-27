package main

import (
	initializers "github.com/aviral/go-crud/Initializers"
	"github.com/aviral/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
