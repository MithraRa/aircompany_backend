package routes

import (
	"airline/controllers"
	"airline/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
)

func InitRoutes() *chi.Mux {
	root := chi.NewRouter()

	root.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	root.Use(middleware.Logger)

	root.Post("/signup", controllers.SignUp)
	root.Post("/signin", controllers.SignIn)
	root.Get("/flights", controllers.GetFlights)

	main := chi.NewRouter()

	root.Mount("/api", main)
	main.Use(middlewares.Auth)
	main.Get("/user/{userId}", controllers.GetUser)
	main.Get("/signout", controllers.SignOut)
	main.Put("/user/{userId}/info", controllers.UpdateUser)
	main.Delete("/user/{userId}/delete", controllers.DeleteUser)

	main.Get("/cart/{flightId}", controllers.GetFreeSeatsOnTicket)
	main.Post("/cart", controllers.CreateNewTicket)
	main.Get("/tickets/{userId}", controllers.GetBoughtTickets)

	log.Println("Created routes")
	return root
}
