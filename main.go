package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"shorthref/routes"
)

func main() {
	connectMongo, err := mgo.Dial("Localhost")

	if err != nil {
		panic(err)
	}
	defer connectMongo.Close()
	db := connectMongo.DB("dbhref")

	m := martini.Classic()
	m.Map(db)

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	staticOption := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOption))

	m.Get("/", routes.HomeRouterHandler)
	m.Get("/view", routes.ViewRouterHandler)
	m.Get("/create", routes.CreateRouteHandler)
	m.Get("/edit/:id", routes.EditRouteHandler)
	m.Post("/saveHref", routes.SafeHrefHandler)
	m.Get("/delete/:id", routes.DeleteRouteHandler)

	m.Run()

}
