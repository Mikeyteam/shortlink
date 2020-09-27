package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
	"net/url"
	"shorthref/db/documents"
	"shorthref/models"
	"shorthref/utils"
)

func CreateRouteHandler(render render.Render) {
	render.HTML(200, "create", nil)
}

func EditRouteHandler(render render.Render, params martini.Params, db *mgo.Database) {
	hrefsCollection := db.C("Hrefs")
	id := params["id"]

	hrefDocument := documents.HrefDocument{}
	err := hrefsCollection.FindId(id).One(&hrefDocument)

	if err != nil {
		render.Redirect("/")
	}

	href := models.Href{Id: hrefDocument.Id, LongHref: hrefDocument.LongHref, ShortHref: hrefDocument.ShortHref}

	render.HTML(200, "create", href)
}

func DeleteRouteHandler(render render.Render, w http.ResponseWriter, r *http.Request, params martini.Params, db *mgo.Database) {
	id := params["id"]
	hrefsCollection := db.C("Hrefs")

	if id == "" {
		http.NotFound(w, r)
	}
	_ = hrefsCollection.RemoveId(id)

	render.Redirect("/view")
}

func SafeHrefHandler(render render.Render, r *http.Request, db *mgo.Database) {
	id := r.FormValue("id")
	longHref := r.FormValue("longHref")
	shortHref := r.FormValue("shortHref")

	_, errLong := url.ParseRequestURI(longHref)
	_, errShort := url.ParseRequestURI(shortHref)
	if errShort != nil || errLong != nil  {
		render.Redirect("/create")
		return
	}

	hrefsCollection := db.C("Hrefs")
	hrefDocument := documents.HrefDocument{Id: id, LongHref: longHref, ShortHref: shortHref}

	if id != "" {
		_ = hrefsCollection.UpdateId(id, hrefDocument)
	} else {
		id = utils.GenerateId()
		hrefDocument.Id = id
		_ = hrefsCollection.Insert(hrefDocument)
	}

	render.Redirect("/view")

}
