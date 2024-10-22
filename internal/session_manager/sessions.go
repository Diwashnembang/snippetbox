package sessionmanager

import (
	"crypto/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Token       string 
	Value map[string]any
}
// type contextKey string
type SessionManager struct {
	IdleTimeout time.Duration
	Lifetime time.Duration
	Store Store
	Cookie SessionCookie
	Codec Codec
	ErrorFunc func(http.ResponseWriter, *http.Request, error)
	// contextKey contextKey
}

type SessionCookie struct{
	// Name sets the name of the session cookie. It should not contain
	// whitespace, commas, colons, semicolons, backslashes, the equals sign or
	// control characters as per RFC6265. The default cookie name is "session".
	// If your application uses two different sessions, you must make sure that
	// the cookie name for each is unique.
	Name string
	Domain string
	HttpOnly bool
	Path string
	Persist bool
	SameSite http.SameSite
	Secure bool
}


func (s *SessionManager) save(session Session){	

	

}
func randonTokenGenerator()uuid.UUID{
	token:=uuid.New()
	return token
}
func NewSession() *Session{
	return &Session{
		Token: randonTokenGenerator().String(),
		Value:make(map[string]any),
	}
}
