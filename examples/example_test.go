package examples

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestForeach(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

}

func TestHello(t *testing.T) {
	go func() {
		fmt.Println("Hello")
	}()
	go func() {
		fmt.Println("World")
	}()

	select {}
}

func TestLock(t *testing.T) {
	var mutext sync.Mutex
	counter := 0
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutext.Lock()
			defer mutext.Unlock()
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func TestCurrentMapReadWrite(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var m = map[int]int{
		1: 1,
	}
	//read
	go func(m map[int]int) {
		for {
			fmt.Println(m[1])
		}
	}(m)
	//write
	go func(m map[int]int) {
		for {
			m[1] = 1
		}
	}(m)

	select {}
}

func TestCurrentMapWriteDelete(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var m = map[int]int{
		1: 1,
	}
	//write
	go func(m map[int]int) {
		for {
			m[1] = 1
		}
	}(m)

	//delete
	go func(m map[int]int) {
		for {
			delete(m, 1)
		}
	}(m)

	select {}
}

func TestCurrentMapReadRead(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var m = map[int]int{
		1: 1,
	}
	//read
	go func(m map[int]int) {
		for {
			fmt.Println(m[1])
		}
	}(m)
	//read
	go func(m map[int]int) {
		for {
			fmt.Println(m[1])
		}
	}(m)

	select {}
}

func TestAdd(t *testing.T) {

	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}

func TestChan(t *testing.T) {
	channel := make(chan string)
	go sendData(channel)
	go getData(channel)
}

func sendData(channel1 chan string) {
	channel1 <- "a"
	channel1 <- "b"
	channel1 <- "c"
	channel1 <- "d"
	channel1 <- "e"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}

func TestOutput(t *testing.T) {
	a := "012"
	pa := &a
	b := []byte(*pa)
	pb := &b

	a += "3"
	*pa += "4"
	b[1] = '5'
	println(*pa)
	println(string(*pb))
}

func TestOutput2(t *testing.T) {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

var a int
var c = make(chan int)

func TestOutput3(t *testing.T) {
	go f()
	c <- 0
	print(a)
}

func f() {
	a = 1
	<-c
}

func TestOutput4(t *testing.T) {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	strChan := make(chan string, 1)
	intChan <- 1
	strChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-strChan:
		panic(value)

	}
}
