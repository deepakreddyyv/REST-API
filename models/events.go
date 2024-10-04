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

func (e *Events) RegisterEvent(userID int64) error {
	insertQuery := "INSERT INTO REGISTRATIONS(user_id, event_id) values(?, ?)"

	_, err := db.DB.Exec(insertQuery, userID, e.Id)

	if err != nil {
		return err
	}

	return nil

}

func (e *Events) CancleRegistration(userId int64) (int64, error) {
	cancleRegisteration := "DELETE FROM REGISTRATIONS WHERE USER_ID = ? and EVENT_ID = ?"
    
	res, err := db.DB.Exec(cancleRegisteration, userId, e.Id);
    
	if err != nil {
		return -1, err 
	}
	rowsEff, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowsEff == 0 {
		return 0, errors.New("you havent registered for this event")
	}
	return rowsEff, nil

}
