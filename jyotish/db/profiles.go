package db

import (
	"database/sql"
	"errors"
	"fmt"
	"jyotish/models"
	"log"
)

func ProfilesList(db *sql.DB, email string) ([]models.Profile, error) {
	query := `SELECT id, name, dob, city, state, country
		FROM profiles WHERE email = $1`

	log.Printf("sql - %s", query)

	rows, err := db.Query(query, email)
	if err != nil {
		log.Printf("unable to get profiles for %s: %s", email, err)
		return nil, err
	}

	defer rows.Close()

	profiles := make([]models.Profile, 0)
	index := 0

	for rows.Next() {
		var p models.Profile

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.DateOfBirth,
			&p.City,
			&p.State,
			&p.Country)
		if err != nil {
			log.Printf("failed to fetch row at index %d: %s", index, err)
			return nil, err
		}

		index++
		profiles = append(profiles, p)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("failed to fetch row at index %d: %s", index, err)
		return nil, err
	}

	return profiles, nil
}

func ProfileInsert(db *sql.DB, email string, profile *models.Profile) error {
	query := `INSERT INTO users (email, name, dob, city, state, country, details) VALUES
		($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, email, profile.Name, profile.DateOfBirth, profile.City,
		profile.State, profile.Country, profile.Details)

	var id string
	err := row.Scan(&id)
	if err != nil {
		log.Printf("unable to insert %s, %s: %s", email, profile.Name, err)
		return err
	} else {
		profile.ID = id
	}

	return nil
}

func ProfileGet(db *sql.DB, id string) (*models.Profile, error) {
	query := `SELECT name, dob, city, state, country FROM profiles
		WHERE id = $1`

	log.Printf("sql - %s", query)

	row := db.QueryRow(query, id)

	var p models.Profile
	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.DateOfBirth,
		&p.City,
		&p.State,
		&p.Country,
		&p.Details)
	if err != nil {
		log.Printf("unable to get %s: %s", id, err)
		return nil, err
	}

	return &p, nil
}

func ProfileUpdate(db *sql.DB, profile *models.Profile) error {
	query := `UPDATE profiles SET name = $2, dob = $3, city = $4,
		state = $5, country = $6, details = $7 WHERE id = $1`

	log.Printf("sql - %s", query)

	result, err := db.Exec(query,
		profile.ID,
		profile.Name,
		profile.DateOfBirth,
		profile.City,
		profile.State,
		profile.Country,
		profile.Details)
	if err != nil {
		log.Printf("unable to update %s: %s", profile.ID, err)
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Printf("unable to update %s: %s", profile.ID, err)
		return err
	}

	if rowsUpdated != 1 {
		err := errors.New(fmt.Sprintf("updated %d rows", rowsUpdated))
		log.Printf("unable to update %s: %s", profile.ID, err)
		return err
	}

	return nil
}

func ProfileDelete(db *sql.DB, id string) error {
	query := `DELETE FROM profiles WHERE id = $1`

	log.Printf("sql - %s", query)

	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("unable to update %s: %s", id, err)
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Printf("unable to update %s: %s", id, err)
		return err
	}

	if rowsDeleted != 1 {
		err := errors.New(fmt.Sprintf("deleted %d rows", rowsDeleted))
		log.Printf("unable to update %s: %s", id, err)
		return err
	}

	return nil
}
