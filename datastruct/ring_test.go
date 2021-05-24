package datastruct

import (
	"fmt"
	"testing"
)

func TestAddRing(t *testing.T) {
	fmt.Println("testing add ring")
	r := &Ring{Value: -1}
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		fmt.Println(node.Value)
		node = node.Next()
		//如果回到了起点，则结束
		if node == r {
			return
		}
	}
}

func TestDeleteRing(t *testing.T) {
	fmt.Println("testing delete ring")
	r := &Ring{Value: -1}
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	tmp := r.Unlink(3)
	node := r
	for {
		fmt.Println(node.Value)
		node = node.Next()
		if node == r {
			break
		}
	}
	fmt.Println("---------")

	node = tmp
	for {
		fmt.Println(node.Value)
		node = node.Next()

		if node == tmp {
			break
		}
	}
}
