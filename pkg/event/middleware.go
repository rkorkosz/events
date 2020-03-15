package event

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

// DefaultOptions
var DefaultOptions = &Options{
	VersionExtractor: PathVersionExtractor,
	TypeExtractor:    PathTypeExtractor,
	Logger:           log.New(os.Stdout, "[Events]", log.LstdFlags),
}

// ParamExtractor defines request param extractor
type ParamExtractor func(r *http.Request) string

// Options contains middleware options
type Options struct {
	VersionExtractor, TypeExtractor ParamExtractor
	Logger                          *log.Logger
}

// PathVersionExtractor extract version from path
func PathVersionExtractor(r *http.Request) string {
	return strings.Split(r.URL.Path, "/")[1]
}

// PathTypeExtractor extracts type from path
func PathTypeExtractor(r *http.Request) string {
	return strings.Split(r.URL.Path, "/")[2]
}

// Middleware stores incoming event in event store
func Middleware(store Store, opts *Options) func(next http.Handler) http.Handler {
	if opts == nil {
		opts = DefaultOptions
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawRequest, err := httputil.DumpRequest(r, true)
			if err != nil {
				opts.Logger.Printf("Error dumping request: %v\n", err)
			}
			event := Event{
				ID:              uuid.New(),
				Source:          r.URL.String(),
				SpecVersion:     opts.VersionExtractor(r),
				Type:            opts.TypeExtractor(r),
				DataContentType: r.Header.Get("Content-Type"),
				Time:            time.Now().UTC(),
				Data:            rawRequest,
			}
			err = store.Put(&event)
			if err != nil {
				opts.Logger.Printf("Error putting event to store: %v\n", err)
			}
			next.ServeHTTP(w, r)
		})
	}
}
