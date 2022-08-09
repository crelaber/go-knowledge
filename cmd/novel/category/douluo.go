package category

type DouluoNovel struct {
}

func NewNovelDouLuo() *DouluoNovel {
	return &DouluoNovel{}
}

// GetResourceList 这里只需要把语音对应的集数文件名称写出来就可以，在命令行里面会做拼接，真是的资源文件在cmd/novel/resource目录下
func (d *DouluoNovel) GetResourceList() []string {
	m := []string{
		"992.json",
		"1012.json",
		"1032.json",
		"1052.json",
		"1072.json",
		"1092.json",
		"1112.json",
	}
	return m
}

//实现type接口用于指定资源文件的路径
var _ NovelResource = (*DouluoNovel)(nil)
