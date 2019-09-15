package blog

import (
	"github.com/kataras/iris"
	"../db"
	"../utils"
	"strconv"
)

func GetArticleDetails(ctx iris.Context) {
	idStr := ctx.FormValue("id")
	if idStr == "" {
		ctx.JSON(iris.Map{
			"err": "WRONG_ARTICLE_ID",
			"ok": false,
		})
		return
	}

	id, err := strconv.Atoi(idStr)
  if err != nil {
		utils.Log("routers.GetArticleDetails, id: " + err.Error())
		ctx.JSON(iris.Map{
			"err": "WRONG_ARTICLE_ID",
			"ok": false,
		})
		return
	}

	article := db.ArticleGetDetails(id)

	if article.Id == 0 {
		ctx.JSON(iris.Map{
			"err": "WRONG_ARTICLE_ID",
			"ok": false,
		})
		return
	}

	ctx.JSON(iris.Map{
		"ok": true,
		"article": article,
	})
}