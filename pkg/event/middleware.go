package event

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Middleware stores incoming event in event store
func Middleware(store Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pathChunks := strings.Split(r.URL.Path, "/")
			rawRequest, err := httputil.DumpRequest(r, true)
			if err != nil {
				log.Println(err)
			}
			event := Event{
				ID:              uuid.New(),
				Source:          r.URL.String(),
				SpecVersion:     pathChunks[0],
				Type:            pathChunks[1],
				DataContentType: r.Header.Get("Content-Type"),
				Time:            time.Now().UTC(),
				Data:            rawRequest,
			}
			err = store.Put(event)
			if err != nil {
				log.Println(err)
			}
			next.ServeHTTP(w, r)
		})
	}
}
