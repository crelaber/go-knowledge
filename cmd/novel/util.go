package novel

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// GetNovelType 获取小说的类型
func GetNovelType() []string {
	return []string{
		NovelRenShiJian,
		NovelDouLuo,
		NovelSanti,
	}
}

// GetNovelBasePath 获取json资源文件的路径
func GetNovelBasePath(novelType string) string {
	return fmt.Sprintf("./cmd/novel/resource/%s/", novelType)
}

//获取小说所有的json文件的路径
func GetNovelJsonFiles(novelType string) []string {
	path := GetNovelBasePath(novelType)
	_, fileList := GetDirFiles(path)
	var result []string
	for _, name := range fileList {
		filePath := path + name
		result = append(result, filePath)
	}
	return result
}

// IsNovelValid 检查小说类型是否合法
func IsNovelValid(novel string) bool {
	novelTypes := GetNovelType()
	for _, v := range novelTypes {
		if v == novel {
			return true
		}
	}
	return false
}

func GetDirFiles(dir string) ([]fs.FileInfo, []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln("read dir error")
	}

	var fileNames []string
	var myFile []fs.FileInfo
	for _, file := range files {
		if file.IsDir() {
			path := strings.TrimSuffix(dir, "/") + "/" + file.Name()
			subFile, _ := GetDirFiles(path)
			if len(subFile) > 0 {
				myFile = append(myFile, subFile...)
			}
		} else {
			myFile = append(myFile, file)
			fileNames = append(fileNames, file.Name())
		}
	}
	return myFile, fileNames
}

// GetDirFiles2 遍历文件夹下所有文件
func GetDirFiles2(dir string) ([]string, error) {
	var files []string
	var walkFunc = func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}
	err := filepath.Walk(dir, walkFunc)
	return files, err
}
