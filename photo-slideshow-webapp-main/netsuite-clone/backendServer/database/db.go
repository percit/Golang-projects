package database
import (
	"fmt"
)


//when everything works with routes, add mongoDB here

type DB struct {
	database map[string]int//database that will take date as string, and number of hours
}

func initDB() {
	db := DB{database: make(map[string]int)}
	db.database["24.04.2023"] = 8
	db.database["25.04.2023"] = 8
	db.database["26.04.2023"] = 8
	db.database["27.04.2023"] = 8
}

func (db *DB) getHoursByDate(date string) int {
	return db.database[date]
}

func (db *DB) setHoursByDate(date string) int {
	return db.database[date]
}

func (db *DB) deleteDate(date string) {
	delete(db.database, date)
}