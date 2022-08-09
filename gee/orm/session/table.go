package session

import (
	"fmt"
	"go-knowledge/gee/orm/log"
	"go-knowledge/gee/orm/schema"
	"reflect"
	"strings"
)

//用来防止操作数据库相关的代码

// Model  方法用于给 refTable 赋值。
//解析操作是比较耗时的，因此将解析的结果保存在成员变量 refTable 中，即使 Model() 被调用多次，如果传入的结构体名称不发生变化，则不会更新 refTable 的值。
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// RefTable  方法返回 refTable 的值，如果 refTable 未被赋值，则打印错误日志。
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

// CreateTable 创建表
func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	return err
}

// DropTable 删除表
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXIST %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable 判断表是否存在
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
