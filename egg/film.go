package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	//getDOubanMovie()
	go func() {
		for {
			log.Println("currnet =======>")
		}
	}()
	http.ListenAndServe("0.0.0.0:60600", nil)

}

func getDOubanMovie() {
	//service.Start()
	//defer model.DB.Close()
}

func traceTest() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "test"
	}()
	<-ch
}
