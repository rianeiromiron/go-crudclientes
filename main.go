package main

import (
	"html/template"
	"log"
	"net/http"

	"crudclientes/db"
	"crudclientes/handlers"
	"crudclientes/models"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("error conectando a la base de datos: %v", err)
	}
	defer conn.Close()

	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	repo := models.NewClienteRepository(conn)
	clienteHandler := handlers.NewClienteHandler(repo, tmpl)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("GET /clientes", clienteHandler.List)
	mux.HandleFunc("GET /clientes/nuevo", clienteHandler.NewForm)
	mux.HandleFunc("POST /clientes/nuevo", clienteHandler.Create)
	mux.HandleFunc("GET /clientes/{id}/editar", clienteHandler.EditForm)
	mux.HandleFunc("POST /clientes/{id}/editar", clienteHandler.Update)
	mux.HandleFunc("POST /clientes/{id}/eliminar", clienteHandler.Delete)

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/clientes", http.StatusSeeOther)
	})

	addr := ":8085"
	log.Printf("servidor escuchando en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
