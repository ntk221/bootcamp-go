package status

import (
	"net/http"

	"github.com/go-chi/chi"
	"yatter-backend-go/app/app"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	// r.Use(auth.Middleware(app))

	h := &handler{app: app}
	r.Post("/", h.Post)
	r.Get("/{id}", h.GetStatusByID)
	r.Delete("/{id}", h.DeleteStatusByID)

	return r
}
