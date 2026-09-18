package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"crudclientes/models"
)

type ClienteHandler struct {
	Repo      *models.ClienteRepository
	Templates *template.Template
}

func NewClienteHandler(repo *models.ClienteRepository, tmpl *template.Template) *ClienteHandler {
	return &ClienteHandler{Repo: repo, Templates: tmpl}
}

func (h *ClienteHandler) render(w http.ResponseWriter, name string, data any) {
	if err := h.Templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// List muestra todos los clientes.
func (h *ClienteHandler) List(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.render(w, "clientes_list.html", clientes)
}

// NewForm muestra el formulario para crear un cliente.
func (h *ClienteHandler) NewForm(w http.ResponseWriter, r *http.Request) {
	h.render(w, "cliente_form.html", struct {
		Cliente models.Cliente
		Accion  string
	}{Accion: "/clientes/nuevo"})
}

// Create procesa el formulario de creación.
func (h *ClienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := models.Cliente{
		Nombre:    r.FormValue("nombre"),
		Apellido:  r.FormValue("apellido"),
		Email:     r.FormValue("email"),
		Telefono:  r.FormValue("telefono"),
		Direccion: r.FormValue("direccion"),
	}

	if err := h.Repo.Create(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}

// EditForm muestra el formulario de edición con los datos actuales.
func (h *ClienteHandler) EditForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	c, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "cliente no encontrado", http.StatusNotFound)
		return
	}

	h.render(w, "cliente_form.html", struct {
		Cliente models.Cliente
		Accion  string
	}{Cliente: *c, Accion: "/clientes/" + r.PathValue("id") + "/editar"})
}

// Update procesa el formulario de edición.
func (h *ClienteHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := models.Cliente{
		ID:        id,
		Nombre:    r.FormValue("nombre"),
		Apellido:  r.FormValue("apellido"),
		Email:     r.FormValue("email"),
		Telefono:  r.FormValue("telefono"),
		Direccion: r.FormValue("direccion"),
	}

	if err := h.Repo.Update(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}

// Delete elimina un cliente.
func (h *ClienteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}
