package main

import (
	"raynet/go-backend/configs"
	"raynet/go-backend/models"

)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}