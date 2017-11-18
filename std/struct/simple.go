package main

import "fmt"

type Article struct {
	title   string
	author  string
	content string
}

func main() {
	var article Article

	article.title = "hello"
	article.author = "wcy100"
	article.content = "world"
	printArticle(article)
	printArticle2(&article)
}

func printArticle(article Article) {
	fmt.Println(article)
}

func printArticle2(article *Article) {
	fmt.Println(article.author)
	fmt.Println(article.title)
	fmt.Println(article.content)
}
