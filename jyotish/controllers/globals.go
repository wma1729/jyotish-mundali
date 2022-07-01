package controllers

import (
	"database/sql"
	"jyotish/authn"
	"jyotish/config"
	"jyotish/db"
)

type Globals struct {
	Authenticator *authn.Authenticator
	DB            *sql.DB
}

func InitGlobals(config *config.Config) (*Globals, error) {
	var a authn.Authenticator
	a.Init(config)

	globals := &Globals{
		Authenticator: &a,
		DB:            db.ConnectToDB(config),
	}

	return globals, nil
}
