package routes

import (
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"shorthref/db/documents"
	"shorthref/models"
)

func ViewRouterHandler(render render.Render, db *mgo.Database) {
	hrefsCollection := db.C("Hrefs")
	var hrefDocuments []documents.HrefDocument
	_ = hrefsCollection.Find(nil).All(&hrefDocuments)

	var hrefs []models.Href

	for _, doc := range hrefDocuments {
		href := models.Href{Id: doc.Id, LongHref: doc.LongHref, ShortHref: doc.ShortHref}
		hrefs = append(hrefs, href)
	}

	render.HTML(200, "view", hrefs)
}
