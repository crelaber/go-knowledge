package category

type SantiNovel struct {
}

func NewNovelSanti() *SantiNovel {
	return &SantiNovel{}
}

// GetResourceList 这里只需要把语音对应的集数文件名称写出来就可以，在命令行里面会做拼接，真是的资源文件在cmd/novel/resource目录下
func (d *SantiNovel) GetResourceList() []string {
	m := []string{
		"1-20.json",
		"21-40.json",
		"41-60.json",
	}
	return m
}

//实现type接口用于指定资源文件的路径
var _ NovelResource = (*SantiNovel)(nil)
