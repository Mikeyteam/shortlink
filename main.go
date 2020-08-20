package main

import (
	"blog/db/documents"
	"blog/models"
	"blog/session"
	"blog/utils"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"html/template"
	"net/http"
	"time"
)

const COOKIE_NAME = "sessionId"

//Глобальная коллекция к ней обращаемся из наших handler
var postsCollection *mgo.Collection
var counter int64
//Указатель на сессию
var inMemorySession *session.Session

//render который мы подключили(import) будет автоматически инжектится
func homeRouterHandler(render render.Render, r *http.Request) {
	cookie,_ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(inMemorySession.Get(cookie.Value))
	}

	var postDocuments []documents.PostDocument
	postsCollection.Find(nil).All(&postDocuments)

	var posts []models.Post

	//Конвертируем из документов модели
	for _, doc := range postDocuments {
		post := models.Post{Id: doc.Id, Title: doc.Title, ContentHtml: doc.ContentHtml, ContentMarkDown: doc.ContentMarkDown}
		posts = append(posts, post)
	}

	fmt.Println(counter)
	render.HTML(200, "index", posts)
}

func writeRouteHandler(render render.Render) {
	render.HTML(200, "write", nil)
}

func getLoginHandler(render render.Render) {
	render.HTML(200, "login", nil)
}

func postLoginHandler(r *http.Request, w http.ResponseWriter, render render.Render) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(password)

    sessionId := inMemorySession.Init(username)
    cookie := &http.Cookie{
    	Name: COOKIE_NAME,
    	Value: sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	render.Redirect("/")
}

func editRouteHandler(render render.Render, params martini.Params) {
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

func deleteRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]

	if id == "" {
		http.NotFound(w, r)
	}
	postsCollection.RemoveId(id)

	render.Redirect("/")
}

func safePostHandler(render render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkDown := r.FormValue("content")
	contentHtml := utils.ConvertMarkDownToHtml(contentMarkDown)

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

func getHtmlHadler(render render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkDownToHtml(md)

	render.JSON(200, map[string]interface{}{"html": string(html)})
}

//Принимает строку и возвращает интерфайс. Отображение маркдауна как html
func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {
	//Создали новую сессию
	inMemorySession = session.NewSession()

	session, err := mgo.Dial("Localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	postsCollection = session.DB("blog").C("posts")

	m := martini.Classic()
	//Добавили свою кастомную функцию(динамически) в пакет template и назвалии ее FuncMap
	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	// Универсальный handler вызывается на каждом запросе
	counter = 0
	m.Use(func(r *http.Request) {
		if r.URL.Path == "/write" {
			counter++
		}
	})

	//Настройки пакета github.com/martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{unescapeFuncMap}, //зарегистрировали функцию которая будет доступна в template. И отрабаотает
		Charset:    "UTF-8",
		IndentJSON: true, // Output human readable JSON
	}))

	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))

	m.Get("/", homeRouterHandler)
	m.Get("/write", writeRouteHandler)
	m.Get("/edit/:id", editRouteHandler)
	m.Post("/savePost", safePostHandler)
	m.Get("/delete/:id", deleteRouteHandler)
	m.Post("/gethtml", getHtmlHadler)
	m.Get("/login", getLoginHandler)
	m.Post("/login", postLoginHandler)

	m.Get("/test", func() string {
		return "dsdsa"
	})

	m.Run()
}
