package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("mysql", "root:Deepak@1492@/PLANB")

	DB = db

	if err != nil {
		panic(err)
	}

	/*
		      The db connections are opened lazyily, the connections open only while processing the new request.
			  if the client request reaches MaxOpenConnections than others have to wait till there is open connection in the pool.
			  if there is no more new requests than its trims down the open connections to MaxIdelConnections
	*/

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createUsers := `
	    CREATE TABLE IF NOT EXISTS USERS(
		    id INT(10) primary key auto_increment,
			email varchar(500) not null unique,
			password varchar(500) not null
		)
	`

	_, err := DB.Exec(createUsers)

	if err != nil {
		panic(err)
	}

	createEvents := `
	    create table if not exists events(
        id INT(10) auto_increment, 
        name varchar(500), 
        description varchar(500), 
        location varchar(500), 
        event_date DATETIME, 
        user_id INT(10),
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users(id)
        )
	`

	_, err = DB.Exec(createEvents)

	if err != nil {
		panic(err)
	}

	createRegistration := `
	    create table if not exists registrations (
        id INT(10) auto_increment, 
        user_id int(10), 
        event_id int(10), 
		PRIMARY KEY (id),
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
       )
	`
    
	_, err = DB.Exec(createRegistration)
    
	if err != nil {
		panic(err)
	}

}
