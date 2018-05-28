package dao

import (
	"os"
	"testing"
)

var todoDao TodoDao

func TestMain(m *testing.M) {

	// master & slave
	dbmConfig := DbConfig{
		User:     "yatteiki_admin",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "yatteiki_cloud",
	}

	// slave
	dbsConfig := DbConfig{
		User:     "yatteiki_admin",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "yatteiki_cloud",
	}

	// setup database
	dbm = setupDbMap(dbmConfig)
	dbs = setupDbMap(dbsConfig)

	// setup tables
	todoDao = &todoDaoImpl{baseDao: NewDao("todos", []string{"id"}, []string{"id", "note"})}

	os.Exit(m.Run())
}
