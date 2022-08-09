package session

import (
	"database/sql"
	"go-knowledge/gee/orm/clause"
	"go-knowledge/gee/orm/dialect"
	"go-knowledge/gee/orm/log"
	"go-knowledge/gee/orm/schema"
	"strings"
)

//用于实现与数据库的交互的封装
//封装有 2 个目的，一是统一打印日志（包括 执行的SQL 语句和错误日志）。
//二是执行完成后，清空 (s *Session).sql 和 (s *Session).sqlVars 两个变量。这样 Session 可以复用，开启一次会话，可以执行多次 SQL。

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect //
	tx       *sql.Tx         //事务相关的操作
	refTable *schema.Schema  //操作的对象
	clause   clause.Clause   //操作的类型
	sql      strings.Builder
	sqlVars  []interface{}
}

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Rows
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw
//Session 结构体目前只包含三个成员变量，第一个是 db *sql.DB，即使用 sql.Open() 方法连接数据库成功之后返回的指针。
//第二个和第三个成员变量用来拼接 SQL 语句和 SQL 语句中占位符的对应值。用户调用 Raw() 方法即可改变这两个变量的值。
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec 执行sql的方法
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 获取单条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars)
}

// QueryRows 从数据库中查询一个列表
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
