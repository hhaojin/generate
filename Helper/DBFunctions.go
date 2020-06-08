package Helper

import (
	"regexp"
	"strings"
)

var (
	DBTypeMap = map[string]string{
		"string":    "varchar,char,text,tinytext,mediumtext",
		"int":       "int,tinyint,smallint,mediumint,integer,year,boolean",
		"int64":     "bigint,datetime,timestamp",
		"bool":      "bit",
		"float":     "float,double,decimal",
		"[]byte":    "blob,longblob,mediumblob",
		"time.Time": "date,time",
	}
)

type typelist []string

func (this typelist) contains(key string) bool {
	for _, item := range this {
		if item == key {
			return true
		}
	}
	return false
}
func MapDBType(t string) string {
	regx := regexp.MustCompile(`\(\d+\)`)
	t = strings.ToLower(regx.ReplaceAllString(t, ""))
	for k, v := range DBTypeMap {
		tlist := typelist(strings.Split(v, ","))
		if tlist.contains(t) {
			return k
		}
	}
	return "string"
}
