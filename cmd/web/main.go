package main

import (
	blog_go_gin "github.com/dazai404/blog-go-gin"
	"github.com/dazai404/blog-go-gin/pkg/handler"
	"github.com/dazai404/blog-go-gin/pkg/mysql"
	"github.com/dazai404/blog-go-gin/pkg/repository"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func main() {
	db, err := mysql.NewMySQLConnection("root", "root", "3306", "blog_go_gin")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db, es)
	handlers := handler.NewHandler(repos)

	srv := new(blog_go_gin.Server)

	err = srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}

}
