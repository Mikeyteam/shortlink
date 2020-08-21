package main

import (
	"blog/routes"
	"blog/session"
	"blog/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"html/template"
)

func main() {

	connectMongo, err := mgo.Dial("Localhost")
	if err != nil {
		panic(err)
	}
	defer connectMongo.Close()

	db := connectMongo.DB("blog")

	m := martini.Classic()
	//Замапили обект db в martini. И этот обьект мы уже можем использовать в роутах.
	//т.е мы в аргументах у роутов передаем этот обьект db
	//Так избавляемся от глобальных обьетов. martini инжектит его сам.
	m.Map(db)
	//Добавляем в martini обработчик
	m.Use(session.Middleware)

	//Добавили свою кастомную функцию(динамически) в пакет template и назвалии ее FuncMap
	unescapeFuncMap := template.FuncMap{"unescape": utils.Unescape}

	//Настройки пакета github.com/martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{unescapeFuncMap}, //зарегистрировали функцию которая будет доступна в template. И отрабаотает
		Charset:    "UTF-8",
		IndentJSON: true, // Output human readable JSON
	}))

	//Подключаем css и js
	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))

	m.Get("/", routes.HomeRouterHandler)
	m.Get("/write", routes.WriteRouteHandler)
	m.Get("/edit/:id", routes.EditRouteHandler)
	m.Post("/savePost", routes.SafePostHandler)
	m.Get("/delete/:id", routes.DeleteRouteHandler)
	m.Post("/gethtml", routes.GetHtmlHadler)
	m.Get("/login", routes.GetLoginHandler)
	m.Post("/login", routes.PostLoginHandler)

	//martini запускает web серевре и слушает порт 3000
	m.Run()
}
