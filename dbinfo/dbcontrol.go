package dbinfo

import (
	"database/sql"
	"log"
	"strconv"

	"gitlab.com/ServerUtility/code"
	"gitlab.com/ServerUtility/messagehandle"
)

type SqlCLi struct {
	DB *sql.DB
}
type SqlQuary struct {
	Quary string
	Args  []interface{}
}

// CallRead call stored procedure
func CallRead(db *sql.DB, name string, args ...interface{}) ([]interface{}, messagehandle.ErrorMsg) {
	QueryStr := MakeProcedureQueryStr(name, len(args))
	request, err := Query(db, QueryStr, args...)
	return request, err
}

// CallReadOutMap call stored procedure
func CallReadOutMap(db *sql.DB, name string, args ...interface{}) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	QueryStr := MakeProcedureQueryStr(name, len(args))
	request, err := QueryOutMap(db, QueryStr, args...)
	return request, err

}

// CallReadOutMultipleMap call stored procedure
func CallReadOutMultipleMap(db *sql.DB, name string, args ...interface{}) ([][]map[string]interface{}, messagehandle.ErrorMsg) {
	QueryStr := MakeProcedureQueryStr(name, len(args))
	request, err := QueryMultipleResult(db, QueryStr, args...)
	return request, err

}

// CallWrite ...
func CallWrite(db *sql.DB, name string, args ...interface{}) (sql.Result, messagehandle.ErrorMsg) {
	request, err := Exec(db, name, args...)
	return request, err
}

// Query Use to SELECT return array, first is Keys
func Query(db *sql.DB, query string, args ...interface{}) ([]interface{}, messagehandle.ErrorMsg) {
	err := messagehandle.New()

	res, errMsg := db.Query(query, args...)
	if errMsg != nil {
		err.ErrorCode = code.FailedPrecondition
		err.Msg = "DBExecFail"
		messagehandle.ErrorLogPrintln("DB", errMsg, query, args)
		return nil, err
	}

	request := MakeScanArray(res)
	defer res.Close()

	return request, err
}

// QueryMultipleResult ...
func QueryMultipleResult(db *sql.DB, query string, args ...interface{}) ([][]map[string]interface{}, messagehandle.ErrorMsg) {
	err := messagehandle.New()

	res, errMsg := db.Query(query, args...)
	if errMsg != nil {
		err.ErrorCode = code.FailedPrecondition
		err.Msg = "DBExecFail"
		messagehandle.ErrorLogPrintln("DB", errMsg, query, args)
		return nil, err
	}

	var request [][]map[string]interface{}
	var resultMap []map[string]interface{}

	resultMap = MakeScanMap(res)
	request = append(request, resultMap)

	for res.NextResultSet() {
		resultMap = MakeScanMap(res)
		request = append(request, resultMap)
	}
	defer res.Close()

	return request, err

}

// QueryOutMap Use to SELECT return array map
func QueryOutMap(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, messagehandle.ErrorMsg) {
	err := messagehandle.New()

	res, errMsg := db.Query(query, args...)
	if errMsg != nil {
		err.ErrorCode = code.FailedPrecondition
		err.Msg = "DBExecFail"
		messagehandle.ErrorLogPrintln("DB", errMsg, query, args)
		return nil, err
	}

	request := MakeScanMap(res)
	defer res.Close()

	return request, err
}

// Exec Use to INSTER, UPDATE, DELETE
func Exec(db *sql.DB, query string, args ...interface{}) (sql.Result, messagehandle.ErrorMsg) {
	err := messagehandle.New()

	res, errMsg := db.Exec(query, args...)
	if errMsg != nil {
		err.ErrorCode = code.FailedPrecondition
		err.Msg = "DBExecFail"
		messagehandle.ErrorLogPrintln("DB", errMsg, query, args)
		return nil, err
	}
	return res, err
}

// MakeProcedureQueryStr ...
func MakeProcedureQueryStr(name string, args int) string {
	query := "CALL " + name + "("

	if args > 0 {
		for i := 0; i < args; i++ {
			query += "?,"
		}
		query = query[:len(query)-1]
	}
	query += ");"
	return query
}

// MakeScanArray ...
func MakeScanArray(rows *sql.Rows) []interface{} {
	var Result []interface{}
	Keys, err := rows.Columns()
	Result = append(Result, Keys)

	for rows.Next() {
		Row := make([]interface{}, len(Keys))
		for i := range Keys {
			Row[i] = new(sql.RawBytes)
		}

		err = rows.Scan(Row...)
		if err != nil {
			log.Fatalln(err)
		}

		Result = append(Result, Row)

	}

	return Result
}

// MakeScanMap ...
func MakeScanMap(rows *sql.Rows) []map[string]interface{} {
	Keys, err := rows.Columns()
	types, err := rows.ColumnTypes()
	scanArgs := make([]interface{}, len(Keys))
	values := make([][]byte, len(Keys))
	Result := []map[string]interface{}{}

	for rows.Next() {
		for i := range Keys {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatalln(err)
		}

		tomap := map[string]interface{}{}
		for i, key := range Keys {
			if types[i].DatabaseTypeName() == "BIGINT" {
				val, errMsg := strconv.ParseInt(string(values[i]), 10, 64)
				if errMsg != nil {
					panic("makeScanMap Error BIGINT")
				}
				tomap[key] = val
			} else if types[i].DatabaseTypeName() == "INT" {
				val, errMsg := strconv.ParseInt(string(values[i]), 10, 0)
				if errMsg != nil {
					panic("makeScanMap Error INT")
				}
				tomap[key] = val
			} else if types[i].DatabaseTypeName() == "TINYINT" {
				val, errMsg := strconv.ParseInt(string(values[i]), 10, 0)
				if errMsg != nil {
					panic("makeScanMap Error TINYINT")
				}
				tomap[key] = val == 1
			} else if types[i].DatabaseTypeName() == "VARCHAR" {
				tomap[key] = string(values[i])
			} else if types[i].DatabaseTypeName() == "TEXT" {
				tomap[key] = string(values[i])
			} else {
				tomap[key] = values[i]
			}
		}

		Result = append(Result, tomap)
	}

	return Result
}
