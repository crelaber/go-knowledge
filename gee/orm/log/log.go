package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	//[info ] 颜色为蓝色，[error] 为红色。
	errLog  = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errLog, infoLog}
	mu      sync.Mutex
)

var (
	Error  = errLog.Println
	Errorf = errLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

//日志的层级
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel 设置日志层级
//这一部分的实现非常简单，三个层级声明为三个常量，通过控制 Output，来控制日志是否打印。
//如果设置为 ErrorLevel，infoLog 的输出会被定向到 ioutil.Discard，即不打印该日志。
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
