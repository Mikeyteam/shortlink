package routes

import (
	"blog/session"
	"fmt"
	"github.com/martini-contrib/render"
	"net/http"
)

func GetLoginHandler(render render.Render) {
	render.HTML(200, "login", nil)
}
//Инжектим обьект сессии который зависит от контекста реквеста
func PostLoginHandler(r *http.Request, render render.Render, s *session.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(password)

	//Меняем обьект. Сессия сама сохранится с новыми данными
	s.Username = username

	render.Redirect("/")
}
