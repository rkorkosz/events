package event

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const EventContentType = "application/cloudevents+json"

type Event struct {
	ID              uuid.UUID       `json:"id"`
	Source          string          `json:"source"`
	SpecVersion     string          `json:"specversion"`
	Type            string          `json:"type"`
	DataContentType string          `json:"datacontenttype"`
	DataSchema      string          `json:"dataschema"`
	Subject         string          `json:"subject"`
	Time            time.Time       `json:"time"`
	Data            json.RawMessage `json:"data"`
}
