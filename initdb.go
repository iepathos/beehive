package main

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

// database name
var DbName = "beehive"

// app table names
var AppTables = []string{
	"users",
	"messages",
	"questions",
	"answers",
}

func createDatabase(databaseName string) {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Creating database", databaseName)
	_, err = r.DBCreate(databaseName).Run(session)
	if err != nil {
		log.Println(err.Error())
	}
}

func createTable(tableName string) {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db := r.DB(DbName)

	log.Println("Creating table", tableName)
	if _, err := db.TableCreate(tableName).RunWrite(session); err != nil {
		log.Println(err)
	}
}

func main() {
	createDatabase(DbName)

	for _, v := range AppTables {
		createTable(v)
	}
}
