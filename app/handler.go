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
	r.Route("/items", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Post("/", h.Create)
		r.Route("/{ID}", func(r chi.Router) {
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
	var itemIDs []uuid.UUID

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, id := range r.Form["id"] {
		var itemID uuid.UUID

		var err error

		if itemID, err = uuid.Parse(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		itemIDs = append(itemIDs, itemID)
	}

	if err := h.service.Sort(r.Context(), itemIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h Handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")

	list, err := h.service.Search(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		err = partials.RenderList(list).Render(r.Context(), w)
	} else {
		err = pages.TodosPage(list, search).Render(r.Context(), w)
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

	item, err := h.service.Add(r.Context(), description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		err = partials.RenderItem(item).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "ID")

	var itemID uuid.UUID

	var err error

	if itemID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var completed = r.Form.Get("completed") == "true"
	var description = r.Form.Get("description")

	item, err := h.service.Update(r.Context(), itemID, completed, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		err = partials.RenderItem(item).Render(r.Context(), w)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "ID")

	var itemID uuid.UUID

	var err error

	if itemID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.Get(r.Context(), itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		err = partials.EditItemForm(item).Render(r.Context(), w)
	} else {
		err = pages.TodoPage(item).Render(r.Context(), w)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "ID")

	var itemID uuid.UUID

	var err error

	if itemID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Remove(r.Context(), itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isHTMX(r) {
		_, err = w.Write([]byte(""))
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isHTMX(r *http.Request) bool {
	// Check for "HX-Request" header
	return r.Header.Get("HX-Request") != ""
}
