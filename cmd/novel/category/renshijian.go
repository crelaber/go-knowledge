package category

type RenshijianNovel struct {
}

func NewNovelRenshijian() *RenshijianNovel {
	return &RenshijianNovel{}
}

func (r *RenshijianNovel) GetResourceList() []string {
	return []string{
		"1-20.json", //1-20集
		"41-60.json",
		"61-80.json",
	}
}

var _ NovelResource = (*RenshijianNovel)(nil)
