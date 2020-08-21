package session

import (
	"blog/utils"
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

const CookieName = "sessionId"

type Session struct {
	id       string
	Username string
}

type Store struct {
	data map[string]*Session
}

func NewSessionStore() *Store {
	session := new(Store)
	session.data = make(map[string]*Session)

	return session
}

func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(CookieName)
	if cookie != nil {
		return cookie.Value
	}
	sessionId := utils.GenerateId()
	cookie = &http.Cookie{
		Name:    CookieName,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)

	return sessionId

}

func (store *Store) Get(sessionId string) *Session {
	session := store.data[sessionId]
	if session == nil {
		return &Session{id: sessionId}
	}

	return session
}

func (store *Store) Set(session *Session) {
	store.data[session.id] = session
}

var sessionStore = NewSessionStore()

//При реквесте мы получаем его данные и создаем обьект на основании полученных данных
//Типо обработчик промежуточный. Для сессии нужен текущий контекст
func Middleware(context martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := ensureCookie(r, w)
	session := sessionStore.Get(sessionId)
	context.Map(session)
	context.Next()
	// Что после Next вызывается когда Request уже произошел
	//Подменили обьект. Т.е перед реквестом получили, выполнили реквест и после него обратно сохранили
	sessionStore.Set(session)
}
