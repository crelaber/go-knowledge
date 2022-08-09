package schema

import (
	"go-knowledge/gee/orm/dialect"
	"go/ast"
	"reflect"
)

//Field 包含 3 个成员变量，字段名 Name、类型 Type、和约束条件 Tag
type Field struct {
	Name string
	Type string
	Tag  string
}

//Schema 主要包含被映射的对象 Model、表名 Name 和字段 Fields。
//FieldNames 包含所有的字段名(列名)，fieldMap 记录字段名和 Field 的映射关系，方便之后直接使用，无需遍历 Fields。

//TypeOf() 和 ValueOf() 是 reflect 包最为基本也是最重要的 2 个方法，分别用来返回入参的类型和值。因为设计的入参是一个对象的指针，因此需要 reflect.Indirect() 获取指针指向的实例。
//modelType.Name() 获取到结构体的名称作为表名。
//NumField() 获取实例的字段的个数，然后通过下标获取到特定字段 p := modelType.Field(i)。
//p.Name 即字段名，p.Type 即字段类型，通过 (Dialect).DataTypeOf() 转换为数据库的字段类型，p.Tag 即额外的约束条件。

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// RecordValues 为insert做铺垫的函数
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
