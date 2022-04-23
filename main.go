package main

import (
	"fmt"
	_ "log"
	"os"
	"projserver/pkg/app"
	"projserver/pkg/route"

	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func main() {
	fiberApp := app.CreateFiberApp()

	_ = godotenv.Load(".env")
	fiberApp.Use(cors.New())
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "",
	}))

	route.SetupRouter(fiberApp)

	// port
	port := os.Getenv("PORT")
	fmt.Printf("starting server at port %s", port)
	fiberApp.Listen(":" + port)
}
