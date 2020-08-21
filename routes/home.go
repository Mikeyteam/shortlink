package routes

import (
	"blog/db/documents"
	"blog/models"
	"blog/session"
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

//render который мы подключили(import) будет автоматически инжектится
func HomeRouterHandler(render render.Render, db *mgo.Database, s *session.Session) {
	fmt.Println(s.Username)

	postsCollection := db.C("Posts")
	var postDocuments []documents.PostDocument
	postsCollection.Find(nil).All(&postDocuments)

	var posts []models.Post

	//Конвертируем из документов модели
	for _, doc := range postDocuments {
		post := models.Post{Id: doc.Id, Title: doc.Title, ContentHtml: doc.ContentHtml, ContentMarkDown: doc.ContentMarkDown}
		posts = append(posts, post)
	}

	render.HTML(200, "index", posts)
}
