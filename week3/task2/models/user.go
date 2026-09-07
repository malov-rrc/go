package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

func SerializeUser(user User) string {
	serializedUser, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка сериализации: %v\n", err)
		os.Exit(1)
	}
	return string(serializedUser)
}

func DeserializeUser(userJson string) User {
	var user = User{}
	err := json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка при десериализации: %v\n", err)
		os.Exit(1)
	}
	return user
}
