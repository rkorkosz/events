package event

import (
	"testing"

	"github.com/google/uuid"
)

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	event := &Event{
		ID:          uuid.New(),
		SpecVersion: "v1",
	}
	err := store.Put(event)
	if err != nil {
		t.Error(err)
	}
	storedEvent, err := store.Get(event.ID)
	if err != nil {
		t.Error(err)
	}
	if event != storedEvent {
		t.Errorf("Wrong event version: got %s, want %s", storedEvent.SpecVersion, event.SpecVersion)
	}
}
