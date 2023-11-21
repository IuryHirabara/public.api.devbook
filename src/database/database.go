package database

import (
	"database/sql"

	"api.devbook/src/config"
	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect abre a conexão com o banco de dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DatabaseStringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
