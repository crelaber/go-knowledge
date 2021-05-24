package datastruct

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array :", a.Print(a))

	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array :", a.Print(a))

	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array :", a.Print(a))

	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array :", a.Print(a))

}
