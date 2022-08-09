package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	novel2 "go-knowledge/cmd/novel"
	"go-knowledge/cmd/novel/category"
	. "go-knowledge/cmd/novel/model"
	"io/ioutil"
	"os"
	"sort"
)

var (
	novel string
)

var dlCmd = &cobra.Command{
	Use:     "yunting",
	Short:   "抓取云听的音频",
	Long:    "抓取云听的音频",
	Example: "go-knowledge yunting -n douluo",
	PreRun:  checkNovelType,
	Run:     entrance,
}

func init() {
	dlCmd.PersistentFlags().StringVarP(&novel, "novel", "n", novel, "novel类型")
	RootCmd.AddCommand(dlCmd)
}

//检查输入的小说类型参数是否正确
func checkNovelType(cmd *cobra.Command, args []string) {
	isValid := novel2.IsNovelValid(novel)
	if !isValid {
		fmt.Printf("小说类型不正确，必须是以下几个的一个：%v", novel2.GetNovelType())
		return
	}
}

//入口方法
func entrance(cmd *cobra.Command, args []string) {
	//fileList := getNovelResource()
	//basePath := novel2.GetNovelBasePath(novel)
	fileList := novel2.GetNovelJsonFiles(novel)
	playList := make([]PlayList, 0)
	novelName := ""
	for _, jsonFile := range fileList {
		//path := basePath + file
		//fmt.Println("path：" + path)
		//list, novelName := getNovelList(path)
		list, name1 := getNovelList(jsonFile)
		if novelName == "" {
			novelName = name1
		}
		playList = append(playList, list...)
	}
	printNovelInfo(playList, novelName)
}

func getNovelResource() []string {
	//定义接口示例
	var inst category.NovelResource
	switch novel {
	case novel2.NovelRenShiJian:
		inst = category.NewNovelRenshijian()
	case novel2.NovelDouLuo:
		inst = category.NewNovelDouLuo()
	case novel2.NovelSanti:
		inst = category.NewNovelSanti()
	}
	return inst.GetResourceList()
}

func getNovelList(file string) ([]PlayList, string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Printf("open file error:%v", err)
		return []PlayList{}, ""
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error read json file")
		return []PlayList{}, ""
	}
	dbNovel := TingShuNovel{}
	json.Unmarshal(jsonData, &dbNovel)

	result := getPlayList(dbNovel)
	return result, dbNovel.ColumnName
}

func getPlayList(novel TingShuNovel) []PlayList {
	if len(novel.PlayList) == 0 {
		return []PlayList{}
	}

	list := novel.PlayList
	var result []PlayList
	for _, novel := range list {
		//name := novel.Name
		//nameArr := strings.Split(name, " ")
		//if len(nameArr) > 0 {
		//	playIndex, err := strconv.Atoi(nameArr[1])
		//	if err != nil {
		//		playIndex = 0
		//	}
		//	novel.PlayIndex = playIndex
		//}
		novel.PlayIndex = novel.Index
		result = append(result, novel)
	}
	return result
}

//打印语音的信息
func printNovelInfo(list []PlayList, novelName string) {
	if len(list) == 0 {
		return
	}
	var keys []int
	//var playMap map[int]PlayList
	playMap := make(map[int]PlayList, 0)
	for _, k := range list {
		playMap[k.PlayIndex] = k
		keys = append(keys, k.PlayIndex)
	}
	sort.Ints(keys)
	for _, v := range keys {
		if novel, ok := playMap[v]; ok {
			fmt.Printf("%s第%d集：%s", novelName, novel.PlayIndex, novel.PlayURL)
			fmt.Println()
		}
	}
}
