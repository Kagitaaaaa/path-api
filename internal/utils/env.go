package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)

	}

	return scanner.Err()
}

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
