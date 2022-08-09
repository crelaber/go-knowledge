package model

import "database/sql"

type TableSchema struct {
	TableName              string
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           string
	ColumnType             string
	ColumnKey              string
	Comment                string
}

type ColumnSchema struct {
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	ColumnType             string
	ColumnKey              string
	Comment                string
}

type Table struct {
	Imports         map[string]struct{} //import表
	HasPrimaryKey   bool                //是否主键
	CamelPrimaryKey string              //转驼峰的主键字段
	PrimaryKey      string              //主键字段名
	PrimaryKeyType  string              //主键字段类型
	Columns         []Column            //所有字段
}

type Column struct {
	Name      string //字段名
	CamelName string //驼峰字段名
	Type      string //mysql中原始数据类型
	ColumnKey string //primary说明是主键
	Comment   string //mysql中原始注释

	GoType    string  //go结构体字段类型
	GoJsonTag string  //go结构体中的json标签
	GoComment Comment //从注释中json反序列化的Comment
}

type Comment struct {
	Data string `json:"data"`
	Type string `json:"type"`
}
