package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("secret123"))
	fmt.Println(store)
	m.Use(sessions.Sessions("my_session", store))

	m.Get("/set", func(session sessions.Session) string {
		session.Set("hello", "world")
		return "OK"
	})

	m.Get("/get", func(session sessions.Session) string {
		v := session.Get("hello")
		if v == nil {
			return ""
		}
		return v.(string)
	})

	m.Run()
}
