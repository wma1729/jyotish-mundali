package db

import (
	"database/sql"
	"fmt"
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
	query := `INSERT INTO users (email, name) VALUES ($1, $2)
		RETURNING email, name, description, lang, astrologer, public`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email, name)

	var user models.User
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.Description,
		&user.Lang,
		&user.Astrologer,
		&user.Public)
	if err != nil {
		log.Printf("unable to insert %s, %s: %s", email, name, err)
		return nil, err
	}

	return &user, nil
}

func UserGet(db *sql.DB, email string) (*models.User, error) {
	query := `SELECT email, name, description, lang, astrologer, public
		FROM users WHERE email = $1`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email)

	var user models.User
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.Description,
		&user.Lang,
		&user.Astrologer,
		&user.Public)
	if err != nil {
		log.Printf("unable to get %s: %s", email, err)
		return nil, err
	}

	return &user, nil
}

func UserUpdate(db *sql.DB, user *models.User) error {
	query := `UPDATE users SET name = $2, description = $3, lang = $4,
		astrologer = $5, public = $6 WHERE email = $1`

	log.Printf("sql - %s", query)

	result, err := db.Exec(query,
		user.Email,
		user.Name,
		user.Description,
		user.Lang,
		user.Astrologer,
		user.Public)
	if err != nil {
		log.Printf("unable to update %s: %s", user.Email, err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Printf("unable to update %s: %s", user.Email, err)
		return err
	}

	if rowsUpdated != 1 {
		err := fmt.Errorf("updated %d rows", rowsUpdated)
		log.Printf("unable to update %s: %s", user.Email, err)
		return err
	}

	return nil
}
