package main

import (
	"context"
	"log"
	"path-api/internal/auth"
	"path-api/internal/db"
	"path-api/internal/kelas"
	"path-api/internal/pengguna"
	"path-api/internal/postingan"
	"path-api/internal/utils"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// @title PATH API
// @version 1.0
// @description Personal Adaptive Teaching Hub

// @contact.name MBKM PATH
// @contact.email mbkmpath@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := utils.ReadEnv()
	if err != nil {
		log.Printf("Failed to read .env: %v", err)
	}

	ctx := context.Background()
	q, pool, dbConn, err := db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	defer pool.Close()

	e := echo.New()

	e.Validator = utils.NewValidator()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	authHandler := auth.New(q)
	authHandler.Routes(e)

	penggunaHandler := pengguna.New(q)
	penggunaHandler.Routes(e)

	kelasHandler := kelas.New(q)
	kelasHandler.Routes(e)

	postinganHandler := postingan.New(q)
	postinganHandler.Routes(e)

	port := utils.GetPort()
	err = e.Start(port)
}
