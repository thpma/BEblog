package main

import (
	"github.com/kataras/iris"
	"github.com/rs/cors"
	"./blog/routers"
)

func main() {
	app := iris.New()
  corsOptions := cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowedHeaders:   []string{"*"},
    AllowCredentials: true,
  }
  corsWrapper := cors.New(corsOptions).ServeHTTP
  app.WrapRouter(corsWrapper)
	app.StaticWeb("/static", "./assets")
	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("index.html", true)
	})

	v1 := app.Party("blog")
	{
		v1.Post("/login", blog.Login)

		v1.Get("/article/list", blog.GetArticleList) // 获取首页文章列表
		v1.Post("/article/create", blog.CreateArticle) // 创建新文章
		v1.Get("/article/get", blog.GetArticleDetails) //获取某个文章
	}

	app.Run(iris.Addr(":8081"))
}