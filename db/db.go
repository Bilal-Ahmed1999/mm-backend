package db

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	// Load .env if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	// Load Aiven CA certificate
	caCertPath := os.Getenv("DB_CA_CERT") // e.g., "./ca.pem"
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatal("Error reading CA certificate:", err)
	}

	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate")
	}

	// Register TLS config with MySQL driver
	tlsConfig := &tls.Config{
		RootCAs: rootCertPool,
	}
	err = mysql.RegisterTLSConfig("aiven", tlsConfig)
	if err != nil {
		log.Fatal("Error registering TLS config:", err)
	}

	// Build DSN with TLS
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=aiven&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("✅ Connected to MySQL database with TLS")
}
