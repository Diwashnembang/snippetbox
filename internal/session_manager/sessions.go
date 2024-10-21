package sessionmanager

import (
	"net/http"
	"time"
)

type Session struct {
	ID        string
	userID    string
	createdAt time.Time
	expiresAt time.Time
}
type contextKey string
type SessionManager struct {
	// IdleTimeout controls the maximum length of time a session can be inactive
	// before it expires. For example, some applications may wish to set this so
	// there is a timeout after 20 minutes of inactivity. By default IdleTimeout
	// is not set and there is no inactivity timeout.
	IdleTimeout time.Duration

	// Lifetime controls the maximum length of time that a session is valid for
	// before it expires. The lifetime is an 'absolute expiry' which is set when
	// the session is first created and does not change. The default value is 24
	// hours.
	Lifetime time.Duration

	// Store controls the session store where the session data is persisted.
	Store Store

	// Cookie contains the configuration settings for session cookies.
	Cookie SessionCookie

	// Codec controls the encoder/decoder used to transform session data to a
	// byte slice for use by the session store. By default session data is
	// encoded/decoded using encoding/gob.
	Codec Codec

	// ErrorFunc allows you to control behavior when an error is encountered by
	// the LoadAndSave middleware. The default behavior is for a HTTP 500
	// "Internal Server Error" message to be sent to the client and the error
	// logged using Go's standard logger. If a custom ErrorFunc is set, then
	// control will be passed to this instead. A typical use would be to provide
	// a function which logs the error and returns a customized HTML error page.
	ErrorFunc func(http.ResponseWriter, *http.Request, error)

	// // HashTokenInStore controls whether or not to store the session token or a hashed version in the store.
	// HashTokenInStore bool

	// contextKey is the key used to set and retrieve the session data from a
	// context.Context. It's automatically generated to ensure uniqueness.
	contextKey contextKey
}

type SessionCookie struct{
	// Name sets the name of the session cookie. It should not contain
	// whitespace, commas, colons, semicolons, backslashes, the equals sign or
	// control characters as per RFC6265. The default cookie name is "session".
	// If your application uses two different sessions, you must make sure that
	// the cookie name for each is unique.
	Name string

	// Domain sets the 'Domain' attribute on the session cookie. By default
	// it will be set to the domain name that the cookie was issued from.
	Domain string

	// HttpOnly sets the 'HttpOnly' attribute on the session cookie. The
	// default value is true.
	HttpOnly bool

	// Path sets the 'Path' attribute on the session cookie. The default value
	// is "/". Passing the empty string "" will result in it being set to the
	// path that the cookie was issued from.
	Path string

	// Persist sets whether the session cookie should be persistent or not
	// (i.e. whether it should be retained after a user closes their browser).
	// The default value is true, which means that the session cookie will not
	// be destroyed when the user closes their browser and the appropriate
	// 'Expires' and 'MaxAge' values will be added to the session cookie. If you
	// want to only persist some sessions (rather than all of them), then set this
	// to false and call the RememberMe() method for the specific sessions that you
	// want to persist.
	Persist bool

	// SameSite controls the value of the 'SameSite' attribute on the session
	// cookie. By default this is set to 'SameSite=Lax'. If you want no SameSite
	// attribute or value in the session cookie then you should set this to 0.
	SameSite http.SameSite

	// Secure sets the 'Secure' attribute on the session cookie. The default
	// value is false. It's recommended that you set this to true and serve all
	// requests over HTTPS in production environments.
	// See https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Session_Management_Cheat_Sheet.md#transport-layer-security.
	Secure bool
}


func (s *SessionManager) Load(){
	
}
