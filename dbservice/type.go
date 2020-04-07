package dbservice

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.fbk168.com/gamedevjp/backend-utility/utility/messagehandle"
)

// ConnSetting ...
type ConnSetting struct {
	DBUser, DBPassword, DBIP, DBPORT string
}

// // IDB ...
// type IDB interface {
// 	SetSetting(ConnSetting)
// 	GetDB() *sql.DB
// 	ConnectDB(ConnSetting) (*sql.DB, error)
// }

// DB ...
type DB struct {
	setting ConnSetting
	conn    *sql.DB
}

// SetSetting ...
func (db *DB) SetSetting(setting ConnSetting) {
	db.setting = setting
}

// GetDB ...
func (db *DB) GetDB() *sql.DB {
	return db.conn
}

// ConnectDB ...
func (db *DB) ConnectDB(dbName string, setting ConnSetting) error {
	if db.conn == nil {
		sqlstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=30s", setting.DBUser, setting.DBPassword, setting.DBIP, setting.DBPORT, dbName)
		messagehandle.LogPrintln("DB Connect:", sqlstr)
		sqldb, err := sql.Open("mysql", sqlstr)
		if err != nil {
			return err
		}

		connMaxLifetime := 59 * time.Second
		maxIdleConns := 50
		maxOpenConns := 50

		messagehandle.LogPrintf("connMaxLifetime:%d\n", connMaxLifetime/time.Second)
		sqldb.SetConnMaxLifetime(time.Duration(connMaxLifetime))

		messagehandle.LogPrintf("maxIdleConns:%d\n", maxIdleConns)
		sqldb.SetMaxIdleConns(maxIdleConns)

		messagehandle.LogPrintf("maxOpenConns:%d\n", maxOpenConns)
		sqldb.SetMaxOpenConns(maxOpenConns)

		db.conn = sqldb
	}

	return nil
}
