package models

type Post struct {
	Id string
	Title string
	Content string
}

//Конструктор модельки post
func NewPost(id, title, content string) *Post {
	//Возвращаем ссылку на модельку
	return &Post{id,title,content}
}
