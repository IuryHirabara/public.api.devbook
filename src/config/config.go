package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseStringConnection é a string de conexão com o MySQL
	DatabaseStringConnection = ""

	// Port é a porta onde a API vai estar rodando
	Port = 0

	// SecretKey é a chave que vai ser usada para assinar os tokens
	SecretKey []byte
)

// Inicializa as variaveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	dbPort, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	DatabaseStringConnection = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
	DatabaseStringConnection = fmt.Sprintf(DatabaseStringConnection,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
