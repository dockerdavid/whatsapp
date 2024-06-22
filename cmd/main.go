package main

import (
	"compi-whatsapp/internal/common"
	"compi-whatsapp/pkg/client"
	"compi-whatsapp/pkg/jobs"
	"compi-whatsapp/pkg/middlewares"
	"compi-whatsapp/pkg/routes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	common.InitEnv()
	client.InitClient()
	jobs.InitCron()

}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middlewares.CorsMiddleware())
	r.Use(middleware.Recoverer)

	routes.ErrorRoutes(r)
	routes.InitRoutes(r)

	fmt.Println("Server running on port", common.PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", common.PORT), r)
}
