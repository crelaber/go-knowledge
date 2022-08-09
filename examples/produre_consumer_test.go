package examples

import (
	"testing"
	"time"
)

func TestProducerConsumer(t *testing.T) {
	ch := make(chan OrderInfo, 20)
	go producerOrder(ch)
	go consumeOrder(ch)
	time.Sleep(2 * time.Second)
}
