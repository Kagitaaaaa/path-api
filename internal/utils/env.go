package utils

import (
	"fmt"
	"log"
	"os"
)

func GetDatabaseURL() string {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatalln("DATABASE_URL belum diset!")
	}

	return dbUrl
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT tidak diset, default ke 3000")
		port = "3000"
	}

	return ":" + port
}

func GetSecret() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatalln("SECRET belum diset!")
	}

	return secret
}
