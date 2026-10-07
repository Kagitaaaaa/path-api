package main

import (
	"context"
	"log"
	"path-api/internal/db"
	"path-api/internal/kelas"
	"path-api/internal/pengguna"
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
	queries, pool, dbConn, err := db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	defer pool.Close()

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	penggunaHandler := pengguna.New(queries)
	penggunaHandler.Register(e)

	kelasHandler := kelas.New(queries)
	kelasHandler.Register(e)

	port := utils.GetPort()
	err = e.Start(port)
}
