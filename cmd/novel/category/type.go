package category

// NovelResource 资源的接口
type NovelResource interface {
	// GetResourceList 这里只需要把语音对应的集数文件名称写出来就可以，在命令行里面会做拼接，真是的资源文件在cmd/novel/resource目录下
	GetResourceList() []string
}
