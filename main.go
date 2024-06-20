package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/myschool-petproj/class-ms/config"
	"github.com/myschool-petproj/class-ms/routes"
	"log"
	"net/http"
)

func main() {
	_ = godotenv.Load()
	conf := config.NewConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/class/api", routes.InitVersionedRoutes)

	log.Println("Starting server on :" + conf.AppPort)
	_ = http.ListenAndServe(":"+conf.AppPort, r)
}
