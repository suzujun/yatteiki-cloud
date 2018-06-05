package dao

import (
	"os"
	"testing"

	"github.com/suzujun/yatteiki-cloud/goapp/model"
)

var todoDao TodoDao

func TestMain(m *testing.M) {

	// master & slave
	dbmConfig := DbConfig{
		User:     "yatteiki_admin",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "yatteiki",
	}

	// slave
	dbsConfig := DbConfig{
		User:     "yatteiki_admin",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "yatteiki",
	}

	// setup database
	dbm = setupDbMap(dbmConfig)
	dbs = setupDbMap(dbsConfig)

	// setup tables
	todoDao = &todoDaoImpl{baseDao: newDao(model.Todo{})}

	os.Exit(m.Run())
}
