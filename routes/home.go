package routes

import (
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func HomeRouterHandler(render render.Render, db *mgo.Database) {
	render.HTML(200, "index", nil)
}
