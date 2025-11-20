package sessionInfra

import (
	"net/http"
	"time"
)

type CookieManager struct {
	name string
}

func NewCookieManager(name string) *CookieManager {
	return &CookieManager{name: name}
}

func (c *CookieManager) Set(w http.ResponseWriter, token string, expiresAt time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name: c.name,
		Value: token,
		Path: "/",
		Expires: expiresAt,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func (c *CookieManager) Get(r *http.Request) (string, error) {
	cookie, err := r.Cookie(c.name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (c *CookieManager) Clear(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: c.name,
		Value: "",
		Path: "/",
		MaxAge: -1,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteStrictMode,
	})
}