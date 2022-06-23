package authn

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type User struct {
	Token          *oauth2.Token
	IDToken        *oidc.IDToken
	RawIDToken     string
	User           *oidc.UserInfo
	Name           string `json:"name,omitempty"`
	GivenName      string `json:"given_name,omitempty"`
	FamilyName     string `json:"family_name,omitempty"`
	LastAccessTime time.Time
	ExpiresOn      time.Time
	SessionId      string
}

type UserStore struct {
	mutex   sync.Mutex
	userMap map[string]*User
}

var store UserStore

func init() {
	store.userMap = make(map[string]*User)
}

func AddUserToStore(u *User) (string, error) {
	sessionId, err := GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	u.LastAccessTime = time.Now()
	u.ExpiresOn = u.IDToken.Expiry
	u.SessionId = sessionId

	store.userMap[sessionId] = u

	return sessionId, nil
}

func RemoveUserFromStore(sessionId string) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	delete(store.userMap, sessionId)
}

func GetUserFromStore(sessionId string) (user *User) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	var ok bool
	user, ok = store.userMap[sessionId]
	if !ok {
		log.Printf("user with session ID (%s) not found", sessionId)
		return nil
	}

	user.LastAccessTime = time.Now()
	return user
}

func GetUserSession(r *http.Request) (*User, error) {
	sessionId, err := r.Cookie("sid")
	if err != nil {
		log.Printf("failed to get session ID from cookie: %s", err)
		return nil, err
	}

	if sessionId.Value == "" {
		return nil, fmt.Errorf("no user session found")
	}

	user := GetUserFromStore(sessionId.Value)
	if user == nil {
		return nil, fmt.Errorf("user session ID (%s) not found", sessionId.Value)
	}

	now := time.Now()
	if now.After(user.ExpiresOn) {
		RemoveUserFromStore(sessionId.Value)
		return nil, fmt.Errorf("user session has expired")
	}

	return user, nil
}

func SetUserSession(w http.ResponseWriter, r *http.Request, u *User) error {
	sessionId, err := AddUserToStore(u)
	if err != nil {
		log.Printf("failed to add user to store")
		return err
	}

	cookie := &http.Cookie{
		Domain:   "localhost",
		Path:     "/",
		Name:     "sid",
		Value:    sessionId,
		MaxAge:   12 * 60 * 60,
		Secure:   r.TLS != nil,
		HttpOnly: false,
	}

	http.SetCookie(w, cookie)

	return nil
}

func ResetUserSession(w http.ResponseWriter, r *http.Request, u *User) error {
	RemoveUserFromStore(u.SessionId)

	cookie := &http.Cookie{
		Domain:   "localhost",
		Name:     "sid",
		Value:    "",
		MaxAge:   -1,
		Secure:   r.TLS != nil,
		HttpOnly: false,
	}

	http.SetCookie(w, cookie)
	return nil
}
