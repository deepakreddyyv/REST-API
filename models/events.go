package models

import (
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

func GetEventById(p int64) (Events, error) {

	var e Events
	tx := db.DB.Find(&e, p)

	return e, tx.Error

}

func GetEvents() ([]Events, error) {
	var events = []Events{}

	tx := db.DB.Find(&events)

	return events, tx.Error
}

func (e *Events) Save() error {

	tx := db.DB.Create(&e)

	return tx.Error

}

func (e *Events) UpdateEvents(id int64) error {

	event, err := GetEventById(id)
	if err != nil {
		return err
	}

	tx := db.DB.Model(&event).Updates(Events{Name: e.Name, Description: e.Description, Location: e.Location, UserId: e.UserId, EventDate: e.EventDate})
	return tx.Error
}

func DeleteEvents(id int64) error {

	event, err := GetEventById(id)
	if err != nil {
		return err
	}

	tx := db.DB.Delete(&event, id)

	return tx.Error
}

func (e *Events) RegisterEvent(userID int64) error {

	tx := db.DB.Create(&db.Registrations{UserID: uint(userID), EventID: uint(e.Id)})

	return tx.Error

}

func (e *Events) CancleRegistration(userId int64) (int64, error) {

	var registerEvent db.Registrations
	tx := db.DB.Find(&registerEvent, "event_id = ? and user_id = ?", e.Id, userId)

	if tx.Error != nil {
		return -1, tx.Error
	}

	tx = db.DB.Delete(&registerEvent)

	return tx.RowsAffected, tx.Error

}
