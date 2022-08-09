package test_case

import (
	"fmt"
	"testing"
	"time"
)

//循环调用panic保证程序不退出
func TestPanicTrigger(t *testing.T) {
	go func() {
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()

			}
		}
	}()

}

func proc() {
	panic("ok")
}
