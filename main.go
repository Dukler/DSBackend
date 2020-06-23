package main

import (
	"DSBackend/controllers"
	"DSBackend/data"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	router := *chi.NewRouter()

	c := cors.New(cors.Options{
		//AllowedOrigins: []string{"https://duckstack.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},

		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With",  "Access-Control-Allow-Headers"},
		// AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(c.Handler)

	router.Get("/api/ui/update/{Screen}", controllers.UIEndpoint)
	router.Post("/api/do/{Action}", controllers.DoEndpoint)

	http.Handle("/", &router)

	port := os.Getenv("PORT")

	data.InitBucket("duckstackui")
	//REST
	log.Fatal(http.ListenAndServe(":"+port, &router))
}
