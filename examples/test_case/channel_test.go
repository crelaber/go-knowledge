package test_case

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			info := fmt.Sprintf("current index : %d", i)
			fmt.Println(info)
			out <- rand.Intn(5)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
