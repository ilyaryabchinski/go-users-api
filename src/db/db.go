package db

import "github.com/syndtr/goleveldb/leveldb"

var database *leveldb.DB

func ConnectToDatabase(path string) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic("Database cannot be connected")
	}

	database = db
}

func GetDatabaseInstance() *leveldb.DB {
	return database
}

func CloseDatabaseConnection() error {
	return database.Close()
}
