package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/c7/todo.c7.se/templates/pages"
	"github.com/c7/todo.c7.se/templates/partials"
	"github.com/c7/todo.c7.se/todo"
)

type Handler struct {
	service todo.Service
}

func NewHandler(service todo.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func Mount(r chi.Router, h todo.Handler) {
	r.Get("/", h.Home)
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Post("/", h.Create)
		r.Route("/{todoId}", func(r chi.Router) {
			r.Patch("/", h.Update)
			r.Post("/edit", h.Update)
			r.Get("/", h.Get)
			r.Delete("/", h.Delete)
			r.Post("/delete", h.Delete)
		})
		r.Post("/sort", h.Sort)
	})
}

func (h Handler) Home(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := pages.HomePage(list).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Sort(w http.ResponseWriter, r *http.Request) {
	var todoIDs []uuid.UUID
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, id := range r.Form["id"] {
		var todoID uuid.UUID
		var err error
		if todoID, err = uuid.Parse(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todoIDs = append(todoIDs, todoID)
	}
	if err := h.service.Sort(r.Context(), todoIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h Handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")
	todos, err := h.service.Search(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodos(todos).Render(r.Context(), w)
	default:
		err = pages.TodosPage(todos, search).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var description = r.Form.Get("description")

	todo, err := h.service.Add(r.Context(), description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodo(todo).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var completed = r.Form.Get("completed") == "true"
	var description = r.Form.Get("description")

	todo, err := h.service.Update(r.Context(), todoID, completed, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodo(todo).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo, err := h.service.Get(r.Context(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.EditTodoForm(todo).Render(r.Context(), w)
	default:
		err = pages.TodoPage(todo).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Remove(r.Context(), todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		_, err = w.Write([]byte(""))
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isHTMX(r *http.Request) bool {
	// Check for "HX-Request" header
	if r.Header.Get("HX-Request") != "" {
		return true
	}

	return false
}
