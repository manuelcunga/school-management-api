package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/manuelcunga/school-management-api/infra/rest/server"
	"github.com/manuelcunga/school-management-api/utils/firebase"
)

func init() {
	err := firebase.InitFirebase()
	if err != nil {
		// Lidar com o erro de inicialização do Firebase
		panic(err)
	}
	_ = godotenv.Load()
}

func main() {
	server := server.NewServer(echo.New())
	server.Start()
}
