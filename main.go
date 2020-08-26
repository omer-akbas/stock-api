package main

import (
	"log"
	"net/http"
	"os"

	"github.com/omer-akbas/stock-api/config"
	"github.com/omer-akbas/stock-api/models"
	"github.com/omer-akbas/stock-api/routes"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

////// go run main.go dev|test|pro
func main() {
	cfg := config.Start(os.Args[1])
	models.Init()
	// mongoModels.Init()
	logFile, _ := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	r := routes.NewRouter()

	http.Handle("/", handlers.LoggingHandler(logFile, r))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8081", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})
	h := c.Handler(r)

	log.Println("Server starting on port :" + cfg.Server.Port)
	http.ListenAndServe(":"+cfg.Server.Port, h)
}
