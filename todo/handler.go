package todo

import "net/http"

type Handler interface {
	// Home : GET /
	Home(w http.ResponseWriter, r *http.Request)

	// Search : GET /todos
	Search(w http.ResponseWriter, r *http.Request)

	// Create : POST /todos
	Create(w http.ResponseWriter, r *http.Request)

	// Update : PATCH /todos/{todoId}
	// Update : POST /todos/{todoId}/edit
	Update(w http.ResponseWriter, r *http.Request)

	// Get : GET /todos/{todoId}
	Get(w http.ResponseWriter, r *http.Request)

	// Delete : DELETE /todos/{todoId}
	// Delete : POST /todos/{todoId}/delete
	Delete(w http.ResponseWriter, r *http.Request)

	// Sort : POST /todos/sort
	Sort(w http.ResponseWriter, r *http.Request)
}
