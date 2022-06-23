package authn

import (
	"fmt"
	"log"
	"net/http"
)

func SetStateInCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	state, err := GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	SetCookie(w, r, "state", state)
	return state, nil
}

func GetState(r *http.Request, fromCookie bool) (string, error) {
	if fromCookie {
		state, err := r.Cookie("state")
		if err != nil {
			log.Printf("failed to get state from cookie: %s", err)
			return "", err
		}
		return state.Value, nil
	}

	params := r.URL.Query()
	if params.Encode() == "" && r.Method == http.MethodPost {
		return r.FormValue("state"), nil
	}
	return params.Get("state"), nil
}

func ValidateState(r *http.Request) error {
	stateFromCookie, err := GetState(r, true)
	if err != nil {
		return err
	}

	stateFromRequest, err := GetState(r, false)
	if err != nil {
		return err
	}

	if stateFromRequest != stateFromCookie {
		return fmt.Errorf("state received (%s) in the request does not match the state in cookie (%s)",
			stateFromRequest, stateFromCookie)
	}

	return nil
}
