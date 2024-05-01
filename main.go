package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/rest/server"
	util "github.com/manuelcunga/school-management-api/utils"
	"github.com/manuelcunga/school-management-api/utils/firebase"
)

func init() {
	err := firebase.InitFirebase()
	if err != nil {
		panic(err)
	}
	_ = godotenv.Load()

	util.Swagger()
}

func main() {
	server := server.NewServer(echo.New())
	server.Start()
}
