package blog

import (
	"github.com/kataras/iris"
	"fmt"
	"../utils"
	"../db"
)

type loginJSON struct {
	Name			string		`json:"username"`
	Password	string		`json:"password"`
}

func Login(ctx iris.Context) {
	if db.IsUserLogged(ctx.GetHeader("auth-token")) == true {
		ctx.JSON(iris.Map{
			"err": "LOGGED",
			"ok": false,
		})
		return
	}

	var p loginJSON
  if err := ctx.ReadJSON(&p); err != nil {
		fmt.Println(err)
	}

	user := db.IsUserExist(p.Name)

	if user.Id == 0 {
		ctx.JSON(iris.Map{
			"err": "USER_NOT_EXIST",
			"ok": false,
		})
		return
	}

	if !utils.PasswordComparison(p.Password, user.Password) {
		ctx.JSON(iris.Map{
			"err": "PASSWORD_INCORRECT",
			"ok": false,
		})
		return
	}

	token := db.UpdateUserToken(user.Id)
	if token == "" {
		ctx.JSON(iris.Map{
			"err": "DATABASE_ERROR",
			"ok": false,
		})
		return
	}

	ctx.JSON(iris.Map{
		"ok": true,
		"auth": token,
	})
}