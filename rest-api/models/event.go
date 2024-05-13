package models

import (
	"time"

	"elbolaky.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

// Save saves the event details to the database.
//
// No parameters.
// Returns an error.
func (event *Event) Save() error {
	query := `
		INSERT INTO events (
			name,
			description,
			location,
			dateTime,
			user_id
		) VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	event.ID = id
	return err
}

// GetAllEvents retrieves all events from the database.
//
// No parameters.
// Returns a slice of Event and an error.
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetEventByID retrieves an event by its ID from the database.
//
// Parameters:
// - id: The ID of the event to retrieve.
// Return type:
// - *Event: The event object.
// - error: An error if the retrieval fails.
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update updates the event details in the database.
//
// No parameters.
// Returns an error.
func (event *Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&event.Name, &event.Description, &event.Location, &event.DateTime, &event.ID)

	return err
}

// Delete deletes the event from the database.
//
// No parameters.
// Returns an error if the deletion fails.
func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&event.ID)

	return err
}

// Register registers a user for an event.
//
// Parameters:
// - userID: The ID of the user to register.
//
// Returns:
// - error: An error if the registration fails.
func (event *Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) values (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userID)

	return err
}

// CancelRegistration cancels a user's registration for an event.
//
// Parameters:
// - userID: The ID of the user to unregister.
//
// Returns:
// - error: An error if the cancellation fails.
func (event *Event) CancelRegistration(userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userID)

	return err
}
