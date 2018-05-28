package dao

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"ghe.ca-tools.org/valencia/valencia-api/modelgen"
	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

// DbConfig database config struct
type DbConfig struct {
	Host            string
	User            string
	Password        string
	Dbname          string
	Port            string
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type Dao struct {
}

func newDao(dbm, dbs *gorp.DbMap) *Dao {
	m := modelgen.Auth{}
	tableName := m.TableName()
	pks := m.PrimaryKeys()
	dbm.AddTableWithName(m, tableName).SetKeys(true, pks...)
	dbs.AddTableWithName(m, tableName).SetKeys(true, pks...)
	dao := AuthDao{}
	dao.baseDao = newBaseDao(dbm, dbs)
	dao.tableName = tableName
	dao.columnsName = strings.Join(m.ColumnNames(), ",")
	return &dao
}

func hoge() *gorp.DbMap {
	c := DbConfig{
		User:     "root",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "yatteiki_cloud",
	}
	// see also https://github.com/go-sql-driver/mysql#timetime-support
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.Dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(c.ConnMaxLifetime)

	return &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}
}
