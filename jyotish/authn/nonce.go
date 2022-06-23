package authn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
)

func SetNonceInCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	nonce, err := GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	SetCookie(w, r, "nonce", nonce)
	return nonce, nil
}

func ValidateNonce(r *http.Request, idToken *oidc.IDToken) error {
	nonce, err := r.Cookie("nonce")
	if err != nil {
		log.Printf("failed to get nonce from cookie: %s", err)
		return err
	}

	if nonce.Value != idToken.Nonce {
		return fmt.Errorf("nonce received (%s) in the cookie does not match the ID token nonce (%s)",
			nonce.Value, idToken.Nonce)
	}

	return nil
}
