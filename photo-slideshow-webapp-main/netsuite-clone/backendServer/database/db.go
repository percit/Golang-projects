package database
import (
	"fmt"
)


//when everything works with routes, add mongoDB here
//https://github.com/anthdm/catfacter/blob/master/main.go
//tu powyzej jest przyklad wlasnie mongoDB

type DB struct {
	database map[string]int//database that will take date as string, and number of hours
}

func (db *DB) InitDB() {
	db.database = make(map[string]int)
	db.database["24.04.2023"] = 8
	db.database["25.04.2023"] = 8
	db.database["26.04.2023"] = 8
	db.database["27.04.2023"] = 8
}

func (db *DB) GetHoursByDate(date string) int {
	fmt.Println("getHoursByDate")
	return db.database[date]
}

func (db *DB) SetHoursByDate(date string, time int) {
	fmt.Println("setHoursByDate")
	db.database[date] = time
}

func (db *DB) deleteDate(date string) {
	fmt.Println("deleteDate")
	delete(db.database, date)
}