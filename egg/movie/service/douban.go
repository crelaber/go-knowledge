package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-knowledge/egg/movie/model"
	"go-knowledge/egg/movie/parse"
	"log"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

func Add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			log.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}

func Start() {
	var movies []parse.DoubanMovie
	pages := parse.GetPages(BaseUrl)
	for _, page := range pages {
		doc, err := goquery.NewDocument(strings.Join([]string{BaseUrl, page.Url}, ""))
		if err != nil {
			log.Println(err)
		}
		movies = append(movies, parse.AnalysisMovie(doc)...)
	}
	fmt.Printf("%+v", movies)
	//Add(movies)
}
