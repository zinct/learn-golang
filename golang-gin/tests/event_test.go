package tests

import (
	"golang-gin/db"
	"golang-gin/models"
	"testing"
	"time"
)

func init() {
	db.InitDB()
}

func createEvent() {
	event := models.Event{
		ID:          1,
		Name:        "Test Event",
		Description: "Test Description",
		Location:    "Test Location",
		Datetime:    time.Now(),
		UserId:      1,
	}
	event.Save()
}

func TestGetAllEvents(t *testing.T) {
	createEvent()
	events, err := models.GetAllEvents()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(events) == 0 {
		t.Errorf("Expected at least one event, got %d", len(events))
	}
}
