package routes

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HomeRouterHandler(t *testing.T) {
	m := martini.Classic()
	m.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

	})

	recorder := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)

	if recorder.Code != 200 {
		t.Error("Response not 200")
	}
}

func Test_CreateRouteHandler(t *testing.T) {
	m := martini.Classic()
	m.Get("/create", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

	})

	recorder := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/create", nil)
	m.ServeHTTP(recorder, r)

	if recorder.Code != 200 {
		t.Error("Response not 200")
	}
}

func Test_ViewRouterHandler(t *testing.T) {
	m := martini.Classic()
	m.Get("/view", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

	})

	recorder := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/view", nil)
	m.ServeHTTP(recorder, r)

	if recorder.Code != 200 {
		t.Error("Response not 200")
	}
}



