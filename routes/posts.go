package routes

import (
	"blog/db/documents"
	"blog/models"
	"blog/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"html/template"
	"net/http"
)

func WriteRouteHandler(render render.Render) {
	render.HTML(200, "write", nil)
}

func EditRouteHandler(render render.Render, params martini.Params, db *mgo.Database) {
	postsCollection := db.C("Posts")

	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)

	if err != nil {
		render.Redirect("/")
	}

	//Создаем новый пост
	post := models.Post{Id: postDocument.Id, Title: postDocument.Title, ContentHtml: postDocument.ContentHtml, ContentMarkDown: postDocument.ContentMarkDown}

	render.HTML(200, "write", post)
}

func DeleteRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params, db *mgo.Database) {
	id := params["id"]
	postsCollection := db.C("Posts")

	if id == "" {
		http.NotFound(w, r)
	}
	postsCollection.RemoveId(id)

	render.Redirect("/")
}

func SafePostHandler(render render.Render, r *http.Request, db *mgo.Database) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkDown := r.FormValue("content")
	contentHtml := utils.ConvertMarkDownToHtml(contentMarkDown)

	postsCollection := db.C("Posts")
	postDocument := documents.PostDocument{Id: id, Title: title, ContentHtml: contentHtml, ContentMarkDown: contentMarkDown}

	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	render.Redirect("/")

}

func GetHtmlHadler(render render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkDownToHtml(md)

	render.JSON(200, map[string]interface{}{"html": string(html)})
}

//Принимает строку и возвращает интерфайс. Отображение маркдауна как html
func unescape(x string) interface{} {
	return template.HTML(x)
}
