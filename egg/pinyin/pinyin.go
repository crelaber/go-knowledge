package pinyin

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"testing"
)

func TestPinyin(t *testing.T) {
	hans := "我是中国人"
	args := pinyin.NewArgs()
	str := pinyin.Pinyin(hans, args)
	fmt.Printf("%s", str)
}

func main() {
	hans := "我是中国人"
	args := pinyin.NewArgs()
	str := pinyin.Pinyin(hans, args)
	fmt.Printf("%s", str)
}
