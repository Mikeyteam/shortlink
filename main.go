package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"html/template"
	"blog/models"
)

var posts map[string]*models.Post

func homeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Andrey!</h1>") // отправляем данные на клиентскую сторону. Записываем в перемную w -
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html" ) //ParseFiles возврщает на обьект шаблона и ошибку
	//так как в индексе есть header и т.д их тоже нужно парсить

	//Если есть ошибка просто выведем в браузере
	if err != nil {
		fmt.Fprintf(w, "Не найден шаблон")
	}
	t.ExecuteTemplate(w, "index", posts) //Выполняем - подключаем шаблон

	fmt.Println(posts)
}

func writeRouteHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Не найден шаблон")
	}
	t.ExecuteTemplate(w, "write", nil) //Выполняем - подключаем шаблон
}

func editRouteHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	id := r.FormValue("id")
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
		return
	}

	t.ExecuteTemplate(w, "write", post)
}

func deleteRouteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	http.Redirect(w,r, "/",302)
}

func safePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = GenerateId()
		// Создали модельку post из данных формы
		post := models.NewPost(id, title, content)
		posts[post.Id] = post //сохранили модельку в памяти. В массиве posts
	}

	http.Redirect(w,r, "/",302)

}

func main() {
	posts = make(map[string]*models.Post, 0) //создали мапу Post-ов

	http.HandleFunc("/", homeRouterHandler) // установим обработчик, т,е функция HomeRouterHandler к
	http.HandleFunc("/write", writeRouteHandler)
	http.HandleFunc("/edit", editRouteHandler)
	http.HandleFunc("/savePost", safePostHandler)
	http.HandleFunc("/delete", deleteRouteHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	err := http.ListenAndServe(":3000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}