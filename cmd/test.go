package cmd

import (
	"github.com/spf13/cobra"
)

var (
	testType string //测试的类型
)

const (
	categoryDir string = "dir_file"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "测试各种用例",
	Long:    "测试各种用例",
	Example: "go-knowledge test -tt dir",
	PreRun:  checkCategory,
	Run:     testEntrance,
}

func init() {
	testCmd.PersistentFlags().StringVarP(&testType, "testType", "t", testType, "类型")
	RootCmd.AddCommand(testCmd)
}

//入口方法
func testEntrance(cmd *cobra.Command, args []string) {
	switch testType {
	case categoryDir:
		printFileList()
	default:
		printFileList()
	}
}

func printFileList() {
	//novelPath := "douluo"
	//path := fmt.Sprintf("./cmd/novel/resource/%s/", novelPath)
	////path := "./cmd"
	//_, fileNames := FileForeach(path)
	//for _,name := range fileNames {
	//	fmt.Println(name)
	//}
}

func checkCategory(cmd *cobra.Command, args []string) {
}
