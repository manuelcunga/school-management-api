package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/rest/server"
	"github.com/manuelcunga/school-management-api/utils/firebase"
)

// @title School Management Api
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func init() {
	err := firebase.InitFirebase()
	if err != nil {
		panic(err)
	}
	_ = godotenv.Load()
}

func main() {
	server := server.NewServer(echo.New())
	server.Start()
}
