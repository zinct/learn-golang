package models

import (
	"golang-gin/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserId      int
}

func (event Event) Save() error {
	_, err := db.DB.Exec(`
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`, event.Name, event.Description, event.Location, event.Datetime, event.UserId)

	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query(`
	SELECT id, name, description, location, datetime, user_id
	FROM events
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		var dateTimeStr string
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)

		if err != nil {
			return nil, err
		}
		event.Datetime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int) (*Event, error) {
	query := `
		SELECT id, name, description, location, datetime, user_id
		FROM events
		WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	var dateTimeStr string

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)
	if err != nil {
		return nil, err
	}

	event.Datetime, err = time.Parse("2006-01-02 15:04:05-07:00", dateTimeStr)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
