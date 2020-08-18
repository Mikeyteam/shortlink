package main

import (
	"blog/models"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
)

var posts map[string]*models.Post
var counter int64

//render который мы подключили(import) будет автоматически инжектится
func homeRouterHandler(render render.Render) {
	fmt.Println(counter)
	render.HTML(200, "index", posts)
}

func writeRouteHandler(render render.Render) {
	render.HTML(200, "write", nil)
}

func editRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
		return
	}
	render.HTML(200, "write", post)
}

func deleteRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params) {
	id := params["id"]

	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	render.Redirect("/")
}

func safePostHandler(render render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkDown  := r.FormValue("content")
	contentHtml := string(blackfriday.MarkdownBasic([]byte(contentMarkDown)))

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkDown = contentMarkDown
	} else {
		id = GenerateId()
		// Создали модельку post из данных формы
		post := models.NewPost(id, title, contentHtml, contentMarkDown)
		posts[post.Id] = post //сохранили модельку в памяти. В массиве posts
	}

	render.Redirect("/")

}

func getHtmlHadler(render render.Render, r *http.Request) {
	md := r.FormValue("md")
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))

	render.JSON(200, map[string]interface{}{"html": string(htmlBytes)})
}
//Принимает строку и возвращает интерфайс. Отображение маркдауна как html
func unescape (x string) interface{} {
	return template.HTML(x)
}

func main() {
	m := martini.Classic()
    //Добавили свою кастомную функцию(динамически) в пакет template и назвалии ее FuncMap
	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	posts = make(map[string]*models.Post, 0) //создали мапу Post-ов
	// Универсальный handler вызывается на каждом запросе
	counter = 0
	m.Use(func(r *http.Request) {
		if r.URL.Path == "/write" {
			counter++
		}
	})

	//Настройки пакета github.com/martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Funcs: []template.FuncMap{unescapeFuncMap}, //зарегистрировали функцию которая будет доступна в template. И отрабаотает
		//Delims: render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
		//IndentXML: true, // Output human readable XML
		//HTMLContentType: "application/xhtml+xml", // Output XHTML content type instead of default "text/html"
	}))

	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))

	m.Get("/", homeRouterHandler) // установим обработчик, т,е функция HomeRouterHandler к
	m.Get("/write", writeRouteHandler)
	m.Get("/edit/:id", editRouteHandler)
	m.Post("/savePost", safePostHandler)
	m.Get("/delete/:id", deleteRouteHandler)
	m.Post("/gethtml", getHtmlHadler)

	m.Get("/test", func() string {
		return "dsdsa"
	})

	m.Run()
}
