package dialect

import (
	"fmt"
	"reflect"
	"time"
)

var _ Dialect = (*sqlite3)(nil)

type sqlite3 struct {
}

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

//DataTypeOf 将 Go 语言的类型映射为 SQLite 的数据类型
func (s sqlite3) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s(%s)", typ.Type().Name(), typ.Kind()))
}

//TableExistSQL 返回了在 SQLite 中判断表 tableName 是否存在的 SQL 语句。
func (s sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select name from sqlite_master where type = 'table' and name = ?", args
}
