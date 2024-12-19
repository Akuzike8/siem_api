package main

import (
	"log"
	"net/http"

	"github.com/Akuzike8/siem_api/handles"
	"github.com/go-chi/cors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/wazuh",func(r chi.Router){
		r.Put("/host_restart",handles.WazuhHostRestart)
	})
	
	r.Route("/velociraptor",func(r chi.Router){
		r.Post("/quaratine",handles.VelociraptorQuarantine)
		r.Post("/unquaratine",handles.VelociraptorUnQuarantine)
	})

	log.Println("server listening on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}