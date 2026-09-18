package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Connect abre la conexión a MySQL y crea la tabla clientes si no existe.
func Connect() (*sql.DB, error) {
	user := getEnv("DB_USER", "contabilidad")
	pass := getEnv("DB_PASSWORD", "contadocker")
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	name := getEnv("DB_NAME", "crudclientes")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, pass, host, port, name)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("no se pudo conectar a MySQL: %w", err)
	}

	if err := ensureSchema(conn); err != nil {
		return nil, err
	}

	return conn, nil
}

func ensureSchema(conn *sql.DB) error {
	const schema = `
	CREATE TABLE IF NOT EXISTS clientes (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		apellido VARCHAR(100) NOT NULL,
		email VARCHAR(150) NOT NULL,
		telefono VARCHAR(30),
		direccion VARCHAR(255),
		creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	_, err := conn.Exec(schema)
	return err
}
