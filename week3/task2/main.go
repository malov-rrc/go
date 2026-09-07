package main

import (
	"fmt"
	"task2/models"
)

func main() {
	user := models.User{ID: 1, Username: "trinity", IsActive: true}
	userJson := "{\"id\": 2, \"username\": \"neo\", \"is_active\": true}"
	fmt.Println(models.SerializeUser(user))
	fmt.Println(models.DeserializeUser(userJson).Username)
}
