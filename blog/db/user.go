package db

import (
	"time"
	"../models"
	"../utils"
	"strconv"
)

func GetUserName(auth string) string {
	user := new(models.Users)
	user.Token = auth
	has, err := mysql.Get(user)
	if err != nil {
		utils.Log("db.GetUserName: " + err.Error())
	}
	if !has || err != nil {
		return "";
	}
	return user.Name;
}

func IsUserLogged(auth string) bool {
	user := new(models.Users)
	user.Token = auth
	has, err := mysql.Get(user)
	if err != nil {
		utils.Log("db.IsUserLogged: " + err.Error())
		return false
	}
	if !has {
		return false
	}
	if user.ExpireTime - time.Now().Unix() < 1 {
		return false
	}
	return true;
}

func IsUserExist(name string) *models.Users {
	user := new(models.Users)
	user.Name = name
	has, err := mysql.Get(user)
	if err != nil {
		utils.Log("db.IsUserExist: " + err.Error())
	}
	if !has || err != nil {
		user.Id = 0
	}
	return user
}

func UpdateUserToken(uid int) string {
	user := new(models.Users)
	user.Id = uid
	user.ExpireTime = time.Now().Unix() + 86400
	user.Token = utils.GETMD5(utils.GetRandomString(32)+strconv.FormatInt(time.Now().Unix(), 10)+utils.GetRandomString(32))
	_, err := mysql.ID(user.Id).Update(user)
	if err != nil {
		utils.Log("db.UpdateUserToken: " + err.Error())
		return ""
	}
	return user.Token
}