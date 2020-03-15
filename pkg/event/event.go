package event

import (
	"time"

	"github.com/google/uuid"
)

// Event represents application event
type Event struct {
	ID              uuid.UUID `json:"id"`
	Source          string    `json:"source"`
	SpecVersion     string    `json:"specversion"`
	Type            string    `json:"type"`
	DataContentType string    `json:"datacontenttype"`
	DataSchema      string    `json:"dataschema"`
	Subject         string    `json:"subject"`
	Time            time.Time `json:"time"`
	Data            []byte    `json:"data"`
}
