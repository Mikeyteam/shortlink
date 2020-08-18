package models

type Post struct {
	Id              string
	Title           string
	ContentHtml     string
	ContentMarkDown string
}

//Конструктор модельки post
func NewPost(id, title, contentHtml, ContentMarkDown string) *Post {
	//Возвращаем ссылку на модельку
	return &Post{id, title, contentHtml, ContentMarkDown}
}
