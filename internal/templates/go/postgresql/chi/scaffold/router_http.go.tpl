package router

import "github.com/go-chi/chi"

// RegisterUserRoutes registers the user-related routes in the provided chi.Router.
func RegisterUserRoutes(r chi.Router, uh *UserHandler) {
    // POST   /user
    r.Post("/user", uh.CreateUser)

    // GET    /user/{id}
    r.Get("/user/{id}", uh.GetUser)

    // PUT    /user/{id}
    r.Put("/user/{id}", uh.UpdateUser)

    // DELETE /user/{id}
    r.Delete("/user/{id}", uh.DeleteUser)
}