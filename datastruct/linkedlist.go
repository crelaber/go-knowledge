package datastruct

import "fmt"

type LinkListInst struct {
}

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func (l *LinkListInst) Demo() {
	node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.NextNode
		}
		break
	}
}
