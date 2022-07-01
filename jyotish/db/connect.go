package db

import (
	"database/sql"
	"fmt"
	"jyotish/config"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func ConnectToDB(config *config.Config) *sql.DB {
	sb := strings.Builder{}

	sb.WriteString("postgres://")
	sb.WriteString(config.Database.User)
	sb.WriteByte(':')
	sb.WriteString(config.Database.Password)
	sb.WriteByte('@')
	sb.WriteString(config.Database.Host)
	sb.WriteByte(':')
	sb.WriteString(fmt.Sprintf("%d", config.Database.Port))
	sb.WriteString("/jyotish")

	log.Printf("Opening postgreSQL db using connection string %s", sb.String())

	db, err := sql.Open("postgres", sb.String())
	if err != nil {
		log.Fatalf("unable to open database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("unable to connect to the database: %s", err)
	}

	return db
}
