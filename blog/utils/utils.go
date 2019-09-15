package utils

import (
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"time"
	"path"
	"os"
	"fmt"
)

func Log(str string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	f, err := os.OpenFile("./runtime.log", os.O_APPEND|os.O_CREATE, 0666)
  defer f.Close()
  if err != nil {
		fmt.Println(err)
  } else {
    _,err=f.Write([]byte("["+time+"] "+str+"\r\n"))
    if err != nil {
			fmt.Println(err)
		}
	}
}

func GetFileSuffixName(filename string) string {
	return path.Ext(path.Base(filename))
}

func PasswordComparison(password string, password2 string) bool {
	output := GETMD5("!#)!@*%#^"+password+"t(%@am&")
	if output == password2 {
		return true
	}
	return false
}

func GETMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
