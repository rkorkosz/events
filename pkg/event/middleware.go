package event

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// EventMiddleware stores incoming event in event store
func EventMiddleware(store Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pathChunks := strings.Split(r.URL.Path, "/")
			event := Event{
				ID:              uuid.New(),
				Source:          r.URL.String(),
				SpecVersion:     pathChunks[0],
				Type:            pathChunks[1],
				DataContentType: r.Header.Get("Content-Type"),
				Time:            time.Now().UTC(),
			}
			if r.Body != nil {
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("[Event] - error reading request body: %w", err)
				} else {
					event.Data = json.RawMessage(body)
					r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
					err = store.Put(event)
					if err != nil {
						log.Fatalf("[Event] - error storing event: %v", err)
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
