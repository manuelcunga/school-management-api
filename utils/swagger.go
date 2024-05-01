package util

import (
	"os"

	"github.com/manuelcunga/school-management-api/api"
)

func Swagger() {
	api.SwaggerInfo.Host = os.Getenv("REMOTEHOST")
	api.SwaggerInfo.BasePath = "/"
	api.SwaggerInfo.Schemes = []string{"http", "https"}
}
