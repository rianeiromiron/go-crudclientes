package models

import (
	"database/sql"
	"time"
)

type Cliente struct {
	ID        int
	Nombre    string
	Apellido  string
	Email     string
	Telefono  string
	Direccion string
	CreadoEn  time.Time
}

type ClienteRepository struct {
	DB *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepository {
	return &ClienteRepository{DB: db}
}

func (r *ClienteRepository) GetAll() ([]Cliente, error) {
	rows, err := r.DB.Query(`SELECT id, nombre, apellido, email, telefono, direccion, creado_en FROM clientes ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Apellido, &c.Email, &c.Telefono, &c.Direccion, &c.CreadoEn); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}
	return clientes, rows.Err()
}

func (r *ClienteRepository) GetByID(id int) (*Cliente, error) {
	var c Cliente
	err := r.DB.QueryRow(`SELECT id, nombre, apellido, email, telefono, direccion, creado_en FROM clientes WHERE id = ?`, id).
		Scan(&c.ID, &c.Nombre, &c.Apellido, &c.Email, &c.Telefono, &c.Direccion, &c.CreadoEn)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ClienteRepository) Create(c Cliente) error {
	_, err := r.DB.Exec(
		`INSERT INTO clientes (nombre, apellido, email, telefono, direccion) VALUES (?, ?, ?, ?, ?)`,
		c.Nombre, c.Apellido, c.Email, c.Telefono, c.Direccion,
	)
	return err
}

func (r *ClienteRepository) Update(c Cliente) error {
	_, err := r.DB.Exec(
		`UPDATE clientes SET nombre = ?, apellido = ?, email = ?, telefono = ?, direccion = ? WHERE id = ?`,
		c.Nombre, c.Apellido, c.Email, c.Telefono, c.Direccion, c.ID,
	)
	return err
}

func (r *ClienteRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM clientes WHERE id = ?`, id)
	return err
}
