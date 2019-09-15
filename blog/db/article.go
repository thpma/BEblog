package db

import (
	"../models"
	"../utils"
	"fmt"
)

var article_total_count int64

func InitArticleTotalCount() {
	article_total_count = totalArticles()
	fmt.Println("article total count:", article_total_count)
}

func GetArticleTotalCount() int64 {
	return article_total_count
}

func ArticleGetDetails(id int) *models.Articles {
	article := new(models.Articles)
	article.Id = id
	has, err := mysql.Get(article)
	if err != nil {
		utils.Log("db.ArticleGetDetails: " + err.Error())
	}
	if !has || err != nil {
		article.Id = 0
	}
	return article
}

func totalArticles() int64 {
	counts, err := mysql.Count(new(models.Articles))
	if err != nil {
		utils.Log("db.TotalArticles: " + err.Error())
		return 0
	}
	return counts
}

func ArticleGetList(page int) []models.Articles {
	if page == 1 {
		page = 0
	} else {
		page = (page-1) * 20
	}
	var list []models.Articles
	rows, err := mysql.OrderBy("id DESC").Limit(20, page).Rows(new(models.Articles))
	defer rows.Close()
	if err != nil {
		utils.Log("db.ArticleGetList, rows: " + err.Error())
		return list
	}
	var p models.Articles
	for rows.Next() {
		rows.Scan(&p)
		list = append(list, p)
	}
	return list
}

func CreateArticle(title string, content string, author string, poster string, category string) int {
	article := new(models.Articles)
	article.Title = title
	article.Content = content
	article.Author = author
	article.Poster = poster
	article.Category = category
	_, err := mysql.Insert(article)
	if err != nil {
		utils.Log("db.CreateArticle: " + err.Error())
		return 0
	}
	article_total_count ++
	return article.Id
}