package blog

import (
	"github.com/kataras/iris"
	"time"
	"../utils"
	"strconv"
	"os"
	"io"
	"../db"
)

func CreateArticle(ctx iris.Context) {
	auth := ctx.GetHeader("auth-token")
	if db.IsUserLogged(auth) == false {
		ctx.JSON(iris.Map{
			"err": "NOT_LOGGED",
			"ok": false,
		})
		return
	}

	file, info, err := ctx.FormFile("cover")
	defer file.Close()
	if err != nil {
		utils.Log("CreateArticle, file: " + err.Error())
		ctx.JSON(iris.Map{
			"ok": false,
			"err": "UPLOAD_COVER_FAILED",
		})
		return
	}

	filename := utils.GETMD5(info.Filename+strconv.FormatInt(time.Now().Unix(), 10)+utils.GetRandomString(6)) + utils.GetFileSuffixName(info.Filename)

	out, err := os.OpenFile("./assets/images/poster/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	if err != nil {
		utils.Log("CreateArticle, out: " + err.Error())
		ctx.JSON(iris.Map{
			"ok": false,
			"err": "UPLOAD_COVER_FAILED",
		})
		return
	}

	io.Copy(out, file)
	
	articleId := db.CreateArticle(ctx.FormValue("title"), ctx.FormValue("content"), db.GetUserName(auth), filename, ctx.FormValue("category"))
	if articleId == 0 {
		ctx.JSON(iris.Map{
			"ok": false,
			"err": "DATABASE_ERROR",
		})
		return
	}

	ctx.JSON(iris.Map{
		"ok": true,
		"id": articleId,
	})
}