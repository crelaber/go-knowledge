package schema

import (
	"go-knowledge/gee/orm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var testDialect, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, testDialect)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse ")
	}

	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
