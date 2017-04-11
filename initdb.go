package main

import "github.com/iepathos/beehive/rego"

// database name
var DbName = "beehive"

// app table names
var AppTables = []string{
	"users",
	"messages",
	"questions",
	"answers",
}

func main() {
	rego.CreateDatabase(DbName)

	for _, v := range AppTables {
		rego.CreateTable(v)
	}
}
