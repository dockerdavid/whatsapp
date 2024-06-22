package routes

import (
	"compi-whatsapp/pkg/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(mx *chi.Mux) {
	mx.Route("/api", func(r chi.Router) {
		r.Post("/send-message", services.HandleSendMessage)
		r.Post("/send-file", services.HandleSendFile)
	})
}

func ErrorRoutes(mx *chi.Mux) {
	mx.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	mx.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	})
}
