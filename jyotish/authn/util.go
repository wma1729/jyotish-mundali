package authn

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"time"
)

func GenerateRandomString(n int) (string, error) {
	buf := make([]byte, n)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		log.Printf("failed to generate random string: %s", err)
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func SetCookie(w http.ResponseWriter, r *http.Request, key, value string) {
	http.SetCookie(w, &http.Cookie{
		Domain:   "localhost",
		Name:     key,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	})
}
