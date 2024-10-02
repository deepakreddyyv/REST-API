package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"deepak.com/web_rest/db"
)

type Events struct {
	Id          int64
	Name        string    `bindings:"required"`
	Description string    `bindings:"required"`
	Location    string    `bindings:"required"`
	EventDate   time.Time `bindings:"required"`
	UserId      int64
}

func parseTime(dateByte []byte) (time.Time, error) {
	str := string(dateByte)
	return time.Parse("2006-01-02 15:04:05", str)
}

func GetEvents(p ...any) ([]Events, error) {
	var events = []Events{}

	selectQuery := "select * from events"

	if l := len(p); l > 0 {
		selectQuery = `
	        select * from events where id in (?)
	    `
	}

	rows, err := db.DB.Query(selectQuery, p...)

	if err != nil {
		return []Events{}, err
	}

	for rows.Next() {
		var event Events
		var eventDateBytes []byte

		if err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &eventDateBytes, &event.UserId); err != nil {
			fmt.Println(err)
			return []Events{}, err
		}

		eventDate, err := parseTime(eventDateBytes)
		if err != nil {
			eventDate = time.Time{}
		}
		event.EventDate = eventDate

		events = append(events, event)
	} // rows.Close() will close automatically when rown.Next() returns False

	return events, nil
}

func (e *Events) Save() error {

	insertQuery := `
	    insert into events(name, description, location, event_date, user_id)
		values(?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.EventDate, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger := log.Default()
		logger.Fatal("Warning: Setting id to default value 0")
		id = 0
	}

	e.Id = id
	return nil
}

func (e *Events) UpdateEvents(id int64) error {
	updateQuery := `
	UPDATE EVENTS SET 
    NAME = ?,
	DESCRIPTION = ?,
	LOCATION = ?,
	EVENT_DATE = ?,
    USER_ID = ?
	WHERE ID = ?
	`

	_, err := db.DB.Exec(updateQuery, e.Name, e.Description, e.Location, e.EventDate, e.UserId, id)

	return err
}

func DeleteEvents(id int64) error {
	deleteQuery := `
	    DELETE FROM EVENTS WHERE ID = ?
	`

	res, err := db.DB.Exec(deleteQuery, id)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()

	if rows == 0 {
		return errors.New("no event with that id")
	}

	return err
}
