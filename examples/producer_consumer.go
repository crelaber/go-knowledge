package examples

import "fmt"

type OrderInfo struct {
	id int
}

func producerOrder(out chan<- OrderInfo) {
	for i := 0; i < 20; i++ {
		order := OrderInfo{
			id: i + 1,
		}
		fmt.Println("生成订单，订单id为：", order.id)
		out <- order //写入channel
	}
	close(out)
}

func consumeOrder(in <-chan OrderInfo) {
	for order := range in {
		fmt.Println("读取订单id，id为：", order.id)
	}
}
