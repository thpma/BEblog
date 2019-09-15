package blog

import (
	"github.com/kataras/iris"
	"../db"
	"../utils"
	"strconv"
)

func GetArticleList(ctx iris.Context) {
	pageStr := ctx.FormValue("page")
	if pageStr == "" {
		ctx.JSON(iris.Map{
			"err": "WRONG_PAGE",
			"ok": false,
		})
		return
	}

	page, err := strconv.Atoi(pageStr)
  if err != nil {
		utils.Log("routers.GetArticleList, page: " + err.Error())
		ctx.JSON(iris.Map{
			"err": "WRONG_PAGE",
			"ok": false,
		})
		return
	}
	ctx.JSON(iris.Map{
		"ok": true,
		"list": db.ArticleGetList(page),
		"total": db.GetArticleTotalCount(),
	})
}