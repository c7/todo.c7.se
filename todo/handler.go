package todo

import "net/http"

type Handler interface {
	// Home : GET /
	Home(w http.ResponseWriter, r *http.Request)

	// Search : GET /items
	Search(w http.ResponseWriter, r *http.Request)

	// Create : POST /items
	Create(w http.ResponseWriter, r *http.Request)

	// Update : PATCH /items/{ID}
	// Update : POST /items/{ID}/edit
	Update(w http.ResponseWriter, r *http.Request)

	// Get : GET /items/{ID}
	Get(w http.ResponseWriter, r *http.Request)

	// Delete : DELETE /items/{ID}
	// Delete : POST /items/{ID}/delete
	Delete(w http.ResponseWriter, r *http.Request)

	// Sort : POST /items/sort
	Sort(w http.ResponseWriter, r *http.Request)
}
