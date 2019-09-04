package main

import (
	"DSBackend/controllers"
	"DSBackend/data"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)




func main() {
	router := *chi.NewRouter()

	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With",  "Access-Control-Allow-Headers"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(c.Handler)

	router.Get("/api/ui/update/{Screen}", controllers.UIEndpoint)
	router.Post("/api/do/{Action}", controllers.DoEndpoint)

	http.Handle("/", &router)

	//handler := handlers.CORS(
	//	handlers.AllowedOrigins([]string{"*"}),
	//	handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
	//	handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	//)(router)

	port :=  os.Getenv("PORT")

	data.InitBucket("duckstackui")
	//REST
	log.Fatal(http.ListenAndServe(":"+ port, &router))
}

