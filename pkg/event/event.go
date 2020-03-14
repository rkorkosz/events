package event

import (
	"bufio"
	"bytes"
	"net/http"
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

// Replay sends event one more time
func (e *Event) Replay(client *http.Client) error {
	if client == nil {
		client = &http.Client{}
	}
	b := bufio.NewReader(bytes.NewReader(e.Data))
	req, err := http.ReadRequest(b)
	if err != nil {
		return err
	}
	_, err = client.Do(req)
	return err
}
