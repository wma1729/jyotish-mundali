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
	sbLog := strings.Builder{}

	sb.WriteString("postgres://")
	sbLog.WriteString("postgres://")
	sb.WriteString(config.Database.User)
	sbLog.WriteString(config.Database.User)
	sb.WriteByte(':')
	sbLog.WriteByte(':')
	sb.WriteString(config.Database.Password)
	sbLog.WriteString("********")
	sb.WriteByte('@')
	sbLog.WriteByte('@')
	sb.WriteString(config.Database.Host)
	sbLog.WriteString(config.Database.Host)
	sb.WriteByte(':')
	sbLog.WriteByte(':')
	sb.WriteString(fmt.Sprintf("%d", config.Database.Port))
	sbLog.WriteByte(':')
	sb.WriteString("/jyotish?sslmode=disable")
	sbLog.WriteString("/jyotish?sslmode=disable")

	log.Printf("Opening postgreSQL db using connection string %s", sbLog.String())

	db, err := sql.Open("postgres", sb.String())
	if err != nil {
		log.Fatalf("unable to open database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("unable to connect to the database: %s", err)
	}

	return db
}
