package Helper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

//把数据结构 映射成 map切片
func DBMap(columns []string, rows *sql.Rows) ([]interface{}, error) {
	allRows := make([]interface{}, 0) //所有行  大切片
	for rows.Next() {
		oneRow := make([]interface{}, len(columns)) //定义一行切片
		scanRow := make([]interface{}, len(columns))
		fieldMap := make(map[string]interface{})
		for i, _ := range oneRow {
			scanRow[i] = &oneRow[i]
		}
		err := rows.Scan(scanRow...)
		if err != nil {
			return nil, err
		}
		for i, val := range oneRow {
			v, ok := val.([]byte) //断言
			if ok {
				fieldMap[columns[i]] = string(v)
			}
		}
		allRows = append(allRows, fieldMap)
	}
	return allRows, nil
}

type DBModel struct {
	TableName string
	Data      []interface{} //map切片
}
type MyDB struct {
	*sqlx.DB
}

//获取表结构
func (this *MyDB) DescTable(tableName string, prefix string) (*DBModel, error) {
	rows, err := this.Query("desc " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	data, err := DBMap(columns, rows)
	if err != nil {
		return nil, err
	}
	tableName = strings.Replace(tableName, prefix, "", 1)
	return &DBModel{Data: data, TableName: tableName}, nil

}

func NewDB(driver string, dsn string) *MyDB {
	db, err := sqlx.Connect(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &MyDB{db}
}
