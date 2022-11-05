package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
}

func (c *Session) InitSession() *scs.SessionManager {

	var persist, secure bool

	//lifetime of session?

	minutes, err := strconv.Atoi(c.CookieLifetime)

	if err != nil {
		minutes = 60
	}

	//persist?
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	// must cookies be secure?

	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	//which session store

	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":

	default:
		//cookie

	}
	return session
}
