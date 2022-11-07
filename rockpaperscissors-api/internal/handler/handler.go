package handler

import (
	"net/http"
	"net/url"
	"rockpaperscissors-api/internal/db"
	"unicode/utf8"
)

// Handler is a type that ties together the API method with the rest of the application.
// We use this as a way to access both the hub and the db in the API methods
type Handler struct {
	// hub *websocket.Hub
	db db.DB
}

// RunHub starts the hub
// func (h *Handler) RunHub() {
// 	h.hub.Run()
// }

// checkSameOrigin returns true if the origin is not set or is equal to the request host.
func checkOrigins(r *http.Request) bool {
	origin := r.Header["Origin"]
	if len(origin) == 0 {
		return true
	}
	u, err := url.Parse(origin[0])
	if err != nil {
		return false
	}
	for _, origin := range []string{
		//TODO: Change this for a better aproach like a config file
		r.Host,
		"localhost:3000",
	} {
		if equalASCIIFold(u.Host, origin) {
			return true
		}
	}

	return false
}

// equalASCIIFold returns true if s is equal to t with ASCII case folding as
// defined in RFC 4790.
func equalASCIIFold(s, t string) bool {
	for s != "" && t != "" {
		sr, size := utf8.DecodeRuneInString(s)
		s = s[size:]
		tr, size := utf8.DecodeRuneInString(t)
		t = t[size:]
		if sr == tr {
			continue
		}
		if 'A' <= sr && sr <= 'Z' {
			sr = sr + 'a' - 'A'
		}
		if 'A' <= tr && tr <= 'Z' {
			tr = tr + 'a' - 'A'
		}
		if sr != tr {
			return false
		}
	}
	return s == t
}

// NewHandler creates and returns a new handler
func NewHandler() Handler {
	return Handler{
		// hub: websocket.NewHub(),
		db: db.New(),
	}
}
