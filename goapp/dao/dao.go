package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	gorp "gopkg.in/gorp.v1"

	"github.com/suzujun/yatteiki-cloud/goapp/config"
)

type (
	baseDao struct {
		dbm         *gorp.DbMap
		dbs         *gorp.DbMap
		tableName   string
		primaryKeys []string
		columnsName string
	}
	// Table interface
	Table interface {
		Name() string
		PrimaryKeys() []string
		ColumnNames() []string
	}
)

var dbm, dbs *gorp.DbMap

func Initialize(dbmConfig, dbsConfig *config.DbConfig) {
	// setup database
	dbm = setupDbMap(dbmConfig)
	dbs = setupDbMap(dbsConfig)
}

// NewDao is new dao
func newDao(table Table) baseDao {

	dao := baseDao{
		dbm:         dbm,
		dbs:         dbs,
		tableName:   table.Name(),
		primaryKeys: table.PrimaryKeys(),
		columnsName: strings.Join(table.ColumnNames(), ","),
	}
	dbm.AddTableWithName(table, table.Name()).SetKeys(true, table.PrimaryKeys()...)
	dbs.AddTableWithName(table, table.Name()).SetKeys(true, table.PrimaryKeys()...)
	return dao
}

func setupDbMap(c *config.DbConfig) *gorp.DbMap {
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

	dbmap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8mb4",
		},
	}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbmap
}

// PingDb is ping database
func PingDb() bool {
	if err := dbm.Db.Ping(); err != nil {
		fmt.Printf("failed to ping to dbm error: %+v\n", err)
		return false
	}
	if err := dbs.Db.Ping(); err != nil {
		fmt.Printf("failed to ping to dbs errpr: %+v\n", err)
		return false
	}
	return true
}

func (dao baseDao) newSelectBuilder() sq.SelectBuilder {
	return sq.Select(dao.columnsName).From(dao.tableName)
}

func (dao baseDao) newUpdateBuilder() sq.UpdateBuilder {
	return sq.Update(dao.tableName)
}

func (dao baseDao) newDeleteBuilder() sq.DeleteBuilder {
	return sq.Delete(dao.tableName)
}

func (dao baseDao) findOneByBuilder(builder *sq.SelectBuilder, target interface{}) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrapf(err, "build sql failed [sql='%s'][args='%+v']", sql, args)
	}
	if err := dao.dbs.SelectOne(target, sql, args...); err != nil {
		return errors.Wrapf(err, "fetch data failed [sql='%s'][args='%+v']", sql, args)
	}
	return nil
}

func (dao baseDao) findManyByBuilder(builder *sq.SelectBuilder, targets interface{}) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrapf(err, "build sql failed [sql='%s'][args='%+v']", sql, args)
	}
	if _, err := dao.dbs.Select(targets, sql, args...); err != nil {
		return errors.Wrapf(err, "fetch data failed [sql='%s'][args='%+v']", sql, args)
	}
	return nil
}

func (dao baseDao) insert(target interface{}) error {
	return errors.Wrapf(dao.dbm.Insert(target), "insert failed [%+v]", dao.tableName)
}

func (dao baseDao) update(target interface{}) error {
	_, err := dao.dbm.Update(target)
	return errors.Wrapf(err, "update failed [%+v]", dao.tableName)
}

func (dao baseDao) updateByBuilder(builder *sq.UpdateBuilder) (sql.Result, error) {
	builder2 := builder.Set("updated_at", time.Now())
	sql, args, err := builder2.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "build update sql failed [sql='%s'][args='%+v']", sql, args)
	}
	result, err := dao.dbm.Exec(sql, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "update data failed [sql='%s'][args='%+v']", sql, args)
	}
	return result, nil
}

func (dao baseDao) delete(target interface{}) error {
	_, err := dao.dbm.Delete(target)
	return errors.Wrapf(err, "delete failed [%+v]", dao.tableName)
}

func (dao baseDao) deleteByBuilder(builder *sq.DeleteBuilder) (sql.Result, error) {
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "build delete sql failed [sql='%s'][args='%+v']", sql, args)
	}
	result, err := dao.dbm.Exec(sql, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "delete data failed [sql='%s'][args='%+v']", sql, args)
	}
	return result, nil
}
