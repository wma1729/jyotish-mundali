package db

import (
	"database/sql"
	"jyotish/models"
	"log"
	"strings"
)

func GetUserName(email, name, givenName, familyName string) string {
	if name != "" {
		return name
	}

	sb := strings.Builder{}

	if givenName != "" {
		sb.WriteString(givenName)
	}

	if familyName != "" {
		if sb.Len() > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(familyName)
	}

	if sb.Len() > 0 {
		return sb.String()
	}

	result := strings.Split(email, "@")
	return result[0]
}

func UserExists(db *sql.DB, email string) bool {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		log.Printf("unable to determine if %s exists: %s", email, err)
		return false
	}

	return exists
}

func UserInsert(db *sql.DB, email, name string) (*models.User, error) {
	query := `INSERT INTO users (Email, Name) VALUES ($1, $2)
		RETURNING Email, Name, Lang, Description, Astrologer, Public`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email, name)

	var user models.User
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.Lang,
		&user.Description,
		&user.Astrologer,
		&user.Public)
	if err != nil {
		log.Printf("unable to insert %s, %s: %s", email, name, err)
		return nil, err
	}

	return &user, nil
}

func UserGet(db *sql.DB, email string) (*models.User, error) {
	query := `SELECT Email, Name, Lang, Description, Astrologer, Public
		FROM users WHERE Email = $1`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email)

	var user models.User
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.Lang,
		&user.Description,
		&user.Astrologer,
		&user.Public)
	if err != nil {
		log.Printf("unable to get %s: %s", email, err)
		return nil, err
	}

	return &user, nil
}
